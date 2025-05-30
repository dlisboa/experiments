// Generated with tools/gentests. DO NOT EDIT.

//go:build integration

package users

import (
	"testing"

	"demo/internal/testdb"
)

func Test0(t *testing.T) {
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

func Test1(t *testing.T) {
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

func Test2(t *testing.T) {
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

func Test3(t *testing.T) {
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

func Test4(t *testing.T) {
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

func Test5(t *testing.T) {
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

func Test6(t *testing.T) {
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

func Test7(t *testing.T) {
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

func Test8(t *testing.T) {
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

func Test9(t *testing.T) {
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
