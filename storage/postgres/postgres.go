package postgres

import (
	"github.com/jmoiron/sqlx"
)

type crudRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *crudRepo {
	return &crudRepo{db: db}
}
