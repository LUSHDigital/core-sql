package sqltest

import (
	"context"
	"database/sql"
	"testing"
)

// Agent defines a common set of methods for executing stored procedures, statements or queries.
type Agent interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// Truncator represents a common set of methods for truncating tables in a database.
type Truncator interface {
	TruncateAll(t testing.TB) error
	TruncateTables(t testing.TB, tables ...string) error
	MustTruncateAll(t testing.TB)
	MustTruncateTables(t testing.TB, tables ...string)
}

// NewTruncator opens a new truncator for a database.
func NewTruncator(driverName string, agent Agent) Truncator {
	switch driverName {
	case "mysql":
		return &MySQLTruncator{agent}
	default:
		return &GenericTruncator{agent}
	}
}
