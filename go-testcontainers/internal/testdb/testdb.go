package testdb

import (
	"context"
	"database/sql"
	"demo/internal/db"
	"fmt"
	"log"
	"sync"

	"github.com/testcontainers/testcontainers-go"
	// "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
)

var (
	sqldb     *sql.DB
	once      sync.Once
	container testcontainers.Container
)

func Create() {
	ctx := context.TODO()

	once.Do(func() {
		req := testcontainers.ContainerRequest{
			Name:         "mytest_postgres",
			Image:        "postgres:16-alpine",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     "testuser",
				"POSTGRES_PASSWORD": "testpassword",
				"POSTGRES_DB":       "testdb",
			},
			WaitingFor: wait.ForLog("database system is ready to accept connections").WithStartupTimeout(30 * 1000000000),
		}

		ctr, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
			Reuse:            true,
		})

		if err != nil {
			log.Fatal(err)
		}
		container = ctr

		// Fetch host and port
		host, err := container.Host(ctx)
		if err != nil {
			log.Fatal(err)
		}

		port, err := container.MappedPort(ctx, "5432")
		if err != nil {
			log.Fatal(err)
		}

		// Build connection string
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port.Port(),
			"testuser",
			"testpassword",
			"testdb",
		)

		db, err := sql.Open("postgres", dsn)
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
	);

	CREATE table movies (
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

	_, err := sqldb.Exec(`
		INSERT INTO users (name) VALUES ('John Lennon'), ('Paul McCarthy');
		INSERT INTO movies (name) VALUES ('Titanic'), ('Central Station');
	`)
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
