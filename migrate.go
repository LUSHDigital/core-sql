package coresql

import (
	"flag"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

// Migrate represents functionality for
type Migrate interface {
	Up() error
	Down() error
}

// HandleMigrationArgs should be invoked to handle command line arguments for running migrations.
func HandleMigrationArgs(mig Migrate) {
	flag.Parse()
	switch flag.Arg(0) {
	case "migrate":
		switch flag.Arg(1) {
		case "up":
			if err := mig.Up(); err != nil {
				log.Fatalln(err)
			}
		case "down":
			if err := mig.Down(); err != nil {
				log.Fatalln(err)
			}
		default:
			log.Fatalln("migrate needs a sub-command")
		}
		os.Exit(0)
		return
	}
}

// MustMigrateUp will attempt to migrate your database up.
func MustMigrateUp(mig Migrate) {
	if err := mig.Up(); err != nil {
		switch err {
		case migrate.ErrNoChange:
			log.Println("database migration:", "no changes")
		default:
			log.Fatalln("database migration:", err)
		}
	}
}
