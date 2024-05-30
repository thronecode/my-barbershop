package database

import (
	"backend/config"
	"backend/config/database/postgres"

	"database/sql"

	"github.com/Masterminds/squirrel"
)

var (
	connections map[string]Database
	roSuffix    = "-ro"
)

// DBTransaction used to aggregate transactions for the database
type DBTransaction struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
}

// OpenConnections opens connections to the database
func OpenConnections() error {
	initConnectionsMap()

	conf := config.GetConfig()

	for _, db := range conf.Databases {
		if _, set := connections[db.Driver]; set {
			if err := connections[db.Driver].Open(&db); err != nil {
				return err
			}
		}
	}

	return nil
}

// CloseConnections closes all database connections
func CloseConnections() {
	for _, conn := range connections {
		conn.Close()
	}
}

func initConnectionsMap() {
	if connections != nil {
		return
	}
	connections = make(map[string]Database)
	connections["postgres"] = &postgres.Postgres{}
	connections["postgres"+roSuffix] = &postgres.Postgres{}
}

// NewTransaction starts a new database transaction
func NewTransaction(readOnly bool) (*DBTransaction, error) {
	t := &DBTransaction{}

	db := "postgres"
	if readOnly {
		db += roSuffix
	}

	pgsql, err := connections[db].NewTx()
	if err != nil {
		return nil, err
	}

	t.postgres = pgsql.(*sql.Tx)
	t.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(t)

	return t, nil
}

// DBTransaction methods for executing queries

// Exec executes a query that doesn't return rows.
func (t *DBTransaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.postgres.Exec(query, args...)
}

// Query executes a query that returns rows.
func (t *DBTransaction) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return t.postgres.Query(query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
func (t *DBTransaction) QueryRow(query string, args ...interface{}) squirrel.RowScanner {
	return t.postgres.QueryRow(query, args...)
}
