[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/LUSHDigital/core-sql/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/LUSHDigital/core-sql)](https://goreportcard.com/report/github.com/LUSHDigital/core-sql) [![Build Status](https://travis-ci.org/LUSHDigital/core-sql.svg?branch=master)](https://travis-ci.org/LUSHDigital/core-sql)
[![GoDoc](https://godoc.org/github.com/LUSHDigital/core-sql?status.svg)](https://godoc.org/github.com/LUSHDigital/core-sql)


# Core SQL
This package is used to wrap the SQL standard library to satisfy the [health checking server](https://github.com/LUSHDigital/core/tree/master/workers/readysrv#ready-server) in the [LUSH core service library](https://github.com/LUSHDigital/core). We use [golang-migrate/migrate](github.com/golang-migrate/migrate) version 4 for migrations and treats it as a direct dependency.

## Examples

### Use in conjunction with [readysrv](https://github.com/LUSHDigital/core/tree/master/workers/readysrv)

```go
database := coresql.MustOpen("mysql", "tcp(localhost:3306)/mydb")
readysrv.New(readysrv.Checks{
    "mysql": database,
})
```

### Ensure database is migrated up

```go
_, migrations := coresql.MustOpenWithMigrations("mysql", "tcp(localhost:3306)/mydb", "file://path/to/migrations")
coresql.MustMigrateUp(migrations)
```

### Handle migration arguments
You can force your application to respect migration command line arguments:

- `migrate up`: attempt to migrate up
- `migrate down`: attempt to migrate down

```go
_, migrations := coresql.MustOpenWithMigrations("mysql", "tcp(localhost:3306)/mydb", "file://path/to/migrations")
coresql.HandleMigrationArgs(migrations)
```
