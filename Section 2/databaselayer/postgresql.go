package databaselayer

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PQHandler struct {
	*SQLHandler
}

func NewPQHandler(connection string) (*PQHandler, error) {
	db, err := sql.Open("postgres", connection)
	return &PQHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
