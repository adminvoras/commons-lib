package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Client interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Prepare(query string) (*sql.Stmt, error)
	Beginx() (*sqlx.Tx, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
}
