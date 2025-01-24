// Generated with tools/gentests. DO NOT EDIT.

//go:build integration

package movies

import (
	"testing"

	"demo/internal/testdb"
)

func Test_movies0(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies1(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies2(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies3(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies4(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies5(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies6(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies7(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies8(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies9(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies10(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies11(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies12(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies13(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies14(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies15(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies16(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies17(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies18(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies19(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies20(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies21(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies22(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies23(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies24(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies25(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies26(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies27(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies28(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies29(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies30(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies31(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies32(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies33(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies34(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies35(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies36(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies37(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies38(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies39(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies40(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies41(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies42(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies43(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies44(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies45(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies46(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies47(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies48(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies49(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies50(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies51(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies52(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies53(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies54(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies55(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies56(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies57(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies58(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies59(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies60(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies61(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies62(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies63(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies64(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies65(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies66(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies67(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies68(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies69(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies70(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies71(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies72(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies73(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies74(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies75(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies76(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies77(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies78(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies79(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies80(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies81(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies82(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies83(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies84(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies85(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies86(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies87(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies88(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies89(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies90(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies91(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies92(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies93(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies94(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies95(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies96(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies97(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies98(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies99(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies100(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies101(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies102(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies103(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies104(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies105(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies106(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies107(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies108(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies109(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies110(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies111(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies112(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies113(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies114(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies115(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies116(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies117(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies118(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies119(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies120(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies121(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies122(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies123(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies124(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies125(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies126(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies127(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies128(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies129(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies130(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies131(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies132(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies133(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies134(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies135(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies136(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies137(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies138(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies139(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies140(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies141(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies142(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies143(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies144(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies145(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies146(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies147(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies148(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies149(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies150(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies151(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies152(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies153(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies154(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies155(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies156(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies157(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies158(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies159(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies160(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies161(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies162(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies163(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies164(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies165(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies166(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies167(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies168(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies169(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies170(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies171(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies172(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies173(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies174(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies175(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies176(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies177(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies178(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies179(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies180(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies181(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies182(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies183(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies184(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies185(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies186(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies187(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies188(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies189(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies190(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies191(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies192(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies193(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies194(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies195(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies196(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies197(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies198(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies199(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies200(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies201(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies202(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies203(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies204(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies205(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies206(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies207(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies208(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies209(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies210(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies211(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies212(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies213(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies214(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies215(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies216(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies217(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies218(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies219(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies220(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies221(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies222(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies223(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies224(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies225(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies226(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies227(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies228(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies229(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies230(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies231(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies232(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies233(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies234(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies235(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies236(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies237(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies238(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies239(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies240(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies241(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies242(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies243(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies244(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies245(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies246(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies247(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies248(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies249(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies250(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies251(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies252(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies253(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies254(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies255(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies256(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies257(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies258(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies259(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies260(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies261(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies262(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies263(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies264(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies265(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies266(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies267(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies268(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies269(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies270(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies271(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies272(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies273(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies274(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies275(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies276(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies277(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies278(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies279(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies280(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies281(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies282(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies283(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies284(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies285(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies286(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies287(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies288(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies289(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies290(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies291(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies292(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies293(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies294(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies295(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies296(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies297(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies298(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies299(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies300(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies301(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies302(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies303(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies304(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies305(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies306(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies307(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies308(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies309(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies310(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies311(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies312(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies313(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies314(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies315(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies316(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies317(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies318(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies319(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies320(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies321(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies322(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies323(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies324(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies325(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies326(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies327(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies328(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies329(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies330(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies331(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies332(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies333(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies334(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies335(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies336(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies337(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies338(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies339(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies340(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies341(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies342(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies343(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies344(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies345(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies346(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies347(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies348(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies349(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies350(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies351(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies352(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies353(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies354(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies355(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies356(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies357(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies358(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies359(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies360(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies361(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies362(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies363(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies364(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies365(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies366(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies367(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies368(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies369(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies370(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies371(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies372(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies373(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies374(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies375(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies376(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies377(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies378(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies379(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies380(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies381(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies382(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies383(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies384(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies385(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies386(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies387(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies388(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies389(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies390(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies391(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies392(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies393(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies394(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies395(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies396(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies397(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies398(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies399(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies400(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies401(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies402(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies403(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies404(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies405(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies406(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies407(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies408(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies409(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies410(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies411(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies412(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies413(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies414(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies415(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies416(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies417(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies418(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies419(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies420(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies421(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies422(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies423(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies424(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies425(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies426(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies427(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies428(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies429(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies430(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies431(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies432(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies433(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies434(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies435(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies436(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies437(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies438(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies439(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies440(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies441(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies442(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies443(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies444(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies445(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies446(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies447(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies448(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies449(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies450(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies451(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies452(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies453(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies454(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies455(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies456(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies457(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies458(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies459(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies460(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies461(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies462(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies463(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies464(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies465(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies466(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies467(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies468(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies469(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies470(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies471(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies472(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies473(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies474(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies475(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies476(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies477(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies478(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies479(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies480(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies481(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies482(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies483(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies484(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies485(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies486(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies487(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies488(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies489(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies490(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies491(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies492(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies493(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies494(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies495(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies496(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies497(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies498(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies499(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies500(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies501(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies502(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies503(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies504(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies505(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies506(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies507(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies508(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies509(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies510(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies511(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies512(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies513(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies514(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies515(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies516(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies517(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies518(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies519(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies520(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies521(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies522(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies523(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies524(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies525(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies526(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies527(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies528(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies529(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies530(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies531(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies532(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies533(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies534(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies535(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies536(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies537(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies538(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies539(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies540(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies541(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies542(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies543(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies544(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies545(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies546(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies547(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies548(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies549(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies550(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies551(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies552(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies553(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies554(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies555(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies556(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies557(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies558(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies559(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies560(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies561(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies562(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies563(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies564(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies565(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies566(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies567(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies568(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies569(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies570(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies571(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies572(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies573(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies574(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies575(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies576(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies577(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies578(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies579(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies580(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies581(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies582(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies583(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies584(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies585(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies586(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies587(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies588(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies589(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies590(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies591(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies592(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies593(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies594(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies595(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies596(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies597(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies598(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies599(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies600(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies601(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies602(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies603(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies604(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies605(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies606(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies607(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies608(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies609(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies610(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies611(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies612(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies613(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies614(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies615(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies616(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies617(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies618(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies619(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies620(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies621(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies622(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies623(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies624(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies625(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies626(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies627(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies628(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies629(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies630(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies631(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies632(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies633(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies634(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies635(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies636(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies637(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies638(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies639(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies640(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies641(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies642(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies643(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies644(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies645(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies646(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies647(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies648(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies649(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies650(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies651(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies652(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies653(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies654(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies655(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies656(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies657(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies658(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies659(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies660(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies661(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies662(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies663(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies664(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies665(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies666(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies667(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies668(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies669(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies670(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies671(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies672(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies673(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies674(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies675(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies676(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies677(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies678(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies679(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies680(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies681(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies682(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies683(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies684(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies685(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies686(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies687(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies688(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies689(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies690(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies691(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies692(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies693(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies694(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies695(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies696(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies697(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies698(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies699(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies700(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies701(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies702(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies703(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies704(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies705(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies706(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies707(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies708(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies709(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies710(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies711(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies712(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies713(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies714(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies715(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies716(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies717(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies718(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies719(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies720(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies721(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies722(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies723(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies724(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies725(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies726(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies727(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies728(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies729(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies730(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies731(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies732(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies733(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies734(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies735(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies736(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies737(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies738(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies739(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies740(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies741(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies742(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies743(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies744(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies745(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies746(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies747(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies748(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies749(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies750(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies751(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies752(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies753(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies754(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies755(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies756(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies757(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies758(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies759(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies760(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies761(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies762(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies763(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies764(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies765(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies766(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies767(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies768(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies769(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies770(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies771(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies772(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies773(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies774(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies775(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies776(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies777(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies778(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies779(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies780(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies781(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies782(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies783(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies784(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies785(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies786(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies787(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies788(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies789(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies790(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies791(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies792(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies793(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies794(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies795(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies796(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies797(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies798(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies799(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies800(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies801(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies802(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies803(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies804(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies805(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies806(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies807(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies808(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies809(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies810(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies811(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies812(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies813(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies814(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies815(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies816(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies817(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies818(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies819(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies820(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies821(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies822(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies823(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies824(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies825(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies826(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies827(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies828(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies829(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies830(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies831(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies832(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies833(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies834(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies835(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies836(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies837(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies838(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies839(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies840(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies841(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies842(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies843(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies844(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies845(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies846(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies847(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies848(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies849(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies850(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies851(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies852(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies853(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies854(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies855(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies856(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies857(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies858(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies859(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies860(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies861(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies862(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies863(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies864(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies865(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies866(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies867(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies868(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies869(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies870(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies871(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies872(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies873(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies874(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies875(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies876(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies877(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies878(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies879(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies880(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies881(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies882(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies883(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies884(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies885(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies886(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies887(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies888(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies889(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies890(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies891(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies892(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies893(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies894(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies895(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies896(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies897(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies898(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies899(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies900(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies901(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies902(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies903(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies904(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies905(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies906(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies907(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies908(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies909(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies910(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies911(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies912(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies913(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies914(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies915(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies916(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies917(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies918(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies919(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies920(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies921(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies922(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies923(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies924(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies925(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies926(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies927(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies928(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies929(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies930(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies931(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies932(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies933(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies934(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies935(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies936(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies937(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies938(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies939(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies940(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies941(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies942(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies943(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies944(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies945(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies946(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies947(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies948(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies949(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies950(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies951(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies952(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies953(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies954(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies955(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies956(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies957(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies958(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies959(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies960(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies961(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies962(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies963(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies964(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies965(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies966(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies967(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies968(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies969(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies970(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies971(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies972(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies973(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies974(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies975(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies976(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies977(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies978(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies979(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies980(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies981(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies982(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies983(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies984(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies985(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies986(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies987(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies988(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies989(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies990(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies991(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies992(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies993(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies994(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies995(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies996(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies997(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies998(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}

func Test_movies999(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}

}
