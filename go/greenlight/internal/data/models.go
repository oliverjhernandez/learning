package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// type Models struct {
// 	Movies interface {
// 		Insert(movie *Movie) error
// 		Get(id int64) (*Movie, error)
// 		GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error)
// 		Update(movie *Movie) error
// 		Delete(id int64) error
// 	}
// }

type Models struct {
	Movies      MovieModel
	Users       UserModel
	Tokens      TokenModel
	Permissions PermissionModel
	Accounts    AccountsModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Users:       UserModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Accounts:    AccountsModel{DB: db},
	}
}
