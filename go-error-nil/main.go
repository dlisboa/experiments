package main

import (
	"fmt"
)

type E struct{}

func (*E) Error() string { return "e: some error" }

// See https://go.dev/doc/faq#nil_error
//
// Under the covers, interfaces are implemented as two elements, a type T and
// a value V. V is a concrete value such as an int, struct or pointer, never an
// interface itself, and has type T. For instance, if we store the int value 3
// in an interface, the resulting interface value has, schematically, (T=int,
// V=3). The value V is also known as the interface’s dynamic value, since a
// given interface variable might hold different values V (and corresponding
// types T) during the execution of the program.
//
// An interface value is nil only if the V and T are both unset, (T=nil, V is
// not set), In particular, a nil interface will always hold a nil type. If we
// store a nil pointer of type *int inside an interface value, the inner type
// will be *int regardless of the value of the pointer: (T=*int, V=nil). Such
// an interface value will therefore be non-nil even when the pointer value V
// inside is nil.
func returnsError() error {
	var e *E = nil
	return e // Will always return a non-nil error.
}

func main() {
	err := returnsError()                      // err is an interface value
	fmt.Printf("err == nil: %v\n", err == nil) // will print false

	var err2 *E = nil                            // err2 is not an interface
	fmt.Printf("err2 == nil: %v\n", err2 == nil) // will print true
}
