package coresql

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/xo/dburl"
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
func Open(databaseURL string) (*DB, error) {
	database, err := dburl.Open(databaseURL)
	if err != nil {
		return nil, err
	}
	return &DB{database}, nil
}

// MustOpen will crash your program unless a database could be retrieved.
func MustOpen(databaseURL string) *DB {
	db, err := Open(databaseURL)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

// OpenWithMigrations opens a database connection with an associated migration instance.
func OpenWithMigrations(databaseURL, migrationsURL string) (*DB, *migrate.Migrate, error) {
	migration, err := migrate.New(migrationsURL, databaseURL)
	if err != nil {
		return nil, nil, err
	}
	database, err := Open(databaseURL)
	if err != nil {
		return nil, nil, err
	}
	return database, migration, nil
}

// MustOpenWithMigrations opens a database connection with an associated migration instance and crashes if unsuccessful.
func MustOpenWithMigrations(databaseURL, migrationsURL string) (*DB, *migrate.Migrate) {
	database, migrations, err := OpenWithMigrations(databaseURL, migrationsURL)
	if err != nil {
		log.Fatalln(err)
	}
	return database, migrations
}
