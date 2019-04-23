package sqltest

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

const (
	setForeignKeysStmt = "SET FOREIGN_KEY_CHECKS=?"
	truncateStmtFmt    = "TRUNCATE TABLE %s"
	showTablesStmt     = "SHOW TABLES"
)

// Agent defines a common set of methods for executing stored procedures, statements or queries.
type Agent interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// NewTruncator sets up a new truncator.
func NewTruncator(agent Agent) *Truncator {
	return &Truncator{agent}
}

// Truncator represents a common set of methods for truncating a database during testing.
type Truncator struct {
	agent Agent
}

// MustTruncateAll will run TruncateAll and fail test if it can't.
func (tr *Truncator) MustTruncateAll(t testing.TB) {
	if err := tr.TruncateAll(t); err != nil {
		t.Error(err)
	}
}

// TruncateAll will empty all tables in the database.
func (tr *Truncator) TruncateAll(t testing.TB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := tr.agent.QueryContext(ctx, showTablesStmt)
	if err != nil {
		return err
	}
	var tables []string
	for rows.Next() {
		var table string
		rows.Scan(&table)
		switch table {
		case "schema_migrations":
		default:
			tables = append(tables, table)
		}
	}
	return tr.TruncateTables(t, tables...)
}

// MustTruncateTables will run TruncateTables and will fail test if it can't.
func (tr *Truncator) MustTruncateTables(t testing.TB) {
	if err := tr.TruncateTables(t); err != nil {
		t.Error(err)
	}
}

// TruncateTables removes all content in the given tables.
func (tr *Truncator) TruncateTables(t testing.TB, tables ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := tr.agent.ExecContext(ctx, setForeignKeysStmt, false); err != nil {
		return err
	}
	for _, table := range tables {
		if _, err := tr.agent.ExecContext(ctx, fmt.Sprintf(truncateStmtFmt, table)); err != nil {
			return err
		}
	}
	if _, err := tr.agent.ExecContext(ctx, setForeignKeysStmt, true); err != nil {
		return err
	}
	return nil
}
