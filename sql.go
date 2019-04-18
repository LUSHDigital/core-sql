package coresql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

// Pinger defines a common set of methods for pinging an SQL database.
type Pinger interface {
	PingContext(ctx context.Context) error
	Ping() error
}

// Connector defines a common set of methods for dealing with the connection to an SQL datbase.
type Connector interface {
	Driver() driver.Driver
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	SetConnMaxLifetime(d time.Duration)
	Conn(ctx context.Context) (*sql.Conn, error)
	Stats() sql.DBStats
}

// Preparer defines a common set of methods for preparing statements in an SQL database.
type Preparer interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
}

// Transactor defines a common set of methods for working with database transactions.
type Transactor interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
}

// Executor defines a common set of methods for executing stored procedures, statements or queries.
type Executor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// Querier defines a common set of methods for querying an SQL database.
type Querier interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Agent defines a common set of methods for interacting with the data in an SQL database.
type Agent interface {
	Preparer
	Transactor
	Executor
	Querier
}

// Operator defines a common set of methods for operating a connection with an SQL database.
type Operator interface {
	Pinger
	Connector
}
