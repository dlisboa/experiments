Example of how to setup test containers with per-test transaction isolation
and parallelism.

Excluding the time to boot up and shutdown the Postgres image, runs 1000
integration tests in about 500 ms. 10k tests take ~5s.

To generate tests: `go run ./tools/gentests -n 1000 -pkg users >./users/store_gen_test.go`

To run tests:

```shell
# -cache=1 makes sure tests are run and not cached
# not needed in normal usage
$ go test ./users -tags=integration -cache=1
```

See files:

* `internal/testdb/testdb.go` - create/migrate/seed and terminate Postgres container
* `internal/db/db.go` - a `DBTX` interface, similar to what
  [sqlc](https://github.com/sqlc-dev/sqlc) creates. Allows
  models/stores/repositories to think it's using a *sql.DB but it's actually
  inside a transaction
* `users/store.go` - a "user store" pattern, what we're testing
* `users/store_test.go` - implements `TestMain` that uses `testdb` to setup Postgres
* `users/store_gen_test.go` - the actual tests, generated
