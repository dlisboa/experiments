// Generated with tools/gentests. DO NOT EDIT.

//go:build integration

package users

import (
	"testing"

	"demo/internal/testdb"
)

func Test_users0(t *testing.T) {
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

func Test_users1(t *testing.T) {
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

func Test_users2(t *testing.T) {
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

func Test_users3(t *testing.T) {
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

func Test_users4(t *testing.T) {
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

func Test_users5(t *testing.T) {
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

func Test_users6(t *testing.T) {
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

func Test_users7(t *testing.T) {
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

func Test_users8(t *testing.T) {
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

func Test_users9(t *testing.T) {
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

func Test_users10(t *testing.T) {
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

func Test_users11(t *testing.T) {
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

func Test_users12(t *testing.T) {
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

func Test_users13(t *testing.T) {
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

func Test_users14(t *testing.T) {
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

func Test_users15(t *testing.T) {
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

func Test_users16(t *testing.T) {
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

func Test_users17(t *testing.T) {
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

func Test_users18(t *testing.T) {
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

func Test_users19(t *testing.T) {
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

func Test_users20(t *testing.T) {
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

func Test_users21(t *testing.T) {
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

func Test_users22(t *testing.T) {
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

func Test_users23(t *testing.T) {
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

func Test_users24(t *testing.T) {
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

func Test_users25(t *testing.T) {
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

func Test_users26(t *testing.T) {
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

func Test_users27(t *testing.T) {
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

func Test_users28(t *testing.T) {
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

func Test_users29(t *testing.T) {
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

func Test_users30(t *testing.T) {
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

func Test_users31(t *testing.T) {
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

func Test_users32(t *testing.T) {
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

func Test_users33(t *testing.T) {
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

func Test_users34(t *testing.T) {
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

func Test_users35(t *testing.T) {
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

func Test_users36(t *testing.T) {
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

func Test_users37(t *testing.T) {
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

func Test_users38(t *testing.T) {
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

func Test_users39(t *testing.T) {
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

func Test_users40(t *testing.T) {
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

func Test_users41(t *testing.T) {
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

func Test_users42(t *testing.T) {
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

func Test_users43(t *testing.T) {
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

func Test_users44(t *testing.T) {
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

func Test_users45(t *testing.T) {
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

func Test_users46(t *testing.T) {
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

func Test_users47(t *testing.T) {
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

func Test_users48(t *testing.T) {
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

func Test_users49(t *testing.T) {
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

func Test_users50(t *testing.T) {
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

func Test_users51(t *testing.T) {
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

func Test_users52(t *testing.T) {
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

func Test_users53(t *testing.T) {
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

func Test_users54(t *testing.T) {
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

func Test_users55(t *testing.T) {
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

func Test_users56(t *testing.T) {
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

func Test_users57(t *testing.T) {
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

func Test_users58(t *testing.T) {
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

func Test_users59(t *testing.T) {
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

func Test_users60(t *testing.T) {
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

func Test_users61(t *testing.T) {
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

func Test_users62(t *testing.T) {
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

func Test_users63(t *testing.T) {
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

func Test_users64(t *testing.T) {
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

func Test_users65(t *testing.T) {
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

func Test_users66(t *testing.T) {
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

func Test_users67(t *testing.T) {
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

func Test_users68(t *testing.T) {
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

func Test_users69(t *testing.T) {
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

func Test_users70(t *testing.T) {
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

func Test_users71(t *testing.T) {
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

func Test_users72(t *testing.T) {
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

func Test_users73(t *testing.T) {
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

func Test_users74(t *testing.T) {
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

func Test_users75(t *testing.T) {
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

func Test_users76(t *testing.T) {
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

func Test_users77(t *testing.T) {
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

func Test_users78(t *testing.T) {
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

func Test_users79(t *testing.T) {
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

func Test_users80(t *testing.T) {
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

func Test_users81(t *testing.T) {
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

func Test_users82(t *testing.T) {
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

func Test_users83(t *testing.T) {
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

func Test_users84(t *testing.T) {
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

func Test_users85(t *testing.T) {
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

func Test_users86(t *testing.T) {
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

func Test_users87(t *testing.T) {
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

func Test_users88(t *testing.T) {
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

func Test_users89(t *testing.T) {
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

func Test_users90(t *testing.T) {
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

func Test_users91(t *testing.T) {
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

func Test_users92(t *testing.T) {
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

func Test_users93(t *testing.T) {
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

func Test_users94(t *testing.T) {
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

func Test_users95(t *testing.T) {
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

func Test_users96(t *testing.T) {
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

func Test_users97(t *testing.T) {
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

func Test_users98(t *testing.T) {
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

func Test_users99(t *testing.T) {
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

func Test_users100(t *testing.T) {
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

func Test_users101(t *testing.T) {
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

func Test_users102(t *testing.T) {
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

func Test_users103(t *testing.T) {
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

func Test_users104(t *testing.T) {
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

func Test_users105(t *testing.T) {
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

func Test_users106(t *testing.T) {
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

func Test_users107(t *testing.T) {
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

func Test_users108(t *testing.T) {
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

func Test_users109(t *testing.T) {
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

func Test_users110(t *testing.T) {
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

func Test_users111(t *testing.T) {
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

func Test_users112(t *testing.T) {
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

func Test_users113(t *testing.T) {
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

func Test_users114(t *testing.T) {
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

func Test_users115(t *testing.T) {
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

func Test_users116(t *testing.T) {
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

func Test_users117(t *testing.T) {
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

func Test_users118(t *testing.T) {
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

func Test_users119(t *testing.T) {
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

func Test_users120(t *testing.T) {
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

func Test_users121(t *testing.T) {
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

func Test_users122(t *testing.T) {
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

func Test_users123(t *testing.T) {
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

func Test_users124(t *testing.T) {
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

func Test_users125(t *testing.T) {
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

func Test_users126(t *testing.T) {
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

func Test_users127(t *testing.T) {
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

func Test_users128(t *testing.T) {
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

func Test_users129(t *testing.T) {
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

func Test_users130(t *testing.T) {
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

func Test_users131(t *testing.T) {
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

func Test_users132(t *testing.T) {
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

func Test_users133(t *testing.T) {
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

func Test_users134(t *testing.T) {
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

func Test_users135(t *testing.T) {
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

func Test_users136(t *testing.T) {
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

func Test_users137(t *testing.T) {
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

func Test_users138(t *testing.T) {
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

func Test_users139(t *testing.T) {
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

func Test_users140(t *testing.T) {
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

func Test_users141(t *testing.T) {
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

func Test_users142(t *testing.T) {
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

func Test_users143(t *testing.T) {
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

func Test_users144(t *testing.T) {
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

func Test_users145(t *testing.T) {
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

func Test_users146(t *testing.T) {
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

func Test_users147(t *testing.T) {
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

func Test_users148(t *testing.T) {
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

func Test_users149(t *testing.T) {
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

func Test_users150(t *testing.T) {
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

func Test_users151(t *testing.T) {
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

func Test_users152(t *testing.T) {
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

func Test_users153(t *testing.T) {
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

func Test_users154(t *testing.T) {
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

func Test_users155(t *testing.T) {
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

func Test_users156(t *testing.T) {
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

func Test_users157(t *testing.T) {
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

func Test_users158(t *testing.T) {
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

func Test_users159(t *testing.T) {
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

func Test_users160(t *testing.T) {
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

func Test_users161(t *testing.T) {
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

func Test_users162(t *testing.T) {
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

func Test_users163(t *testing.T) {
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

func Test_users164(t *testing.T) {
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

func Test_users165(t *testing.T) {
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

func Test_users166(t *testing.T) {
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

func Test_users167(t *testing.T) {
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

func Test_users168(t *testing.T) {
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

func Test_users169(t *testing.T) {
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

func Test_users170(t *testing.T) {
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

func Test_users171(t *testing.T) {
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

func Test_users172(t *testing.T) {
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

func Test_users173(t *testing.T) {
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

func Test_users174(t *testing.T) {
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

func Test_users175(t *testing.T) {
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

func Test_users176(t *testing.T) {
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

func Test_users177(t *testing.T) {
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

func Test_users178(t *testing.T) {
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

func Test_users179(t *testing.T) {
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

func Test_users180(t *testing.T) {
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

func Test_users181(t *testing.T) {
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

func Test_users182(t *testing.T) {
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

func Test_users183(t *testing.T) {
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

func Test_users184(t *testing.T) {
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

func Test_users185(t *testing.T) {
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

func Test_users186(t *testing.T) {
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

func Test_users187(t *testing.T) {
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

func Test_users188(t *testing.T) {
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

func Test_users189(t *testing.T) {
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

func Test_users190(t *testing.T) {
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

func Test_users191(t *testing.T) {
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

func Test_users192(t *testing.T) {
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

func Test_users193(t *testing.T) {
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

func Test_users194(t *testing.T) {
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

func Test_users195(t *testing.T) {
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

func Test_users196(t *testing.T) {
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

func Test_users197(t *testing.T) {
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

func Test_users198(t *testing.T) {
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

func Test_users199(t *testing.T) {
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

func Test_users200(t *testing.T) {
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

func Test_users201(t *testing.T) {
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

func Test_users202(t *testing.T) {
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

func Test_users203(t *testing.T) {
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

func Test_users204(t *testing.T) {
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

func Test_users205(t *testing.T) {
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

func Test_users206(t *testing.T) {
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

func Test_users207(t *testing.T) {
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

func Test_users208(t *testing.T) {
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

func Test_users209(t *testing.T) {
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

func Test_users210(t *testing.T) {
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

func Test_users211(t *testing.T) {
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

func Test_users212(t *testing.T) {
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

func Test_users213(t *testing.T) {
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

func Test_users214(t *testing.T) {
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

func Test_users215(t *testing.T) {
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

func Test_users216(t *testing.T) {
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

func Test_users217(t *testing.T) {
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

func Test_users218(t *testing.T) {
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

func Test_users219(t *testing.T) {
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

func Test_users220(t *testing.T) {
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

func Test_users221(t *testing.T) {
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

func Test_users222(t *testing.T) {
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

func Test_users223(t *testing.T) {
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

func Test_users224(t *testing.T) {
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

func Test_users225(t *testing.T) {
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

func Test_users226(t *testing.T) {
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

func Test_users227(t *testing.T) {
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

func Test_users228(t *testing.T) {
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

func Test_users229(t *testing.T) {
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

func Test_users230(t *testing.T) {
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

func Test_users231(t *testing.T) {
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

func Test_users232(t *testing.T) {
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

func Test_users233(t *testing.T) {
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

func Test_users234(t *testing.T) {
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

func Test_users235(t *testing.T) {
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

func Test_users236(t *testing.T) {
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

func Test_users237(t *testing.T) {
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

func Test_users238(t *testing.T) {
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

func Test_users239(t *testing.T) {
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

func Test_users240(t *testing.T) {
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

func Test_users241(t *testing.T) {
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

func Test_users242(t *testing.T) {
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

func Test_users243(t *testing.T) {
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

func Test_users244(t *testing.T) {
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

func Test_users245(t *testing.T) {
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

func Test_users246(t *testing.T) {
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

func Test_users247(t *testing.T) {
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

func Test_users248(t *testing.T) {
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

func Test_users249(t *testing.T) {
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

func Test_users250(t *testing.T) {
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

func Test_users251(t *testing.T) {
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

func Test_users252(t *testing.T) {
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

func Test_users253(t *testing.T) {
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

func Test_users254(t *testing.T) {
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

func Test_users255(t *testing.T) {
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

func Test_users256(t *testing.T) {
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

func Test_users257(t *testing.T) {
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

func Test_users258(t *testing.T) {
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

func Test_users259(t *testing.T) {
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

func Test_users260(t *testing.T) {
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

func Test_users261(t *testing.T) {
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

func Test_users262(t *testing.T) {
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

func Test_users263(t *testing.T) {
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

func Test_users264(t *testing.T) {
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

func Test_users265(t *testing.T) {
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

func Test_users266(t *testing.T) {
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

func Test_users267(t *testing.T) {
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

func Test_users268(t *testing.T) {
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

func Test_users269(t *testing.T) {
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

func Test_users270(t *testing.T) {
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

func Test_users271(t *testing.T) {
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

func Test_users272(t *testing.T) {
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

func Test_users273(t *testing.T) {
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

func Test_users274(t *testing.T) {
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

func Test_users275(t *testing.T) {
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

func Test_users276(t *testing.T) {
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

func Test_users277(t *testing.T) {
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

func Test_users278(t *testing.T) {
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

func Test_users279(t *testing.T) {
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

func Test_users280(t *testing.T) {
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

func Test_users281(t *testing.T) {
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

func Test_users282(t *testing.T) {
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

func Test_users283(t *testing.T) {
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

func Test_users284(t *testing.T) {
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

func Test_users285(t *testing.T) {
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

func Test_users286(t *testing.T) {
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

func Test_users287(t *testing.T) {
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

func Test_users288(t *testing.T) {
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

func Test_users289(t *testing.T) {
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

func Test_users290(t *testing.T) {
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

func Test_users291(t *testing.T) {
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

func Test_users292(t *testing.T) {
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

func Test_users293(t *testing.T) {
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

func Test_users294(t *testing.T) {
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

func Test_users295(t *testing.T) {
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

func Test_users296(t *testing.T) {
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

func Test_users297(t *testing.T) {
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

func Test_users298(t *testing.T) {
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

func Test_users299(t *testing.T) {
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

func Test_users300(t *testing.T) {
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

func Test_users301(t *testing.T) {
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

func Test_users302(t *testing.T) {
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

func Test_users303(t *testing.T) {
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

func Test_users304(t *testing.T) {
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

func Test_users305(t *testing.T) {
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

func Test_users306(t *testing.T) {
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

func Test_users307(t *testing.T) {
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

func Test_users308(t *testing.T) {
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

func Test_users309(t *testing.T) {
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

func Test_users310(t *testing.T) {
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

func Test_users311(t *testing.T) {
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

func Test_users312(t *testing.T) {
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

func Test_users313(t *testing.T) {
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

func Test_users314(t *testing.T) {
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

func Test_users315(t *testing.T) {
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

func Test_users316(t *testing.T) {
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

func Test_users317(t *testing.T) {
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

func Test_users318(t *testing.T) {
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

func Test_users319(t *testing.T) {
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

func Test_users320(t *testing.T) {
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

func Test_users321(t *testing.T) {
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

func Test_users322(t *testing.T) {
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

func Test_users323(t *testing.T) {
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

func Test_users324(t *testing.T) {
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

func Test_users325(t *testing.T) {
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

func Test_users326(t *testing.T) {
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

func Test_users327(t *testing.T) {
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

func Test_users328(t *testing.T) {
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

func Test_users329(t *testing.T) {
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

func Test_users330(t *testing.T) {
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

func Test_users331(t *testing.T) {
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

func Test_users332(t *testing.T) {
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

func Test_users333(t *testing.T) {
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

func Test_users334(t *testing.T) {
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

func Test_users335(t *testing.T) {
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

func Test_users336(t *testing.T) {
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

func Test_users337(t *testing.T) {
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

func Test_users338(t *testing.T) {
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

func Test_users339(t *testing.T) {
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

func Test_users340(t *testing.T) {
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

func Test_users341(t *testing.T) {
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

func Test_users342(t *testing.T) {
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

func Test_users343(t *testing.T) {
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

func Test_users344(t *testing.T) {
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

func Test_users345(t *testing.T) {
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

func Test_users346(t *testing.T) {
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

func Test_users347(t *testing.T) {
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

func Test_users348(t *testing.T) {
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

func Test_users349(t *testing.T) {
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

func Test_users350(t *testing.T) {
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

func Test_users351(t *testing.T) {
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

func Test_users352(t *testing.T) {
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

func Test_users353(t *testing.T) {
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

func Test_users354(t *testing.T) {
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

func Test_users355(t *testing.T) {
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

func Test_users356(t *testing.T) {
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

func Test_users357(t *testing.T) {
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

func Test_users358(t *testing.T) {
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

func Test_users359(t *testing.T) {
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

func Test_users360(t *testing.T) {
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

func Test_users361(t *testing.T) {
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

func Test_users362(t *testing.T) {
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

func Test_users363(t *testing.T) {
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

func Test_users364(t *testing.T) {
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

func Test_users365(t *testing.T) {
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

func Test_users366(t *testing.T) {
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

func Test_users367(t *testing.T) {
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

func Test_users368(t *testing.T) {
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

func Test_users369(t *testing.T) {
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

func Test_users370(t *testing.T) {
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

func Test_users371(t *testing.T) {
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

func Test_users372(t *testing.T) {
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

func Test_users373(t *testing.T) {
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

func Test_users374(t *testing.T) {
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

func Test_users375(t *testing.T) {
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

func Test_users376(t *testing.T) {
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

func Test_users377(t *testing.T) {
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

func Test_users378(t *testing.T) {
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

func Test_users379(t *testing.T) {
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

func Test_users380(t *testing.T) {
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

func Test_users381(t *testing.T) {
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

func Test_users382(t *testing.T) {
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

func Test_users383(t *testing.T) {
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

func Test_users384(t *testing.T) {
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

func Test_users385(t *testing.T) {
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

func Test_users386(t *testing.T) {
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

func Test_users387(t *testing.T) {
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

func Test_users388(t *testing.T) {
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

func Test_users389(t *testing.T) {
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

func Test_users390(t *testing.T) {
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

func Test_users391(t *testing.T) {
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

func Test_users392(t *testing.T) {
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

func Test_users393(t *testing.T) {
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

func Test_users394(t *testing.T) {
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

func Test_users395(t *testing.T) {
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

func Test_users396(t *testing.T) {
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

func Test_users397(t *testing.T) {
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

func Test_users398(t *testing.T) {
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

func Test_users399(t *testing.T) {
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

func Test_users400(t *testing.T) {
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

func Test_users401(t *testing.T) {
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

func Test_users402(t *testing.T) {
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

func Test_users403(t *testing.T) {
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

func Test_users404(t *testing.T) {
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

func Test_users405(t *testing.T) {
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

func Test_users406(t *testing.T) {
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

func Test_users407(t *testing.T) {
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

func Test_users408(t *testing.T) {
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

func Test_users409(t *testing.T) {
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

func Test_users410(t *testing.T) {
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

func Test_users411(t *testing.T) {
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

func Test_users412(t *testing.T) {
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

func Test_users413(t *testing.T) {
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

func Test_users414(t *testing.T) {
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

func Test_users415(t *testing.T) {
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

func Test_users416(t *testing.T) {
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

func Test_users417(t *testing.T) {
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

func Test_users418(t *testing.T) {
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

func Test_users419(t *testing.T) {
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

func Test_users420(t *testing.T) {
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

func Test_users421(t *testing.T) {
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

func Test_users422(t *testing.T) {
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

func Test_users423(t *testing.T) {
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

func Test_users424(t *testing.T) {
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

func Test_users425(t *testing.T) {
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

func Test_users426(t *testing.T) {
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

func Test_users427(t *testing.T) {
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

func Test_users428(t *testing.T) {
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

func Test_users429(t *testing.T) {
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

func Test_users430(t *testing.T) {
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

func Test_users431(t *testing.T) {
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

func Test_users432(t *testing.T) {
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

func Test_users433(t *testing.T) {
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

func Test_users434(t *testing.T) {
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

func Test_users435(t *testing.T) {
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

func Test_users436(t *testing.T) {
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

func Test_users437(t *testing.T) {
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

func Test_users438(t *testing.T) {
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

func Test_users439(t *testing.T) {
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

func Test_users440(t *testing.T) {
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

func Test_users441(t *testing.T) {
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

func Test_users442(t *testing.T) {
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

func Test_users443(t *testing.T) {
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

func Test_users444(t *testing.T) {
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

func Test_users445(t *testing.T) {
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

func Test_users446(t *testing.T) {
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

func Test_users447(t *testing.T) {
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

func Test_users448(t *testing.T) {
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

func Test_users449(t *testing.T) {
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

func Test_users450(t *testing.T) {
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

func Test_users451(t *testing.T) {
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

func Test_users452(t *testing.T) {
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

func Test_users453(t *testing.T) {
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

func Test_users454(t *testing.T) {
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

func Test_users455(t *testing.T) {
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

func Test_users456(t *testing.T) {
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

func Test_users457(t *testing.T) {
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

func Test_users458(t *testing.T) {
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

func Test_users459(t *testing.T) {
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

func Test_users460(t *testing.T) {
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

func Test_users461(t *testing.T) {
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

func Test_users462(t *testing.T) {
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

func Test_users463(t *testing.T) {
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

func Test_users464(t *testing.T) {
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

func Test_users465(t *testing.T) {
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

func Test_users466(t *testing.T) {
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

func Test_users467(t *testing.T) {
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

func Test_users468(t *testing.T) {
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

func Test_users469(t *testing.T) {
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

func Test_users470(t *testing.T) {
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

func Test_users471(t *testing.T) {
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

func Test_users472(t *testing.T) {
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

func Test_users473(t *testing.T) {
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

func Test_users474(t *testing.T) {
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

func Test_users475(t *testing.T) {
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

func Test_users476(t *testing.T) {
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

func Test_users477(t *testing.T) {
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

func Test_users478(t *testing.T) {
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

func Test_users479(t *testing.T) {
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

func Test_users480(t *testing.T) {
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

func Test_users481(t *testing.T) {
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

func Test_users482(t *testing.T) {
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

func Test_users483(t *testing.T) {
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

func Test_users484(t *testing.T) {
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

func Test_users485(t *testing.T) {
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

func Test_users486(t *testing.T) {
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

func Test_users487(t *testing.T) {
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

func Test_users488(t *testing.T) {
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

func Test_users489(t *testing.T) {
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

func Test_users490(t *testing.T) {
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

func Test_users491(t *testing.T) {
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

func Test_users492(t *testing.T) {
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

func Test_users493(t *testing.T) {
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

func Test_users494(t *testing.T) {
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

func Test_users495(t *testing.T) {
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

func Test_users496(t *testing.T) {
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

func Test_users497(t *testing.T) {
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

func Test_users498(t *testing.T) {
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

func Test_users499(t *testing.T) {
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

func Test_users500(t *testing.T) {
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

func Test_users501(t *testing.T) {
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

func Test_users502(t *testing.T) {
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

func Test_users503(t *testing.T) {
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

func Test_users504(t *testing.T) {
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

func Test_users505(t *testing.T) {
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

func Test_users506(t *testing.T) {
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

func Test_users507(t *testing.T) {
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

func Test_users508(t *testing.T) {
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

func Test_users509(t *testing.T) {
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

func Test_users510(t *testing.T) {
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

func Test_users511(t *testing.T) {
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

func Test_users512(t *testing.T) {
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

func Test_users513(t *testing.T) {
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

func Test_users514(t *testing.T) {
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

func Test_users515(t *testing.T) {
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

func Test_users516(t *testing.T) {
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

func Test_users517(t *testing.T) {
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

func Test_users518(t *testing.T) {
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

func Test_users519(t *testing.T) {
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

func Test_users520(t *testing.T) {
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

func Test_users521(t *testing.T) {
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

func Test_users522(t *testing.T) {
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

func Test_users523(t *testing.T) {
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

func Test_users524(t *testing.T) {
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

func Test_users525(t *testing.T) {
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

func Test_users526(t *testing.T) {
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

func Test_users527(t *testing.T) {
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

func Test_users528(t *testing.T) {
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

func Test_users529(t *testing.T) {
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

func Test_users530(t *testing.T) {
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

func Test_users531(t *testing.T) {
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

func Test_users532(t *testing.T) {
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

func Test_users533(t *testing.T) {
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

func Test_users534(t *testing.T) {
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

func Test_users535(t *testing.T) {
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

func Test_users536(t *testing.T) {
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

func Test_users537(t *testing.T) {
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

func Test_users538(t *testing.T) {
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

func Test_users539(t *testing.T) {
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

func Test_users540(t *testing.T) {
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

func Test_users541(t *testing.T) {
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

func Test_users542(t *testing.T) {
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

func Test_users543(t *testing.T) {
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

func Test_users544(t *testing.T) {
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

func Test_users545(t *testing.T) {
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

func Test_users546(t *testing.T) {
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

func Test_users547(t *testing.T) {
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

func Test_users548(t *testing.T) {
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

func Test_users549(t *testing.T) {
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

func Test_users550(t *testing.T) {
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

func Test_users551(t *testing.T) {
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

func Test_users552(t *testing.T) {
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

func Test_users553(t *testing.T) {
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

func Test_users554(t *testing.T) {
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

func Test_users555(t *testing.T) {
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

func Test_users556(t *testing.T) {
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

func Test_users557(t *testing.T) {
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

func Test_users558(t *testing.T) {
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

func Test_users559(t *testing.T) {
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

func Test_users560(t *testing.T) {
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

func Test_users561(t *testing.T) {
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

func Test_users562(t *testing.T) {
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

func Test_users563(t *testing.T) {
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

func Test_users564(t *testing.T) {
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

func Test_users565(t *testing.T) {
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

func Test_users566(t *testing.T) {
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

func Test_users567(t *testing.T) {
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

func Test_users568(t *testing.T) {
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

func Test_users569(t *testing.T) {
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

func Test_users570(t *testing.T) {
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

func Test_users571(t *testing.T) {
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

func Test_users572(t *testing.T) {
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

func Test_users573(t *testing.T) {
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

func Test_users574(t *testing.T) {
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

func Test_users575(t *testing.T) {
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

func Test_users576(t *testing.T) {
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

func Test_users577(t *testing.T) {
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

func Test_users578(t *testing.T) {
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

func Test_users579(t *testing.T) {
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

func Test_users580(t *testing.T) {
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

func Test_users581(t *testing.T) {
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

func Test_users582(t *testing.T) {
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

func Test_users583(t *testing.T) {
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

func Test_users584(t *testing.T) {
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

func Test_users585(t *testing.T) {
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

func Test_users586(t *testing.T) {
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

func Test_users587(t *testing.T) {
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

func Test_users588(t *testing.T) {
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

func Test_users589(t *testing.T) {
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

func Test_users590(t *testing.T) {
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

func Test_users591(t *testing.T) {
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

func Test_users592(t *testing.T) {
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

func Test_users593(t *testing.T) {
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

func Test_users594(t *testing.T) {
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

func Test_users595(t *testing.T) {
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

func Test_users596(t *testing.T) {
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

func Test_users597(t *testing.T) {
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

func Test_users598(t *testing.T) {
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

func Test_users599(t *testing.T) {
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

func Test_users600(t *testing.T) {
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

func Test_users601(t *testing.T) {
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

func Test_users602(t *testing.T) {
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

func Test_users603(t *testing.T) {
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

func Test_users604(t *testing.T) {
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

func Test_users605(t *testing.T) {
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

func Test_users606(t *testing.T) {
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

func Test_users607(t *testing.T) {
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

func Test_users608(t *testing.T) {
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

func Test_users609(t *testing.T) {
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

func Test_users610(t *testing.T) {
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

func Test_users611(t *testing.T) {
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

func Test_users612(t *testing.T) {
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

func Test_users613(t *testing.T) {
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

func Test_users614(t *testing.T) {
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

func Test_users615(t *testing.T) {
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

func Test_users616(t *testing.T) {
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

func Test_users617(t *testing.T) {
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

func Test_users618(t *testing.T) {
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

func Test_users619(t *testing.T) {
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

func Test_users620(t *testing.T) {
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

func Test_users621(t *testing.T) {
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

func Test_users622(t *testing.T) {
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

func Test_users623(t *testing.T) {
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

func Test_users624(t *testing.T) {
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

func Test_users625(t *testing.T) {
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

func Test_users626(t *testing.T) {
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

func Test_users627(t *testing.T) {
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

func Test_users628(t *testing.T) {
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

func Test_users629(t *testing.T) {
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

func Test_users630(t *testing.T) {
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

func Test_users631(t *testing.T) {
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

func Test_users632(t *testing.T) {
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

func Test_users633(t *testing.T) {
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

func Test_users634(t *testing.T) {
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

func Test_users635(t *testing.T) {
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

func Test_users636(t *testing.T) {
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

func Test_users637(t *testing.T) {
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

func Test_users638(t *testing.T) {
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

func Test_users639(t *testing.T) {
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

func Test_users640(t *testing.T) {
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

func Test_users641(t *testing.T) {
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

func Test_users642(t *testing.T) {
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

func Test_users643(t *testing.T) {
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

func Test_users644(t *testing.T) {
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

func Test_users645(t *testing.T) {
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

func Test_users646(t *testing.T) {
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

func Test_users647(t *testing.T) {
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

func Test_users648(t *testing.T) {
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

func Test_users649(t *testing.T) {
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

func Test_users650(t *testing.T) {
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

func Test_users651(t *testing.T) {
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

func Test_users652(t *testing.T) {
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

func Test_users653(t *testing.T) {
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

func Test_users654(t *testing.T) {
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

func Test_users655(t *testing.T) {
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

func Test_users656(t *testing.T) {
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

func Test_users657(t *testing.T) {
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

func Test_users658(t *testing.T) {
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

func Test_users659(t *testing.T) {
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

func Test_users660(t *testing.T) {
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

func Test_users661(t *testing.T) {
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

func Test_users662(t *testing.T) {
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

func Test_users663(t *testing.T) {
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

func Test_users664(t *testing.T) {
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

func Test_users665(t *testing.T) {
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

func Test_users666(t *testing.T) {
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

func Test_users667(t *testing.T) {
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

func Test_users668(t *testing.T) {
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

func Test_users669(t *testing.T) {
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

func Test_users670(t *testing.T) {
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

func Test_users671(t *testing.T) {
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

func Test_users672(t *testing.T) {
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

func Test_users673(t *testing.T) {
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

func Test_users674(t *testing.T) {
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

func Test_users675(t *testing.T) {
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

func Test_users676(t *testing.T) {
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

func Test_users677(t *testing.T) {
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

func Test_users678(t *testing.T) {
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

func Test_users679(t *testing.T) {
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

func Test_users680(t *testing.T) {
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

func Test_users681(t *testing.T) {
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

func Test_users682(t *testing.T) {
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

func Test_users683(t *testing.T) {
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

func Test_users684(t *testing.T) {
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

func Test_users685(t *testing.T) {
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

func Test_users686(t *testing.T) {
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

func Test_users687(t *testing.T) {
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

func Test_users688(t *testing.T) {
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

func Test_users689(t *testing.T) {
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

func Test_users690(t *testing.T) {
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

func Test_users691(t *testing.T) {
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

func Test_users692(t *testing.T) {
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

func Test_users693(t *testing.T) {
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

func Test_users694(t *testing.T) {
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

func Test_users695(t *testing.T) {
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

func Test_users696(t *testing.T) {
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

func Test_users697(t *testing.T) {
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

func Test_users698(t *testing.T) {
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

func Test_users699(t *testing.T) {
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

func Test_users700(t *testing.T) {
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

func Test_users701(t *testing.T) {
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

func Test_users702(t *testing.T) {
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

func Test_users703(t *testing.T) {
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

func Test_users704(t *testing.T) {
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

func Test_users705(t *testing.T) {
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

func Test_users706(t *testing.T) {
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

func Test_users707(t *testing.T) {
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

func Test_users708(t *testing.T) {
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

func Test_users709(t *testing.T) {
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

func Test_users710(t *testing.T) {
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

func Test_users711(t *testing.T) {
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

func Test_users712(t *testing.T) {
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

func Test_users713(t *testing.T) {
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

func Test_users714(t *testing.T) {
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

func Test_users715(t *testing.T) {
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

func Test_users716(t *testing.T) {
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

func Test_users717(t *testing.T) {
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

func Test_users718(t *testing.T) {
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

func Test_users719(t *testing.T) {
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

func Test_users720(t *testing.T) {
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

func Test_users721(t *testing.T) {
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

func Test_users722(t *testing.T) {
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

func Test_users723(t *testing.T) {
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

func Test_users724(t *testing.T) {
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

func Test_users725(t *testing.T) {
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

func Test_users726(t *testing.T) {
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

func Test_users727(t *testing.T) {
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

func Test_users728(t *testing.T) {
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

func Test_users729(t *testing.T) {
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

func Test_users730(t *testing.T) {
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

func Test_users731(t *testing.T) {
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

func Test_users732(t *testing.T) {
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

func Test_users733(t *testing.T) {
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

func Test_users734(t *testing.T) {
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

func Test_users735(t *testing.T) {
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

func Test_users736(t *testing.T) {
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

func Test_users737(t *testing.T) {
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

func Test_users738(t *testing.T) {
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

func Test_users739(t *testing.T) {
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

func Test_users740(t *testing.T) {
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

func Test_users741(t *testing.T) {
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

func Test_users742(t *testing.T) {
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

func Test_users743(t *testing.T) {
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

func Test_users744(t *testing.T) {
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

func Test_users745(t *testing.T) {
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

func Test_users746(t *testing.T) {
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

func Test_users747(t *testing.T) {
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

func Test_users748(t *testing.T) {
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

func Test_users749(t *testing.T) {
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

func Test_users750(t *testing.T) {
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

func Test_users751(t *testing.T) {
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

func Test_users752(t *testing.T) {
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

func Test_users753(t *testing.T) {
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

func Test_users754(t *testing.T) {
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

func Test_users755(t *testing.T) {
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

func Test_users756(t *testing.T) {
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

func Test_users757(t *testing.T) {
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

func Test_users758(t *testing.T) {
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

func Test_users759(t *testing.T) {
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

func Test_users760(t *testing.T) {
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

func Test_users761(t *testing.T) {
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

func Test_users762(t *testing.T) {
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

func Test_users763(t *testing.T) {
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

func Test_users764(t *testing.T) {
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

func Test_users765(t *testing.T) {
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

func Test_users766(t *testing.T) {
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

func Test_users767(t *testing.T) {
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

func Test_users768(t *testing.T) {
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

func Test_users769(t *testing.T) {
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

func Test_users770(t *testing.T) {
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

func Test_users771(t *testing.T) {
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

func Test_users772(t *testing.T) {
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

func Test_users773(t *testing.T) {
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

func Test_users774(t *testing.T) {
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

func Test_users775(t *testing.T) {
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

func Test_users776(t *testing.T) {
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

func Test_users777(t *testing.T) {
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

func Test_users778(t *testing.T) {
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

func Test_users779(t *testing.T) {
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

func Test_users780(t *testing.T) {
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

func Test_users781(t *testing.T) {
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

func Test_users782(t *testing.T) {
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

func Test_users783(t *testing.T) {
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

func Test_users784(t *testing.T) {
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

func Test_users785(t *testing.T) {
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

func Test_users786(t *testing.T) {
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

func Test_users787(t *testing.T) {
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

func Test_users788(t *testing.T) {
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

func Test_users789(t *testing.T) {
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

func Test_users790(t *testing.T) {
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

func Test_users791(t *testing.T) {
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

func Test_users792(t *testing.T) {
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

func Test_users793(t *testing.T) {
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

func Test_users794(t *testing.T) {
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

func Test_users795(t *testing.T) {
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

func Test_users796(t *testing.T) {
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

func Test_users797(t *testing.T) {
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

func Test_users798(t *testing.T) {
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

func Test_users799(t *testing.T) {
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

func Test_users800(t *testing.T) {
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

func Test_users801(t *testing.T) {
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

func Test_users802(t *testing.T) {
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

func Test_users803(t *testing.T) {
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

func Test_users804(t *testing.T) {
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

func Test_users805(t *testing.T) {
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

func Test_users806(t *testing.T) {
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

func Test_users807(t *testing.T) {
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

func Test_users808(t *testing.T) {
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

func Test_users809(t *testing.T) {
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

func Test_users810(t *testing.T) {
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

func Test_users811(t *testing.T) {
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

func Test_users812(t *testing.T) {
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

func Test_users813(t *testing.T) {
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

func Test_users814(t *testing.T) {
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

func Test_users815(t *testing.T) {
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

func Test_users816(t *testing.T) {
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

func Test_users817(t *testing.T) {
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

func Test_users818(t *testing.T) {
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

func Test_users819(t *testing.T) {
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

func Test_users820(t *testing.T) {
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

func Test_users821(t *testing.T) {
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

func Test_users822(t *testing.T) {
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

func Test_users823(t *testing.T) {
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

func Test_users824(t *testing.T) {
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

func Test_users825(t *testing.T) {
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

func Test_users826(t *testing.T) {
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

func Test_users827(t *testing.T) {
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

func Test_users828(t *testing.T) {
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

func Test_users829(t *testing.T) {
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

func Test_users830(t *testing.T) {
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

func Test_users831(t *testing.T) {
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

func Test_users832(t *testing.T) {
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

func Test_users833(t *testing.T) {
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

func Test_users834(t *testing.T) {
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

func Test_users835(t *testing.T) {
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

func Test_users836(t *testing.T) {
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

func Test_users837(t *testing.T) {
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

func Test_users838(t *testing.T) {
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

func Test_users839(t *testing.T) {
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

func Test_users840(t *testing.T) {
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

func Test_users841(t *testing.T) {
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

func Test_users842(t *testing.T) {
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

func Test_users843(t *testing.T) {
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

func Test_users844(t *testing.T) {
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

func Test_users845(t *testing.T) {
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

func Test_users846(t *testing.T) {
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

func Test_users847(t *testing.T) {
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

func Test_users848(t *testing.T) {
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

func Test_users849(t *testing.T) {
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

func Test_users850(t *testing.T) {
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

func Test_users851(t *testing.T) {
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

func Test_users852(t *testing.T) {
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

func Test_users853(t *testing.T) {
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

func Test_users854(t *testing.T) {
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

func Test_users855(t *testing.T) {
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

func Test_users856(t *testing.T) {
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

func Test_users857(t *testing.T) {
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

func Test_users858(t *testing.T) {
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

func Test_users859(t *testing.T) {
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

func Test_users860(t *testing.T) {
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

func Test_users861(t *testing.T) {
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

func Test_users862(t *testing.T) {
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

func Test_users863(t *testing.T) {
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

func Test_users864(t *testing.T) {
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

func Test_users865(t *testing.T) {
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

func Test_users866(t *testing.T) {
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

func Test_users867(t *testing.T) {
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

func Test_users868(t *testing.T) {
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

func Test_users869(t *testing.T) {
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

func Test_users870(t *testing.T) {
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

func Test_users871(t *testing.T) {
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

func Test_users872(t *testing.T) {
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

func Test_users873(t *testing.T) {
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

func Test_users874(t *testing.T) {
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

func Test_users875(t *testing.T) {
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

func Test_users876(t *testing.T) {
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

func Test_users877(t *testing.T) {
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

func Test_users878(t *testing.T) {
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

func Test_users879(t *testing.T) {
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

func Test_users880(t *testing.T) {
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

func Test_users881(t *testing.T) {
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

func Test_users882(t *testing.T) {
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

func Test_users883(t *testing.T) {
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

func Test_users884(t *testing.T) {
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

func Test_users885(t *testing.T) {
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

func Test_users886(t *testing.T) {
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

func Test_users887(t *testing.T) {
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

func Test_users888(t *testing.T) {
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

func Test_users889(t *testing.T) {
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

func Test_users890(t *testing.T) {
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

func Test_users891(t *testing.T) {
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

func Test_users892(t *testing.T) {
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

func Test_users893(t *testing.T) {
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

func Test_users894(t *testing.T) {
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

func Test_users895(t *testing.T) {
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

func Test_users896(t *testing.T) {
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

func Test_users897(t *testing.T) {
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

func Test_users898(t *testing.T) {
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

func Test_users899(t *testing.T) {
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

func Test_users900(t *testing.T) {
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

func Test_users901(t *testing.T) {
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

func Test_users902(t *testing.T) {
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

func Test_users903(t *testing.T) {
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

func Test_users904(t *testing.T) {
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

func Test_users905(t *testing.T) {
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

func Test_users906(t *testing.T) {
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

func Test_users907(t *testing.T) {
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

func Test_users908(t *testing.T) {
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

func Test_users909(t *testing.T) {
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

func Test_users910(t *testing.T) {
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

func Test_users911(t *testing.T) {
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

func Test_users912(t *testing.T) {
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

func Test_users913(t *testing.T) {
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

func Test_users914(t *testing.T) {
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

func Test_users915(t *testing.T) {
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

func Test_users916(t *testing.T) {
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

func Test_users917(t *testing.T) {
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

func Test_users918(t *testing.T) {
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

func Test_users919(t *testing.T) {
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

func Test_users920(t *testing.T) {
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

func Test_users921(t *testing.T) {
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

func Test_users922(t *testing.T) {
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

func Test_users923(t *testing.T) {
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

func Test_users924(t *testing.T) {
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

func Test_users925(t *testing.T) {
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

func Test_users926(t *testing.T) {
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

func Test_users927(t *testing.T) {
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

func Test_users928(t *testing.T) {
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

func Test_users929(t *testing.T) {
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

func Test_users930(t *testing.T) {
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

func Test_users931(t *testing.T) {
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

func Test_users932(t *testing.T) {
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

func Test_users933(t *testing.T) {
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

func Test_users934(t *testing.T) {
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

func Test_users935(t *testing.T) {
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

func Test_users936(t *testing.T) {
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

func Test_users937(t *testing.T) {
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

func Test_users938(t *testing.T) {
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

func Test_users939(t *testing.T) {
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

func Test_users940(t *testing.T) {
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

func Test_users941(t *testing.T) {
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

func Test_users942(t *testing.T) {
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

func Test_users943(t *testing.T) {
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

func Test_users944(t *testing.T) {
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

func Test_users945(t *testing.T) {
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

func Test_users946(t *testing.T) {
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

func Test_users947(t *testing.T) {
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

func Test_users948(t *testing.T) {
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

func Test_users949(t *testing.T) {
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

func Test_users950(t *testing.T) {
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

func Test_users951(t *testing.T) {
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

func Test_users952(t *testing.T) {
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

func Test_users953(t *testing.T) {
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

func Test_users954(t *testing.T) {
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

func Test_users955(t *testing.T) {
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

func Test_users956(t *testing.T) {
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

func Test_users957(t *testing.T) {
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

func Test_users958(t *testing.T) {
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

func Test_users959(t *testing.T) {
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

func Test_users960(t *testing.T) {
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

func Test_users961(t *testing.T) {
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

func Test_users962(t *testing.T) {
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

func Test_users963(t *testing.T) {
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

func Test_users964(t *testing.T) {
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

func Test_users965(t *testing.T) {
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

func Test_users966(t *testing.T) {
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

func Test_users967(t *testing.T) {
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

func Test_users968(t *testing.T) {
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

func Test_users969(t *testing.T) {
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

func Test_users970(t *testing.T) {
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

func Test_users971(t *testing.T) {
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

func Test_users972(t *testing.T) {
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

func Test_users973(t *testing.T) {
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

func Test_users974(t *testing.T) {
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

func Test_users975(t *testing.T) {
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

func Test_users976(t *testing.T) {
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

func Test_users977(t *testing.T) {
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

func Test_users978(t *testing.T) {
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

func Test_users979(t *testing.T) {
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

func Test_users980(t *testing.T) {
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

func Test_users981(t *testing.T) {
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

func Test_users982(t *testing.T) {
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

func Test_users983(t *testing.T) {
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

func Test_users984(t *testing.T) {
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

func Test_users985(t *testing.T) {
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

func Test_users986(t *testing.T) {
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

func Test_users987(t *testing.T) {
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

func Test_users988(t *testing.T) {
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

func Test_users989(t *testing.T) {
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

func Test_users990(t *testing.T) {
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

func Test_users991(t *testing.T) {
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

func Test_users992(t *testing.T) {
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

func Test_users993(t *testing.T) {
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

func Test_users994(t *testing.T) {
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

func Test_users995(t *testing.T) {
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

func Test_users996(t *testing.T) {
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

func Test_users997(t *testing.T) {
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

func Test_users998(t *testing.T) {
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

func Test_users999(t *testing.T) {
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
