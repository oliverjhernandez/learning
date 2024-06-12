package db

import (
	"context"

	"finance/types"
)

const (
	DBNAME  = "finance"
	DBURI   = "mongodb://localhost:27017"
	TDBNAME = "test-finance"
	TDBURI  = "mongodb://localhost:27017"
)

type Store struct {
	User   UserStore
	Tx     TransactionStore
	Credit CreditStore
}

type Dropper interface {
	Drop(ctx context.Context) error
}

type UserStore interface {
	GetUsers(ctx context.Context) ([]*types.User, error)
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	InsertUser(ctx context.Context, user *types.User) (*types.User, error)
	UpdateUser(ctx context.Context, filter Params, params *types.UpdateUserParams) error
	DeleteUser(ctx context.Context, id string) error
}

type TransactionStore interface {
	Dropper

	GetTransactions(ctx context.Context) ([]*types.Transaction, error)
	GetTransactionByID(ctx context.Context, id string) (*types.Transaction, error)
	InsertTransaction(ctx context.Context, tx *types.Transaction) (*types.Transaction, error)
	UpdateTransaction(ctx context.Context, filter Params, params *types.UpdateTransactionParams) error
	DeleteTransaction(ctx context.Context, id string) error
}

type CreditStore interface {
	GetCredits(ctx context.Context) ([]*types.Credit, error)
	GetCreditByID(ctx context.Context, id string) (*types.Credit, error)
	InsertCredit(ctx context.Context, cred *types.Credit) (*types.Credit, error)
	UpdateCredit(ctx context.Context, filter Params, update *types.UpdateCreditParams) error
	DeleteCreditByID(ctx context.Context, id string) error
}

type Params map[string]any
