package database

import "backend/config"

// Database interface for multiple databases
type Database interface {
	Open(*config.DatabaseConfig) error
	Close()
	NewTx() (any, error)
}

// Transaction interface for multiple databases
type Transaction interface {
	Commit() error
	Rollback() error
}
