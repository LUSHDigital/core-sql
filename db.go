package coresql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
)

// DB represents a wrapper for SQL DB providing extra methods.
type DB struct {
	*sql.DB
}

// Check will attempt to ping the database to see if the connection is still alive.
func (db *DB) Check() ([]string, bool) {
	if err := db.Ping(); err != nil {
		return []string{err.Error()}, false
	}
	return []string{"database connection ok"}, true
}

// Open will attempt to open a new database connection.
func Open(driverName, dsn string) (*DB, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	// see: https://github.com/go-sql-driver/mysql/issues/674
	db.SetMaxIdleConns(0)
	return &DB{db}, nil
}

// MustOpen will crash your program unless a database could be retrieved.
func MustOpen(driverName, dsn string) *DB {
	db, err := Open(driverName, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

// OpenWithMigrations opens a database connection with an associated migration instance.
func OpenWithMigrations(driverName, dsn, sourceURL string) (*DB, *migrate.Migrate, error) {
	migration, err := migrate.New(sourceURL, fmt.Sprintf("%s://%s", driverName, dsn))
	if err != nil {
		return nil, nil, err
	}
	database, err := Open(driverName, dsn)
	if err != nil {
		return nil, nil, err
	}
	return database, migration, nil
}

// MustOpenWithMigrations opens a database connection with an associated migration instance and crashes if unsuccessful.
func MustOpenWithMigrations(driverName, dsn, sourceURL string) (*DB, *migrate.Migrate) {
	database, migrations, err := OpenWithMigrations(driverName, dsn, sourceURL)
	if err != nil {
		log.Fatalln(err)
	}
	return database, migrations
}
