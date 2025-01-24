package testdb

import (
	"context"
	"database/sql"
	"demo/internal/db"
	"log"
	"sync"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
)

var (
	sqldb     *sql.DB
	once      sync.Once
	container *postgres.PostgresContainer
)

func Create() {
	ctx := context.TODO()

	once.Do(func() {
		ctr, err := postgres.Run(
			ctx,
			"postgres:16-alpine",
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).
					WithStartupTimeout(5*time.Second)),
		)
		if err != nil {
			log.Fatal(err)
		}
		container = ctr

		conn, err := ctr.ConnectionString(ctx, "sslmode=disable")
		if err != nil {
			log.Fatalf("testdb: connection string: %v", err)
		}
		db, err := sql.Open("postgres", conn)
		if err != nil {
			log.Fatalf("testdb: cannot open db: %v", err)
		}
		err = db.Ping()
		if err != nil {
			log.Fatalf("testdb: cannot ping db: %v", err)
		}

		sqldb = db
	})
}

func Migrate() {
	if sqldb == nil {
		log.Fatal("testdb: called Migrate() with no container; call Create() first")
	}

	_, err := sqldb.Exec(`
	CREATE TABLE users (
		name text
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func Seed() {
	if sqldb == nil {
		log.Fatal("testdb: called Seed() with no container; call Create() first")
	}

	_, err := sqldb.Exec(`INSERT INTO users (name) VALUES ('John Lennon'), ('Paul McCarthy')`)
	if err != nil {
		log.Fatal(err)
	}
}

func Terminate() {
	if err := testcontainers.TerminateContainer(container); err != nil {
		log.Fatalf("testdb: cannot terminate container: %v", err)
	}
}

type dbtx struct {
	tx *sql.Tx
}

var _ db.DBTX = (&dbtx{})

func (t *dbtx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}

func (t *dbtx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return t.tx.PrepareContext(ctx, query)
}

func (t *dbtx) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return t.tx.QueryContext(ctx, query, args...)
}

func (t *dbtx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return t.tx.QueryRowContext(ctx, query, args...)
}

func (t *dbtx) StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt {
	return t.tx.StmtContext(ctx, stmt)
}

func (t *dbtx) Teardown() {
	err := t.tx.Rollback()
	if err != nil {
		log.Fatal("testdb: rollback:", err)
	}
}

func DB() *dbtx {
	if sqldb == nil {
		log.Fatal("testdb: called DB() with no container; call Create() first")
	}

	tx, err := sqldb.Begin()
	if err != nil {
		log.Fatal("testdb: cannot begin transaction:", err)
	}

	return &dbtx{tx}
}
