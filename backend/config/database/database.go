package database

import (
	"backend/config"
	"backend/config/database/postgres"
	"backend/sorry"

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
		if db.ReadOnly {
			db.Driver += roSuffix
		}
		if _, set := connections[db.Driver]; set {
			if err := connections[db.Driver].Open(&db); err != nil {
				return sorry.Err(err)
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
		return nil, sorry.Err(err)
	}

	t.postgres = pgsql.(*sql.Tx)
	t.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(t)

	return t, nil
}

// DBTransaction methods for executing queries

// Commit commits pending transactions for all databases open
func (t *DBTransaction) Commit() error {
	err := t.postgres.Commit()
	if err != nil {
		return sorry.Err(err)
	}

	return nil
}

// Rollback rollback pending transaction for all databases open
func (t *DBTransaction) Rollback() {
	_ = t.postgres.Rollback()
}

// Exec executes a query that doesn't return rows.
func (t *DBTransaction) Exec(query string, args ...any) (sql.Result, error) {
	result, err := t.postgres.Exec(query, args...)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return result, nil
}

// Query executes a query that returns rows.
func (t *DBTransaction) Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := t.postgres.Query(query, args...)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return rows, nil
}

// QueryRow executes a query that is expected to return at most one row.
func (t *DBTransaction) QueryRow(query string, args ...any) squirrel.RowScanner {
	return t.postgres.QueryRow(query, args...)
}
