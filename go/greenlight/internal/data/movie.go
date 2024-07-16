package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"greenlight/internal/validator"

	"github.com/lib/pq"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConfilct   = errors.New("edit conflict")
)

type MovieModel struct {
	DB *sql.DB
}

func (m *MovieModel) Insert(movie *Movie) error {
	query := `
    INSERT INTO movies (title, year, runtime, genres) 
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_at, version`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query,
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres)).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m *MovieModel) Get(id int) (*Movie, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	var movie Movie

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
    SELECT id, created_at, title, year, runtime, genres, version 
    FROM movies
    WHERE id = $1`

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &movie, nil
}

func (m *MovieModel) Update(movie *Movie) error {
	query := `
    UPDATE movies
    SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
    WHERE id = $5 AND version = $6
    RETURNING version`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query,
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres),
		movie.ID,
		movie.Version,
	).Scan(&movie.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConfilct
		default:
			return err
		}
	}

	return nil
}

func (m *MovieModel) Delete(id int) error {
	query := `
    DELETE from movies
    WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m *MovieModel) GetAll(title string, genres []string, filter Filters) ([]*Movie, Metadata, error) {
	query := fmt.Sprintf(`
  SELECT count(*) OVER(), id, created_at, title, year, runtime, genres, version
  FROM movies
  WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
    AND (genres @> $2 OR $2 = '{}')
  ORDER BY %s %s, id ASC
  LIMIT $3 OFFSET $4`, filter.sortColumn(), filter.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{title, pq.Array(genres), filter.limit(), filter.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	movies := []*Movie{}

	for rows.Next() {
		var movie Movie

		err := rows.Scan(
			&totalRecords,
			&movie.ID,
			&movie.CreatedAt,
			&movie.Title,
			&movie.Year,
			&movie.Runtime,
			pq.Array(&movie.Genres),
			&movie.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		movies = append(movies, &movie)
	}

	if err := rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	md := calculateMetadata(totalRecords, filter.Page, filter.PageSize)

	return movies, md, nil
}

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, m *Movie) {
	// Title
	v.Check(m.Title != "", "title", "must be provided")
	v.Check(len(m.Title) <= 500, "title", "must not be more than 500 bytes long")

	// Year
	v.Check(m.Year != 0, "year", "must be provided")
	v.Check(m.Year >= 1888, "year", "must be greater than 1888")
	v.Check(m.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	// Runtime
	v.Check(m.Runtime != 0, "runtime", "must be provided")
	v.Check(m.Runtime > 0, "runtime", "must be a positive integer")

	// Genres
	v.Check(m.Genres != nil, "genres", "must be provided")
	v.Check(len(m.Genres) >= 1, "genres", "at least one genre must be provided")
	v.Check(len(m.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(m.Genres), "genres", "must not contain duplicate values")
}
