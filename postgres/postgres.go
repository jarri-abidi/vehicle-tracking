package postgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const dbName = "karma"

type Repository struct{ Conn *pgx.Conn }

func Migrate(migrationURL, connString string) error {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return errors.Wrap(err, "could not open postgres connection")
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{MultiStatementEnabled: true})
	if err != nil {
		return errors.Wrap(err, "could not create postgres driver instance")
	}

	migration, err := migrate.NewWithDatabaseInstance(migrationURL, dbName, driver)
	if err != nil {
		return errors.Wrap(err, "could not create migration instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "failed to run migrate up")
	}

	return nil
}
