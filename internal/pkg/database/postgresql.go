package database

import (
	"database/sql"
	"fmt"

	"github.com/natanaelrusli/segokuning-be/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitPostgres(cfg *config.Config) (*sql.DB, error) {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DbName,
		dbCfg.Port,
		dbCfg.Sslmode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
