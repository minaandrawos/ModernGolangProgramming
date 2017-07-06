package databaselayer

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLHandler struct {
	*SQLHandler
}

func NewMySQLHandler(connection string) (*MySQLHandler, error) {
	db, err := sql.Open("mysql", connection)
	return &MySQLHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
