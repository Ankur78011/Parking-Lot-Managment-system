package database

import (
	"database/sql"
)

type Storage struct {
	Db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{Db: db}
}
