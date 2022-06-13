package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/linkshot.git/internal/config"
)

func Init(cfg *config.Config) (*sqlx.DB, error) {
	var dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Name, cfg.Postgres.User, cfg.Postgres.Password)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *sqlx.DB) error {
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
