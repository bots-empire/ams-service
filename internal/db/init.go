package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/bots-empire/ams-service/db"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

const dbDriver = "postgres"

func InitDataBase(ctx context.Context, cfg *pgxpool.Config) (*pgxpool.Pool, error) {
	dataBase, err := sql.Open(dbDriver, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"psql-ams", 5432, "ams-user", "26538hsvgn8p", "ams-service"))
	if err != nil {
		log.Fatalf("Failed open database: %s\n", err.Error())
	}

	goose.SetBaseFS(db.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(dataBase, "migrations"); err != nil {
		panic(err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create conn pool")
	}

	return pool, nil
}
