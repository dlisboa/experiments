package testdb

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
)

type testdb struct {
	tx *sql.Tx
}

func (t *testdb) Exec(query string, args ...any) (sql.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t *testdb) QueryRow(query string, args ...any) *sql.Row {
	return t.tx.QueryRow(query, args...)
}

func (t *testdb) Teardown() {
	err := t.tx.Rollback()
	if err != nil {
		log.Fatal("testdb: rollback:", err)
	}
}

var (
	sqldb     *sql.DB
	once      sync.Once
	container *postgres.PostgresContainer
)

func Create() {
	log.Println("testdb: creating db container")
	ctx := context.TODO()

	once.Do(func() {
		ctr, err := postgres.Run(
			ctx,
			"postgres:16-alpine",
			postgres.WithDatabase("mydb_test"),
			postgres.WithUsername("testuser"),
			postgres.WithPassword("testpassword"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).
					WithStartupTimeout(5*time.Second)),
		)
		if err != nil {
			log.Fatal(err)
		}

		container = ctr

		connStr, err := container.ConnectionString(ctx, "sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		pool, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		sqldb = pool
	})
}

func DB() *testdb {
	tx, err := sqldb.Begin()
	if err != nil {
		log.Fatal("testdb: cannot begin transaction:", err)
	}

	return &testdb{
		tx: tx,
	}
}

func Migrate() {
	if sqldb == nil {
		log.Fatal("testdb: called Migrate() with no container; call Create() first")
	}

	_, err := sqldb.Exec(`
	CREATE TABLE users (
		name text
	);`)

	if err != nil {
		log.Fatal(err)
	}
}

func Terminate(ctx context.Context) {
	if err := testcontainers.TerminateContainer(container); err != nil {
		log.Fatalf("testdb: failed to terminate container: %v", err)
	}
}
