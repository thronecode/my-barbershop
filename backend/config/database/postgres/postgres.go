package postgres

import (
	"backend/config"
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func (p *Postgres) Open(cfg *config.DatabaseConfig) error {
	dsn := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + "/" + cfg.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	p.DB = db
	return nil
}

func (p *Postgres) Close() {
	if p.DB != nil {
		_ = p.DB.Close()
	}
}

func (p *Postgres) NewTx() (interface{}, error) {
	return p.DB.Begin()
}
