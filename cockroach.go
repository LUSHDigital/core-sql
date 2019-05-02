package coresql

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
)

const (
	// DefaultCockroachURL is the default url to a CockroachDB database.
	DefaultCockroachURL = "tcp(127.0.0.1:26257)/defaultdb"
)

// CockroachURLFromEnv tries to retrieve the cockroach url from the environment.
func CockroachURLFromEnv() string {
	url := os.Getenv("COCKROACH_URL")
	if url == "" {
		url = DefaultCockroachURL
	}
	return url
}

// OpenCockroachWithMigration opens a cockroach database connection with an associated migration instance.
// This assumes you use a postgres driver like https://github.com/lib/pq to interact with your postgres database.
func OpenCockroachWithMigration(dsn, sourceURL string) (*DB, *migrate.Migrate, error) {
	const (
		migDriver = "cockroach"
		dbDriver  = "postgres"
	)
	migration, err := OpenMigration(migDriver, dsn, sourceURL)
	if err != nil {
		return nil, nil, err
	}
	database, err := Open(dbDriver, dsn)
	if err != nil {
		return nil, nil, err
	}
	return database, migration, nil
}
