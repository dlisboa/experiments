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
		return []string{}, ErrUnsupported
	}
	return []string{"/boot", "/etc"}, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: fake-ls -a\n  fake-ls -l forces the error\n")
		os.Exit(1)
	}

	result, err := list(os.Args[1])
	// comparison with sentinel error
	if err == ErrUnsupported {
		fmt.Fprintf(os.Stderr, "fake-ls: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("%v\n", result)
}
