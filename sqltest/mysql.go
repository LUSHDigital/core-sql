package sqltest

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// MySQLTruncator represents a set of methods for truncating for MySQL specifically.
type MySQLTruncator struct {
	agent Agent
}

// MustTruncateAll will run TruncateAll and fail test if it's unsuccessful.
func (tr *MySQLTruncator) MustTruncateAll(t testing.TB) {
	mustTruncateAll(t, tr)
}

// TruncateAll will empty all tables in the database.
func (tr *MySQLTruncator) TruncateAll(t testing.TB) error {
	return truncateAll(t, tr, tr.agent)
}

// MustTruncateTables will run TruncateTables and will fail test if it can't.
func (tr *MySQLTruncator) MustTruncateTables(t testing.TB, tables ...string) {
	mustTruncateTables(t, tr, tables...)
}

// TruncateTables removes all content in the given tables.
func (tr *MySQLTruncator) TruncateTables(t testing.TB, tables ...string) error {
	const (
		setForeignKeysStmt = "SET FOREIGN_KEY_CHECKS=?"
	)
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
