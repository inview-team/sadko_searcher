package db

import (
	"context"
	"fmt"
	config2 "src/config"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Host     string
	Port     int64
	User     string
	Password string
	DbName   string
	SSLMode  string
}

func NewPsqlPoolConnection(conf config2.AppConfig) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		conf.Postgres.User,
		conf.Postgres.Password,
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.DbName,
		conf.Postgres.SSLMode,
	)
	dbURL = fmt.Sprintf("%s&pool_max_conns=%d", dbURL, 5)
	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}
