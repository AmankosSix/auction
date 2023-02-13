package schema

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func Migrator(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/schema",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		return err
	}
	return nil
}
