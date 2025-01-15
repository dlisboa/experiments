package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrUnsupported = errors.New("option not supported")

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
		// an error that is both ErrUnsupported and ErrNotImplemented
		return []string{}, fmt.Errorf("%w: %w", ErrUnsupported, err)
	}
	return []string{"/boot", "/etc"}, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: fake-ls -a\n  fake-ls -l forces the error\n")
		os.Exit(1)
	}

	result, err := list(os.Args[1])

	// this does not compile since ErrNotImplemented is a type and not
	// an expression
	// if err == ErrNotImplemented {
	// 	// ...
	// }

	// this also does not compile, ErrNotImplemented is not an expression
	// if errors.Is(err, ErrNotImplemented) {
	// 	// ...
	// }

	// this comparison will work
	// casts err to notImplementedErr, skipping past the other errors in the chain,
	// and giving us access to things like struct fields
	var notImplementedErr ErrNotImplemented
	if errors.As(err, &notImplementedErr) {
		fmt.Fprintf(
			os.Stderr,
			"fake-ls: %T: %v. Error code: %d\n",
			notImplementedErr,
			notImplementedErr,
			notImplementedErr.Code,
		)

		// err is still ErrUnsupported, we just skipped past it when we created
		// notImplementedErr with errors.As
		if errors.Is(err, ErrUnsupported) {
			fmt.Fprintf(
				os.Stderr,
				"fake-ls: error is main.ErrUnsupported as well\n",
			)
		}
		os.Exit(2)
	}

	fmt.Printf("%v\n", result)
}
