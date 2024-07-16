package db

import (
	"context"
	"database/sql"
	"time"

	"casita/models"
)

type AccountStore interface {
	InsertAccount(ctx context.Context, tx *sql.Tx, account *models.Account) (int, error)
	GetAccountByID(ctx context.Context, tx *sql.Tx, id int) (models.Account, error)
	GetAllAccounts(ctx context.Context, tx *sql.Tx) ([]models.Account, error)
	UpdateAccount(ctx context.Context, tx *sql.Tx, id int, params *models.UpdateAccount) error
	DeleteAccountByID(ctx context.Context, tx *sql.Tx, id int) error
}

type PGAccountStore struct {
	client *sql.DB
}

func NewPGAccountStore(client *sql.DB) *PGAccountStore {
	return &PGAccountStore{
		client: client,
	}
}

func (s *PGAccountStore) InsertAccount(ctx context.Context, tx *sql.Tx, account *models.Account) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var newID int

	query := `
            INSERT into accounts 
              (name, user_id, entity, currency, created_at, updated_at)
            VALUES 
              ($1, $2, $3, $4, $5, $6)
            RETURNING id
    `

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query,
			account.Name,
			account.UserID,
			account.Entity,
			account.Currency,
			time.Now(),
			time.Now(),
		).Scan(&newID)
	} else {
		err = s.client.QueryRowContext(ctx, query,
			account.Name,
			account.UserID,
			account.Entity,
			account.Currency,
			time.Now(),
			time.Now(),
		).Scan(&newID)
	}
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (s *PGAccountStore) GetAccountByID(ctx context.Context, tx *sql.Tx, id int) (models.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var account models.Account

	query := `
            SELECT 
              (name, id, user_id, entity, currency, created_at, updated_at)
            from accounts
            WHERE
            id=$1`
	var err error
	if tx != nil {
		err = s.client.QueryRowContext(ctx, query, id).Scan(
			&account.ID,
			&account.Name,
			&account.UserID,
			&account.Entity,
			&account.Currency,
			time.Now(),
			time.Now(),
		)
	} else {
		err = s.client.QueryRowContext(ctx, query, id).Scan(
			&account.ID,
			&account.Name,
			&account.UserID,
			&account.Entity,
			&account.Currency,
			time.Now(),
			time.Now(),
		)
	}
	if err != nil {
		return account, err
	}

	return account, nil
}

func (s *PGAccountStore) GetAllAccounts(ctx context.Context, tx *sql.Tx) ([]models.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var accounts []models.Account

	query := `
            SELECT 
              id, name, user_id, entity, currency, created_at, updated_at
            FROM accounts`

	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryContext(ctx, query)
	} else {
		rows, err = s.client.QueryContext(ctx, query)
	}
	if err != nil {
		return accounts, err
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err = rows.Scan(
			&account.ID,
			&account.Name,
			&account.UserID,
			&account.Entity,
			&account.Currency,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PGAccountStore) UpdateAccount(ctx context.Context, tx *sql.Tx, id int, params *models.UpdateAccount) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	now := time.Now()

	query := `
        UPDATE accounts SET 
            name = COALESCE($1, name),
            entity = COALESCE($2, entity),
            currency = COALESCE($3, currency),
            updated_at = %5
        WHERE id = $6
    `

	var err error
	if tx != nil {
		_, err = s.client.ExecContext(
			ctx,
			query,
			params.Name,
			params.Entity,
			params.Currency,
			now,
			id)
	} else {
		_, err = s.client.ExecContext(
			ctx,
			query,
			params.Name,
			params.Entity,
			params.Currency,
			now,
			id)
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *PGAccountStore) DeleteAccountByID(ctx context.Context, tx *sql.Tx, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "DELETE FROM accounts WHERE id = $1"

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, id)
	} else {
		_, err = s.client.ExecContext(ctx, query, id)
	}
	if err != nil {
		return err
	}

	return nil
}
