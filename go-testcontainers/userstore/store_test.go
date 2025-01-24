package userstore

import (
	"context"
	"demo/internal/testdb"
	"testing"
)

func TestMain(m *testing.M) {
	testdb.Create()
	testdb.Migrate()
	m.Run()
	testdb.Terminate(context.TODO())
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
	err := store.Insert("john")
	if err != nil {
		t.Fatal(err)
	}

	user, err := store.GetByName("john")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := user.Name, "john"; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}
}

func checkEmptyDB(t *testing.T, store *Store) {
	_, err := store.GetByName("john")
	if err == nil {
		t.Fatalf("there shouldn't be any previous users")
	}
}
