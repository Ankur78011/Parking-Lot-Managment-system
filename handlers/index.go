package handlers

import (
	"ankur/parkinlot/database"
	"database/sql"
)

type Apihandler struct {
	Db      *sql.DB
	Storage *database.Storage
}

func NewApiHandler(db *sql.DB) *Apihandler {
	return &Apihandler{Db: db, Storage: database.NewStorage(db)}
}
