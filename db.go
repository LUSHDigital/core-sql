package coresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	// Since we are most likely going to be only retrieving migrations from file source,
	// it's prudent that we include this side effect inside of this package and not
	// having to import it inside each and every project.
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	waitMaxTries = 60
	waitTimeout  = 5 * time.Second
	waitCooldown = 10 * time.Millisecond
)

var errTimeout = fmt.Errorf("could not connect to database: timed out")

// DB represents a wrapper for SQL DB providing extra methods.
// DB satisfies the sql.DB interface.
type DB struct {
	*sql.DB
	hist prometheus.Histogram
}

// new creates a new DB, wrapping the provided sql.DB and adding metrics.
func newDB(driverName string, db *sql.DB) *DB {
	hist := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:        "database_query_roundtrip_duration_seconds",
		Help:        "A histogram of roundtrip query times for this database.",
		ConstLabels: map[string]string{"driver": driverName},
		// Divide buckets by doubling the threshold, in ms: 5, 10, 20, 40, 80, 160, 320.
		Buckets: prometheus.ExponentialBuckets(.05, 2, 7),
	})

	return &DB{
		DB:   db,
		hist: hist,
	}
}

func (db *DB) observe(start time.Time) {
	db.hist.Observe(float64(time.Since(start).Milliseconds() / 1000))
}

// Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.
// The duration of the call to the underlying database is added to the metrics for this DB.
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	defer db.observe(start)
	return db.DB.Exec(query, args...)
}

// ExecContext executes a query without returning any rows. The args are for any placeholder parameters in the query.
// The duration of the call to the underlying database is added to the metrics for this DB.
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	defer db.observe(start)
	return db.DB.ExecContext(ctx, query, args...)
}

// Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
// The duration of the call to the underlying database is added to the metrics for this DB.
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	defer db.observe(start)
	return db.DB.Query(query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
// The duration of the call to the underlying database is added to the metrics for this DB.
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	defer db.observe(start)
	return db.DB.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row. QueryRow always returns a non-nil value.
// Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards the rest.
// The duration of the call to the underlying database is added to the metrics for this DB.
func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	start := time.Now()
	defer db.observe(start)
	return db.DB.QueryRow(query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row. QueryRowContext always returns a non-nil value.
// Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards the rest.
// The duration of the call to the underlying database is added to the metrics for this DB.
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	start := time.Now()
	defer db.observe(start)
	return db.DB.QueryRowContext(ctx, query, args...)
}
func (db *DB) Check() ([]string, bool) {
	if err := db.Ping(); err != nil {
		return []string{err.Error()}, false
	}
	return []string{"database connection ok"}, true
}

// MustWait will call the Wait method and crash if it cant be performed.
func (db *DB) MustWait() {
	if err := db.Wait(); err != nil {
		log.Fatal(err)
	}
}

// Wait will attempt to connect to a database and block until it connects.
func (db *DB) Wait() error {
	ctx, cancel := context.WithTimeout(context.Background(), waitTimeout)
	defer cancel()

	tries := 0
	doneC := make(chan struct{}, 1)
	sem := make(chan struct{}, 1)

	ping := func(ctx context.Context) {
		err := db.PingContext(ctx)
		if err == nil {
			doneC <- struct{}{}
			return
		}
		time.Sleep(waitCooldown)
		tries++
		<-sem
	}

	for {
		select {
		case sem <- struct{}{}:
			if tries >= waitMaxTries {
				return fmt.Errorf("could not connect to datavase: attempt limit (%d) exceeded", waitMaxTries)
			}
			go ping(ctx)
		case <-ctx.Done():
			return fmt.Errorf("could not connect to database: timeout")
		case <-doneC:
			return nil
		}
	}
}

// Open will attempt to open a new database connection.
func Open(driverName, dsn string) (*DB, error) {
	switch driverName {
	case "mysql":
	default:
		dsn = fmt.Sprintf("%s://%s", driverName, dsn)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbC := make(chan *sql.DB, 1)
	errC := make(chan error, 1)

	go func(driverName, dsn string) {
		db, err := sql.Open(driverName, dsn)
		if err != nil {
			errC <- err
			return
		}
		dbC <- db
	}(driverName, dsn)

	select {
	case db := <-dbC:
		// see: https://github.com/go-sql-driver/mysql/issues/674
		db.SetMaxIdleConns(0)
		return newDB(driverName, db), nil
	case err := <-errC:
		return nil, err
	case <-ctx.Done():
		return nil, errTimeout
	}
}

// MustOpen will crash your program unless a database could be retrieved.
func MustOpen(driverName, dsn string) *DB {
	db, err := Open(driverName, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

// OpenMigration will open a migration instance.
func OpenMigration(driverName, dsn, sourceURL string) (*migrate.Migrate, error) {
	dsn = fmt.Sprintf("%s://%s", driverName, dsn)
	migration, err := migrate.New(sourceURL, dsn)
	if err != nil {
		return nil, err
	}
	return migration, nil
}

// MustOpenMigration will open a migration instance with and crashes if unsuccessful.
func MustOpenMigration(driverName, dsn, sourceURL string) *migrate.Migrate {
	migration, err := OpenMigration(driverName, dsn, sourceURL)
	if err != nil {
		log.Fatalln(err)
	}
	return migration
}

// OpenWithMigrations opens a database connection with an associated migration instance.
func OpenWithMigrations(driverName, dsn, sourceURL string) (*DB, *migrate.Migrate, error) {
	migration, err := OpenMigration(driverName, dsn, sourceURL)
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
