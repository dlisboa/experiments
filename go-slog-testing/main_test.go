package main

import (
	"log/slog"
	"os"
	"testing"

	"github.com/neilotoole/slogt"
)

func TestMain(m *testing.M) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	m.Run()
}

func TestFoo(t *testing.T) {
	t.Log("Using slogt.New(t)")
	log := slogt.New(t)

	if got, want := Foo(log), 0; got != want {
		t.Errorf("Foo()=%v, want=%v", got, want)
	}
}

func TestDefault(t *testing.T) {
	t.Log("Using slog.Default")

	if got, want := Default(), 0; got != want {
		t.Errorf("Default()=%v, want=%v", got, want)
	}
}
