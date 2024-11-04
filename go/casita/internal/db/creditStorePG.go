package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type CreditStore interface {
	InsertCredit(ctx context.Context, tx *sql.Tx, credit *Credit) (*Credit, error)
	GetCreditByID(ctx context.Context, tx *sql.Tx, id int) (*Credit, error)
	GetAllCredits(ctx context.Context, tx *sql.Tx) ([]*Credit, error)
	UpdateCredit(ctx context.Context, tx *sql.Tx, id int, params *UpdateCredit) (*Credit, error)
	DeleteCreditByID(ctx context.Context, tx *sql.Tx, id int) error
}

type PGCreditStore struct {
	client *sql.DB
}

func NewPGCreditStore(client *sql.DB) *PGCreditStore {
	return &PGCreditStore{
		client: client,
	}
}

func (s *PGCreditStore) InsertCredit(ctx context.Context, tx *sql.Tx, credit *Credit) (*Credit, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var newID int

	query := `
            INSERT INTO credits
                (closing_day, due_day, identifier, entity, type, rate, total, user_id, installments, created_at, updated_at)
            VALUES
                ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
            RETURNING id
    `

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query,
			credit.ClosingDay,
			credit.DueDay,
			credit.Identifier,
			credit.Entity,
			credit.Type,
			credit.Rate,
			credit.Total,
			credit.UserID,
			credit.Installments,
			time.Now(),
			time.Now(),
		).Scan(&newID)
	} else {
		err = s.client.QueryRowContext(ctx, query,
			credit.ClosingDay,
			credit.DueDay,
			credit.Identifier,
			credit.Entity,
			credit.Type,
			credit.Rate,
			credit.Total,
			credit.UserID,
			credit.Installments,
			time.Now(),
			time.Now(),
		).Scan(&newID)
	}
	if err != nil {
		return nil, err
	}

	cred, err := s.GetCreditByID(ctx, nil, newID)
	if err != nil {
		return nil, err
	}

	return cred, nil
}

func (s *PGCreditStore) GetCreditByID(ctx context.Context, tx *sql.Tx, id int) (*Credit, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var credit Credit

	query := `
            SELECT 
              id, closing_day, due_day, identifier, entity, type, rate, total, installments, created_at, updated_at
            FROM credits
            WHERE id=$1
    `

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(
			&credit.ID,
			&credit.ClosingDay,
			&credit.DueDay,
			&credit.Identifier,
			&credit.Entity,
			&credit.Type,
			&credit.Rate,
			&credit.Total,
			&credit.Installments,
			&credit.CreatedAt,
			&credit.UpdatedAt,
		)
	} else {
		err = s.client.QueryRowContext(ctx, query, id).Scan(
			&credit.ID,
			&credit.ClosingDay,
			&credit.DueDay,
			&credit.Identifier,
			&credit.Entity,
			&credit.Type,
			&credit.Rate,
			&credit.Total,
			&credit.Installments,
			&credit.CreatedAt,
			&credit.UpdatedAt,
		)
	}
	if err != nil {
		return nil, err
	}

	return &credit, nil
}

func (s *PGCreditStore) GetAllCredits(ctx context.Context, tx *sql.Tx) ([]*Credit, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var credits []*Credit

	query := `
            SELECT 
              id, closing_day, due_day, identifier, entity, type, rate, total, installments, created_at, updated_at
            from credits`

	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryContext(ctx, query)
	} else {
		rows, err = s.client.QueryContext(ctx, query)
	}
	if err != nil {
		return credits, err
	}
	defer rows.Close()

	for rows.Next() {
		var credit Credit
		err := rows.Scan(
			&credit.ID,
			&credit.ClosingDay,
			&credit.DueDay,
			&credit.Identifier,
			&credit.Entity,
			&credit.Type,
			&credit.Rate,
			&credit.Total,
			&credit.Installments,
			&credit.CreatedAt,
			&credit.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		credits = append(credits, &credit)
	}

	return credits, nil
}

func (s *PGCreditStore) UpdateCredit(ctx context.Context, tx *sql.Tx, id int, params *UpdateCredit) (*Credit, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	now := time.Now()
	var newID int

	setClauses := []string{}
	args := []interface{}{}
	argID := 1

	if params.ClosingDay != 0 {
		setClauses = append(setClauses, fmt.Sprintf("closing_day = $%d", argID))
		args = append(args, params.ClosingDay)
		argID++
	}
	if params.DueDay != 0 {
		setClauses = append(setClauses, fmt.Sprintf("due_day = $%d", argID))
		args = append(args, params.DueDay)
		argID++
	}
	if params.Identifier != "" {
		setClauses = append(setClauses, fmt.Sprintf("identifier = $%d", argID))
		args = append(args, params.Identifier)
		argID++
	}
	if params.Entity != 0 {
		setClauses = append(setClauses, fmt.Sprintf("entity = $%d", argID))
		args = append(args, params.Entity)
		argID++
	}
	if params.Type != 0 {
		setClauses = append(setClauses, fmt.Sprintf("type = $%d", argID))
		args = append(args, params.Type)
		argID++
	}
	if params.Rate != 0 {
		setClauses = append(setClauses, fmt.Sprintf("rate = $%d", argID))
		args = append(args, params.Rate)
		argID++
	}
	if params.Total != 0 {
		setClauses = append(setClauses, fmt.Sprintf("total = $%d", argID))
		args = append(args, params.Total)
		argID++
	}
	if params.Installments != 0 {
		setClauses = append(setClauses, fmt.Sprintf("installments = $%d", argID))
		args = append(args, params.Installments)
		argID++
	}
	if params.UserID != 0 {
		setClauses = append(setClauses, fmt.Sprintf("user_id = $%d", argID))
		args = append(args, params.UserID)
		argID++
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", argID))
	args = append(args, now)
	argID++

	args = append(args, id)

	query := fmt.Sprintf(`
        UPDATE credits 
        SET %s
        WHERE id = $%d
    `, strings.Join(setClauses, ", "), argID)

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, args...).Scan(&newID)
	} else {
		err = s.client.QueryRowContext(ctx, query, args...).Scan(&newID)
	}
	if err != nil {
		return nil, err
	}

	cred, err := s.GetCreditByID(ctx, nil, newID)
	if err != nil {
		return nil, err
	}

	return cred, nil
}

func (s *PGCreditStore) DeleteCreditByID(ctx context.Context, tx *sql.Tx, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `DELETE FROM credits WHERE id = $1`

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
