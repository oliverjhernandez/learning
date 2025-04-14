package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"greenlight/internal/validator"
	"time"
)

type Account struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	EntityID  int       `json:"entity_id"`
	Balance   int       `json:"balance"`
}

func ValidateAccount(v *validator.Validator, a *Account) {
	// Title
	v.Check(a.Title != "", "title", "must be provided")
	v.Check(len(a.Title) >= 2, "title", "must be at least 2 bytes long")

	// UserID
	v.Check(a.UserID != 0, "user_id", "must be provided")

	// EntityID
	v.Check(a.EntityID != 0, "entity_id", "must be provided")
}

type AccountsModel struct {
	DB *sql.DB
}

func (m AccountsModel) Insert(account *Account) error {
	query := `
    INSERT INTO accounts
      (title, user_id, entity_id, created_at, updated_at)
    VALUES
      ($1, $2, $3, $4, $5)
    RETURNING id, created_at, updated_at;
  `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{
		account.Title,
		account.UserID,
		account.EntityID,
		time.Now().UTC(),
		time.Now().UTC(),
	}

	fmt.Printf("%+v\n", args...)

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt)
}

func (m AccountsModel) Get(id int64) (*Account, error) {
	query := `
    SELECT
      id,
      created_at,
      title,
      user_id,
      entity_id
    FROM accounts
    WHERE id = $1;
  `

	var account Account

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&account.ID,
		&account.CreatedAt,
		&account.Title,
		&account.UserID,
		&account.EntityID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &account, nil
}

func (m AccountsModel) GetAll(title string, userID int, entityID int, filters Filters) ([]*Account, Metadata, error) {
	query := fmt.Sprintf(`
  SELECT count(*) OVER(), id, created_at, title, user_id, entity_id
  FROM accounts
  WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
    AND (user_id = $2 OR $2 = 0)
    AND (entity_id = $3 OR $3 = 0)
  ORDER BY %s %s, id ASC
  LIMIT $4 OFFSET $5`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{title, userID, entityID, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	accounts := []*Account{}

	for rows.Next() {
		var account Account

		err := rows.Scan(
			&totalRecords,
			&account.ID,
			&account.CreatedAt,
			&account.Title,
			&account.UserID,
			&account.EntityID,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		accounts = append(accounts, &account)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return accounts, metadata, nil
}

func (m AccountsModel) Update(account *Account) error {
	query := `
  UPDATE accounts
  SET title = $1, entity_id = $2
  WHERE id = $3
  RETURNING id
  `

	args := []any{
		account.Title,
		account.EntityID,
		account.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&account.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err

		}
	}

	return nil
}

func (m AccountsModel) Delete(id int64) error {
	query := `
  DELETE FROM movies
  WHERE id = $1
  `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
