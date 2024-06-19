package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"finance/models"
)

type CreditStore interface {
	// TODO: Inject context instead of creating it in each method
	InsertCredit(credit *models.Credit) (int, error)
	GetCreditByID(id int) (models.Credit, error)
	GetAllCredits() ([]models.Credit, error)
	UpdateACreditsName(id int, firstName, lastName string) error
	DeleteCreditByID(id int) error
}

type PGCreditStore struct {
	client *sql.DB
}

func NewPGCreditStore(client *sql.DB) *PGCreditStore {
	return &PGCreditStore{
		client: client,
	}
}

func (s *PGCreditStore) InsertCredit(credit *models.Credit) (int, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	var newID int

	query := `
            INSERT into credits 
                (closing_date, due_date, identifier, entity, type, rate, total, installments, created_at, updated_at)
            values 
                ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
            returning id`

	err := s.client.QueryRowContext(ctx, query,
		credit.ClosingDate,
		credit.DueDate,
		credit.Identifier,
		credit.Entity,
		credit.Type,
		credit.Rate,
		credit.Total,
		credit.Installments,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (s *PGCreditStore) GetCreditByID(id int) (models.Credit, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	var credit models.Credit

	query := `
            SELECT 
              (closing_date, due_date, identifier, entity, type, rate, total, installments, created_at, updated_at)
            from credits
            WHERE
            id=$1`

	err := s.client.QueryRowContext(ctx, query, id).Scan(
		&credit.ID,
		&credit.ClosingDate,
		&credit.DueDate,
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
		return credit, err
	}

	return credit, nil
}

func (s *PGCreditStore) GetAllCredits() ([]models.Credit, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	var credits []models.Credit

	query := `
            SELECT 
              (closing_date, due_date, identifier, entity, type, rate, total, installments, created_at, updated_at)
            from credits`

	rows, err := s.client.QueryContext(ctx, query)
	if err != nil {
		return credits, err
	}

	for rows.Next() {
		var credit models.Credit
		err := s.client.QueryRowContext(ctx, query).Scan(
			&credit.ID,
			&credit.ClosingDate,
			&credit.DueDate,
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
			return credits, err
		}
	}

	return credits, nil
}

func (s *PGCreditStore) UpdateCredit(id int, params *models.UpdateCreditParams) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	now := time.Now()

	query := fmt.Sprintf(`
        UPDATE credits set 
            closing_date = COALESCE($1, closing_date),
            due_date = COALESCE($2, due_date),
            identifier = COALESCE($3, identifier),
            entity = COALESCE($4, entity),
            type = COALESCE($5, type),
            rate = COALESCE($6, rate),
            total = COALESCE($7, total),
            installments = COALESCE($8, installments),
            updated_at = %s
        WHERE
            id = $9
    `, now)

	_, err := s.client.ExecContext(
		ctx,
		query,
		params.ClosingDate,
		params.DueDate,
		params.Identifier,
		params.Entity,
		params.Type,
		params.Rate,
		params.Total,
		params.Installments,
		id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PGCreditStore) DeleteCreditByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM credits WHERE id = $1"

	_, err := s.client.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
