package storage

import (
	"github.com/NajmiddinAbdulhakim/iman/crud-service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	CRUD() repo.CRUDStorage
}

type storagePg struct {
	db   *sqlx.DB
	crud repo.CRUDStorage
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{db: db}
}

func (s storagePg) CRUD() repo.CRUDStorage {
	return s.crud
}
