package postgres

import (
	"backend/config"
	"backend/sorry"

	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	DB *sql.DB
}

func (p *Postgres) Open(cfg *config.DatabaseConfig) error {
	pgxPoolConfig, err := pgxpool.ParseConfig("postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.Port + "?sslmode=disable")
	if err != nil {
		return sorry.Err(err)
	}

	db := stdlib.OpenDB(*pgxPoolConfig.ConnConfig)
	if db == nil {
		return sorry.NewErr("failed to open database")
	}

	if err = db.Ping(); err != nil {
		return sorry.Err(err)
	}

	p.DB = db
	return nil
}

func (p *Postgres) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}

func (p *Postgres) NewTx() (any, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, sorry.Err(err)
	}

	return tx, nil
}
