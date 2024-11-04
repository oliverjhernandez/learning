package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// TODO: centralize this into a single interface
type UserStore interface {
	InsertUser(ctx context.Context, tx *sql.Tx, user *User) (*User, error)
	GetUserByID(ctx context.Context, tx *sql.Tx, id int) (*User, error)
	GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (*User, error)
	GetAllUsers(ctx context.Context, tx *sql.Tx) ([]*User, error)
	UpdateUser(ctx context.Context, tx *sql.Tx, id int, params *UpdateUser) (*User, error)
	DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error
}

type PGUserStore struct {
	DB *sql.DB
}

func NewPGUserStore(client *sql.DB) *PGUserStore {
	return &PGUserStore{
		DB: client,
	}
}

func (s *PGUserStore) InsertUser(ctx context.Context, tx *sql.Tx, user *User) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var newID int

	query := `
    INSERT into users 
      (first_name, last_name, email, passwd_hash, created_at, updated_at)
    VALUES 
      ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id`

	args := []interface{}{
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswdHash,
		time.Now(),
		time.Now(),
	}

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, args...).Scan(&newID)
	} else {
		err = s.DB.QueryRowContext(ctx, query, args...).Scan(&newID)
	}
	if err != nil {
		return nil, err
	}

	userResp, err := s.GetUserByID(ctx, nil, newID)
	if err != nil {
		return nil, err
	}

	return userResp, nil
}

func (s *PGUserStore) GetUserByID(ctx context.Context, tx *sql.Tx, id int) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var user User

	query := `
    SELECT id, first_name, last_name, email, passwd_hash, created_at, updated_at
    FROM users
    WHERE id=$1`

	args := []interface{}{
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswdHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	}

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(args...)
	} else {
		err = s.DB.QueryRowContext(ctx, query, id).Scan(args...)
	}
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (s *PGUserStore) GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var user User

	query := `
    SELECT id, first_name, last_name, email, passwd_hash, created_at, updated_at
    FROM users
    WHERE email=$1`

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, email).Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PasswdHash,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	} else {
		err = s.DB.QueryRowContext(ctx, query, email).Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PasswdHash,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	}
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (s *PGUserStore) GetAllUsers(ctx context.Context, tx *sql.Tx) ([]*User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var users []*User

	query := `
    SELECT id, first_name, last_name, email, created_at, updated_at
    FROM users`

	var err error
	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.QueryContext(ctx, query)
	} else {
		rows, err = s.DB.QueryContext(ctx, query)
	}
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User

		args := []interface{}{
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		}
		err = rows.Scan(args...)
		if err != nil {
			return users, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (s *PGUserStore) UpdateUser(ctx context.Context, tx *sql.Tx, id int, params *UpdateUser) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	now := time.Now()
	var newID int

	setClauses := []string{}
	args := []interface{}{}
	argID := 1

	if params.FirstName != "" {
		setClauses = append(setClauses, fmt.Sprintf("first_name = $%d", argID))
		args = append(args, params.FirstName)
		argID++
	}
	if params.LastName != "" {
		setClauses = append(setClauses, fmt.Sprintf("last_name = $%d", argID))
		args = append(args, params.LastName)
		argID++
	}
	if params.Email != "" {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argID))
		args = append(args, params.Email)
		argID++
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", argID))
	args = append(args, now)
	argID++

	args = append(args, id)

	query := fmt.Sprintf(`
    UPDATE users 
    SET %s
    WHERE id = $%d`,
		strings.Join(setClauses, ", "), argID)

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, args...).Scan(&newID)
	} else {
		err = s.DB.QueryRowContext(ctx, query, args...).Scan(&newID)
	}
	if err != nil {
		return nil, err
	}

	userResp, err := s.GetUserByID(ctx, nil, newID)
	if err != nil {
		return nil, err
	}

	return userResp, nil
}

func (s *PGUserStore) DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
    DELETE 
    FROM users 
    WHERE id = $1`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, id)
	} else {
		_, err = s.DB.ExecContext(ctx, query, id)
	}
	if err != nil {
		return err
	}

	return nil
}
