package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Store struct {
	*sql.DB
	UserStore    UserStore
	AccountStore AccountStore
	TxnStore     TransactionStore
	CreditStore  CreditStore
}

type DBRepo interface {
	Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error
	Drop(ctx context.Context, dbname string) error
}

type DBParams struct {
	Host   string
	Port   string
	Name   string
	User   string
	Passwd string
	SSL    string

	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type DBCfg struct {
	Env     string
	Port    int
	DB      DBParams
	Limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

func (s *Store) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
	tx, err := s.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() error {
		if err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return nil
	}()

	if err := operation(ctx, tx); err != nil {
		return err
	}

	return nil
}

func (s *Store) Drop(ctx context.Context, dbname string) error {
	if err := s.Close(); err != nil {
		return err
	}

	return nil
}

func ConnectSQL(dbParams DBParams) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbParams.Host,
		dbParams.Port,
		dbParams.Name,
		dbParams.User,
		dbParams.Passwd,
		dbParams.SSL)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
