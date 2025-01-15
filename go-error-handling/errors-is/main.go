package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrUnsupported = errors.New("option not supported")

// imagine list is like `ls` but some options are not supported
func list(option string) ([]string, error) {
	if option == "-l" {
		return []string{}, fmt.Errorf("fake-ls: %w", ErrUnsupported)
	}
	return []string{"/boot", "/etc"}, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: fake-ls -a\n  fake-ls -l forces the error\n")
		os.Exit(1)
	}

	result, err := list(os.Args[1])

	// since err is not exactly ErrUnsupported, but an error that wraps it,
	// this comparison will fail and this `if` never enters
	if err == ErrUnsupported {
		fmt.Fprintln(os.Stderr, "unreachable", err)
		os.Exit(2)
	}

	// this comparison will return true, `if` enters
	if errors.Is(err, ErrUnsupported) {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	fmt.Printf("%v\n", result)
}
