package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Vacancies   VacancyModel
	Users       UserModel
	Tokens      TokenModel
	Permissions PermissionModel
	Subscribers SubscriberModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Vacancies:   VacancyModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Subscribers: SubscriberModel{DB: db},
	}
}
