package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"finance/models"
)

type TransactionStore interface {
	InsertTransaction(ctx context.Context, tx *sql.Tx, txn *models.Transaction) (int, error)
	GetTransactionByID(ctx context.Context, tx *sql.Tx, id int) (*models.Transaction, error)
	GetAllTransactions(ctx context.Context, tx *sql.Tx) ([]*models.Transaction, error)
	UpdateTransaction(ctx context.Context, tx *sql.Tx, id int, params *models.UpdateTransaction) error
	DeleteTransactionByID(ctx context.Context, tx *sql.Tx, id int) error
}

type PGTransactionStore struct {
	client *sql.DB
}

func NewPGTransactionStore(client *sql.DB) *PGTransactionStore {
	return &PGTransactionStore{
		client: client,
	}
}

func (s *PGTransactionStore) InsertTransaction(ctx context.Context, tx *sql.Tx, txn *models.Transaction) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var newID int

	query := `
            INSERT into transactions 
                (concept, description, value, date, relevance, currency, account, created_at, updated_at)
            values 
                ($1, $2, $3, $4, $5, $6)
            returning id`

	var err error
	if tx != nil {
		err = s.client.QueryRowContext(ctx, query,
			txn.Concept,
			txn.Description,
			txn.Value,
			txn.Date,
			txn.Relevance,
			txn.Currency,
			txn.Account,
			time.Now(),
			time.Now(),
		).Scan(&newID)
	} else {
		err = s.client.QueryRowContext(ctx, query,
			txn.Concept,
			txn.Description,
			txn.Value,
			txn.Date,
			txn.Relevance,
			txn.Currency,
			txn.Account,
			time.Now(),
			time.Now(),
		).Scan(&newID)
	}
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (s *PGTransactionStore) GetTransactionByID(ctx context.Context, tx *sql.Tx, id int) (*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var txn *models.Transaction

	query := `
            SELECT 
                (concept, description, value, date, relevance, currency, account, created_at, updated_at)
            from transactions
            WHERE
            id=$1`

	var err error
	if tx != nil {
		err = s.client.QueryRowContext(ctx, query, id).Scan(
			&txn.ID,
			&txn.Description,
			&txn.Value,
			&txn.Date,
			&txn.Relevance,
			&txn.Currency,
			&txn.Account,
			&txn.CreatedAt,
			&txn.UpdatedAt,
		)
	} else {
		err = s.client.QueryRowContext(ctx, query, id).Scan(
			&txn.ID,
			&txn.Description,
			&txn.Value,
			&txn.Date,
			&txn.Relevance,
			&txn.Currency,
			&txn.Account,
			&txn.CreatedAt,
			&txn.UpdatedAt,
		)
	}
	if err != nil {
		return txn, err
	}

	return txn, nil
}

func (s *PGTransactionStore) GetAllTransactions(ctx context.Context, tx *sql.Tx) ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var txs []*models.Transaction

	query := `
            SELECT (id, first_name, last_name, email, password, created_at, updated_at)
            from transactions`

	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryContext(ctx, query)
	} else {
		rows, err = s.client.QueryContext(ctx, query)
	}

	if err != nil {
		return txs, err
	}

	for rows.Next() {
		var tx models.Transaction
		err := s.client.QueryRowContext(ctx, query).Scan(
			&tx.ID,
			&tx.Description,
			&tx.Value,
			&tx.Date,
			&tx.Relevance,
			&tx.Currency,
			&tx.Account,
			&tx.CreatedAt,
			&tx.UpdatedAt,
		)
		if err != nil {
			return txs, err
		}
	}

	return txs, nil
}

func (s *PGTransactionStore) UpdateTransaction(ctx context.Context, tx *sql.Tx, id int, params *models.UpdateTransaction) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	now := time.Now()

	query := fmt.Sprintf(`
        UPDATE transactions set 
            concept = COALESCE($1, concept)
            description = COALESCE($2, description)
            value = COALESCE($3, value)
            date = COALESCE($4, date)
            relevance = COALESCE($5, relevance)
            currency = COALESCE($6, currency)
            account = COALESCE($7, account)
            updated_at = %s
        WHERE
            id = $8
    `, now)

	var err error
	if tx != nil {
		_, err = s.client.ExecContext(
			ctx,
			query,
			params.Concept,
			params.Description,
			params.Value,
			params.Date,
			params.Relevance,
			params.Currency,
			params.Account,
			id)
	} else {
		_, err = s.client.ExecContext(
			ctx,
			query,
			params.Concept,
			params.Description,
			params.Value,
			params.Date,
			params.Relevance,
			params.Currency,
			params.Account,
			id)
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *PGTransactionStore) DeleteTransactionByID(ctx context.Context, tx *sql.Tx, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := "DELETE FROM transactions WHERE id = $1"

	var err error
	if tx != nil {
		_, err = s.client.ExecContext(ctx, query, id)
	} else {
		_, err = s.client.ExecContext(ctx, query, id)
	}
	if err != nil {
		return err
	}

	return nil
}
