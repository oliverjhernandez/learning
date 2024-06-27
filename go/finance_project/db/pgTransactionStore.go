package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"finance/models"
)

type TransactionStore interface {
	InsertTransaction(ctx context.Context, tx *sql.Tx, txn *models.Transaction) (int, error)
	GetTransactionByID(ctx context.Context, tx *sql.Tx, id int) (models.Transaction, error)
	GetAllTransactions(ctx context.Context, tx *sql.Tx) ([]models.Transaction, error)
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
                (concept, description, value, date, relevance, account_id, created_at, updated_at)
            values 
                ($1, $2, $3, $4, $5, $6, $7, $8)
            returning id
    `

	var err error
	if tx != nil {
		err = s.client.QueryRowContext(ctx, query,
			txn.Concept,
			txn.Description,
			txn.Value,
			txn.Date,
			txn.Relevance,
			txn.AccountID,
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
			txn.AccountID,
			time.Now(),
			time.Now(),
		).Scan(&newID)
	}
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (s *PGTransactionStore) GetTransactionByID(ctx context.Context, tx *sql.Tx, id int) (models.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var txn models.Transaction

	query := `
            SELECT 
                id, concept, description, value, date, relevance, account_id, created_at, updated_at
            FROM transactions
            WHERE id=$1
    `

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(
			&txn.ID,
			&txn.Concept,
			&txn.Description,
			&txn.Value,
			&txn.Date,
			&txn.Relevance,
			&txn.AccountID,
			&txn.CreatedAt,
			&txn.UpdatedAt,
		)
	} else {
		err = s.client.QueryRowContext(ctx, query, id).Scan(
			&txn.ID,
			&txn.Concept,
			&txn.Description,
			&txn.Value,
			&txn.Date,
			&txn.Relevance,
			&txn.AccountID,
			&txn.CreatedAt,
			&txn.UpdatedAt,
		)
	}
	if err != nil {
		fmt.Printf("error: %v", err)
		return txn, err
	}

	return txn, nil
}

func (s *PGTransactionStore) GetAllTransactions(ctx context.Context, tx *sql.Tx) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var txns []models.Transaction

	query := `
            SELECT 
                id, concept, description, value, date, relevance, account_id, created_at, updated_at
            from transactions
    `

	var err error
	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.QueryContext(ctx, query)
	} else {
		rows, err = s.client.QueryContext(ctx, query)
	}
	if err != nil {
		return txns, err
	}
	defer rows.Close()

	for rows.Next() {
		var txn models.Transaction
		err = rows.Scan(
			&txn.ID,
			&txn.Concept,
			&txn.Description,
			&txn.Value,
			&txn.Date,
			&txn.Relevance,
			&txn.AccountID,
			&txn.CreatedAt,
			&txn.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		txns = append(txns, txn)
	}

	return txns, nil
}

func (s *PGTransactionStore) UpdateTransaction(ctx context.Context, tx *sql.Tx, id int, params *models.UpdateTransaction) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	now := time.Now()

	setClauses := []string{}
	args := []interface{}{}
	argID := 1

	if params.Concept != "" {
		setClauses = append(setClauses, fmt.Sprintf("concept = $%d", argID))
		args = append(args, params.Concept)
		argID++
	}
	if params.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", argID))
		args = append(args, params.Description)
		argID++
	}
	if !params.Date.IsZero() {
		setClauses = append(setClauses, fmt.Sprintf("date = $%d", argID))
		args = append(args, params.Date)
		argID++
	}
	if params.Relevance != 0 {
		setClauses = append(setClauses, fmt.Sprintf("relevance = $%d", argID))
		args = append(args, params.Relevance)
		argID++
	}
	if params.AccountID != 0 {
		setClauses = append(setClauses, fmt.Sprintf("account_id = $%d", argID))
		args = append(args, params.AccountID)
		argID++
	}
	if params.Value != 0 {
		setClauses = append(setClauses, fmt.Sprintf("value = $%d", argID))
		args = append(args, params.Value)
		argID++
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", argID))
	args = append(args, now)
	argID++

	args = append(args, id)

	query := fmt.Sprintf(`
        UPDATE transactions 
        SET %s
        WHERE id = $%d
    `, strings.Join(setClauses, ", "), argID)

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = s.client.ExecContext(ctx, query, args...)
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
