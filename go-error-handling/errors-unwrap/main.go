package main

import (
	"errors"
	"fmt"
	"os"
)

type ErrUnsupported struct {
	Err error
}

func (e ErrUnsupported) Error() string {
	return fmt.Sprintf("option not supported: %s", e.Err)
}

func (e ErrUnsupported) Unwrap() error {
	return e.Err
}

type ErrNotImplemented struct {
	reason string
	Code   int
}

func (e ErrNotImplemented) Error() string {
	return fmt.Sprintf("not implemented yet: %s", e.reason)
}

// imagine list is like `ls` but some options are not supported
func list(option string) ([]string, error) {
	if option == "-l" {
		err := ErrNotImplemented{
			reason: "too lazy to do it",
			Code:   101,
		}
		// ErrNotImplemented is a put inside an ErrUnsupported
		return []string{}, ErrUnsupported{Err: err}
	}
	return []string{"/boot", "/etc"}, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: fake-ls -a\n  fake-ls -l forces the error\n")
		os.Exit(1)
	}

	result, err := list(os.Args[1])

	// cast err to errUnsupported
	var errUnsupported ErrUnsupported
	if errors.As(err, &errUnsupported) {
		// gets the underlying error which is of type ErrNotImplemented
		unwrapped := errUnsupported.Unwrap()
		fmt.Fprintf(os.Stderr, "fake-ls: error of type ErrUnsupported\nunderlying error that caused it: %T: %v\n", unwrapped, unwrapped)
		os.Exit(2)
	}

	fmt.Printf("%v\n", result)
}
