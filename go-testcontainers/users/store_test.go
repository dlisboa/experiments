//go:build integration

package users

import (
	"demo/internal/testdb"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	testdb.Create()
	testdb.Migrate()
	testdb.Seed()

	t := time.Now()
	code := m.Run()
	fmt.Fprintf(os.Stdout, "===== tests took: %v\n", time.Since(t))

	testdb.Terminate()
	os.Exit(code)
}

func TestFoo(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded user
	john := "John Lennon"
	user, err := store.GetByName(t.Context(), john)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, john; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	george := "George Harrison"
	err = store.Insert(t.Context(), george)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), george)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	ringo := "Ringo Starr"
	_, err = store.GetByName(t.Context(), ringo)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}
}

func TestUserStoreInsert(t *testing.T) {
	t.Parallel()
	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}
	checkEmptyDB(t, store)
	testInsert(t, store)
}

func TestUserStoreInsert2(t *testing.T) {
	t.Parallel()
	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}
	checkEmptyDB(t, store)
	testInsert(t, store)
}

func TestUserStoreInsert3(t *testing.T) {
	t.Parallel()
	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}
	checkEmptyDB(t, store)
	testInsert(t, store)
}

func TestUserStoreInsert4(t *testing.T) {
	t.Parallel()
	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}
	checkEmptyDB(t, store)
	testInsert(t, store)
}

func testInsert(t *testing.T, store *Store) {
	err := store.Insert(t.Context(), "john")
	if err != nil {
		t.Fatal(err)
	}

	user, err := store.GetByName(t.Context(), "john")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := user.Name, "john"; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}
}

func checkEmptyDB(t *testing.T, store *Store) {
	_, err := store.GetByName(t.Context(), "john")
	if err == nil {
		t.Fatalf("there shouldn't be any previous users")
	}
}
