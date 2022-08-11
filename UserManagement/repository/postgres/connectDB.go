package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/akbarshoh/microOLX/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresDbname),
	)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"users",
		driver,
	)
	if err != nil {
		return nil, err
	}
	if err := m.Up(); !errors.Is(err, migrate.ErrNoChange) && err != nil {
		return nil, err
	}
	return db, nil
}
