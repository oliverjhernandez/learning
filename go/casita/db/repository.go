package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func NewStore() (*Store, *sql.DB, error) {
	client := connectSQL()
	return &Store{
		DB:           client,
		UserStore:    NewPGUserStore(client),
		TxnStore:     NewPGTransactionStore(client),
		CreditStore:  NewPGCreditStore(client),
		AccountStore: NewPGAccountStore(client),
	}, client, nil
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

func connectSQL() *sql.DB {
	// TODO: create config getter
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "casita"
	dbUser := "postgres"
	dbPasswd := "secret"
	dbSSL := "disable"

	// Connecto to DB
	log.Println("Connecting to dabase")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPasswd, dbSSL)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
