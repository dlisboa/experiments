Example of how to setup test containers with per-test transaction isolation
and parallelism.

Excluding the time to boot up and shutdown the Postgres image, runs 1000 tests
in about 500 ms. 10k tests take ~5s.

To generate tests: `go run ./tools/gentests -n 1000 -pkg users >./users/store_gen_test.go`

To run tests:

```shell
# -cache=1 makes sure tests are run and not cached
# not needed in normal usage
$ go test ./users -tags=integration -cache=1
```
