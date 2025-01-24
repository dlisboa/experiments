package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var number = flag.Int("n", 10, "how many tests to generate")
var pkg = flag.String("pkg", "main", "package name")

func main() {
	flag.Parse()

	var body string
	switch *pkg {
	case "users":
		body = usertest
	case "movies":
		body = movietest
	}

	t := template.Must(template.New("").Parse(fmt.Sprintf(`
func Test_{{.Package}}{{.TestNumber}}(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}
%s
}
`, body)))

	fmt.Fprintf(os.Stdout, `// Generated with tools/gentests. DO NOT EDIT.

//go:build integration

package %s

import (
	"testing"

	"demo/internal/testdb"
)
`, *pkg)

	type data struct {
		Package    string
		TestNumber int
	}
	for i := range *number {
		t.Execute(os.Stdout, data{Package: *pkg, TestNumber: i})
	}
}

var usertest = `
	// seeded user
	john := "John Lennon"
	user, err := store.GetByName(t.Context(), john)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, john; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	george := "George Harrison"
	err = store.Insert(t.Context(), george)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), george)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	ringo := "Ringo Starr"
	_, err = store.GetByName(t.Context(), ringo)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}
`

var movietest = `
	// seeded movie
	titanic := "Titanic"
	user, err := store.GetByName(t.Context(), titanic)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, titanic; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	avatar := "Avatar"
	err = store.Insert(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), avatar)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	terminator := "Terminator"
	_, err = store.GetByName(t.Context(), terminator)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}
`
