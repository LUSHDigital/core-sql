package sqlexec

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Transactor defines a common set of methods for working with database transactions.
type Transactor interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// ExecTxContext executes a transaction with an associated context.
func ExecTxContext(ctx context.Context, trans Transactor, opt *sql.TxOptions, actions func(*sql.Tx) error) (err error) {
	if opt == nil {
		opt = &sql.TxOptions{
			ReadOnly:  false,
			Isolation: sql.LevelSerializable,
		}
	}
	var tx *sql.Tx
	if tx, err = trans.BeginTx(ctx, opt); err != nil {
		return err
	}

	defer func(tx *sql.Tx) {
		if r := recover(); r != nil {
			// Only need to log here because panic won't report whether
			// the rollback was successful or not.
			if txerr := tx.Rollback(); txerr != nil {
				log.Println("database transaction rollback error:", txerr)
			}

			log.Printf("database transaction rolled back transaction\n")
			panic(r)
		} else if err != nil {
			// If we run into issues rolling back, keep track of the error that
			// caused the issue and provide some context on the rollback failure.
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("database transaction error: %v rollback error: %v", err, rerr)
			}
		} else {
			if cerr := tx.Commit(); cerr != nil {
				err = fmt.Errorf("database transaction commit error: %v", cerr)
			}
		}
	}(tx)

	err = actions(tx)
	return err
}

// ExecTx closes over a transaction and automatically commits,
// or rollbacks depending on whether errors were encountered.
func ExecTx(trans Transactor, opt *sql.TxOptions, actions func(*sql.Tx) error) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ExecTxContext(ctx, trans, opt, actions)
}
