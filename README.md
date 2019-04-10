# Core SQL
This package is used to wrap the SQL standard library to satisfy the [health checking server](https://github.com/LUSHDigital/core/tree/master/workers/readysrv#ready-server) in the [LUSH core service library](https://github.com/LUSHDigital/core)

## Examples

### Use in conjunction with [readysrv](https://github.com/LUSHDigital/core/tree/master/workers/readysrv)

```go
db, err := coresql.Open("mysql://localhost:3306/mydb")
if err != nil {
    log.Fatal(err)
}
readysrv.New(readysrv.Checks{
    "mysql": db,
})
```
