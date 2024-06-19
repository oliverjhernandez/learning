package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"finance/models"
)

type UserStore interface {
	// TODO: Inject context instead of creating it in each method
	InsertUser(user *models.User) (int, error)
	GetUserByID(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateAUsersName(id int, firstName, lastName string) error
	DeleteUserByID(id int) error
}

type PGUserStore struct {
	client *sql.DB
}

func NewPGUserStore(client *sql.DB) *PGUserStore {
	return &PGUserStore{
		client: client,
	}
}

func (s *PGUserStore) InsertUser(user *models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	var newID int

	query := `
            INSERT into users 
                (first_name, last_name, email, password, created_at, updated_at)
            values 
                ($1, $2, $3, $4, $5, $6)
            returning id`

	err := s.client.QueryRowContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Passwd,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (s *PGUserStore) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	var user models.User

	query := `
            SELECT (id, first_name, last_name, email, password, created_at, updated_at)
            from users
            WHERE
            id=$1`

	err := s.client.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Passwd,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *PGUserStore) GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	var users []models.User

	query := `
            SELECT (id, first_name, last_name, email, password, created_at, updated_at)
            from users`

	rows, err := s.client.QueryContext(ctx, query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user models.User
		err := s.client.QueryRowContext(ctx, query).Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Passwd,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return users, err
		}
	}

	return users, nil
}

func (s *PGUserStore) UpdateUser(id int, params *models.UpdateUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	now := time.Now()

	query := fmt.Sprintf(`
        UPDATE users set 
            first_name = COALESCE($1, first_name),
            last_name = COALESCE($2, last_name),
            email = COALESCE($3, email),
            updated_at = %s
        WHERE
            id = $4
    `, now)

	_, err := s.client.ExecContext(ctx, query, params.FirstName, params.LastName, params.Email, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PGUserStore) DeleteUserByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id = $1"

	_, err := s.client.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
