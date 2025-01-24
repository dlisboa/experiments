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
	type data struct {
		TestNumber int
	}

	t := template.Must(template.New("").Parse(`
func Test{{.TestNumber}}(t *testing.T) {
	t.Parallel()

	db := testdb.DB()
	t.Cleanup(db.Teardown)

	store := &Store{db: db}

	// seeded user
	john := "John Lennon"
	user, err := store.GetByName(t.Context(), john)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := user.Name, john; got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}

	// insert and retrieve
	george := "George Harrison"
	err = store.Insert(t.Context(), george)
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetByName(t.Context(), george)
	if err != nil {
		t.Fatal(err)
	}

	// retrieve nothing
	ringo := "Ringo Starr"
	_, err = store.GetByName(t.Context(), ringo)
	if err == nil {
		// found something
		t.Errorf("got=%v, want=%v", err, nil)
	}
}
`))

	flag.Parse()
	fmt.Fprintf(os.Stdout, `// Generated with tools/gentests. DO NOT EDIT.

//go:build integration

package %s

import (
	"testing"

	"demo/internal/testdb"
)
`, *pkg)

	for i := range *number {
		t.Execute(os.Stdout, data{i})
	}
}
