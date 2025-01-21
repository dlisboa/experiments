package user

import "context"

// How the context package implements context keys
//
// See:
// https://cs.opensource.google/go/go/+/refs/tags/go1.23.5:src/context/context.go;l=363-377

// Use unexported variables
//
// Here we're hiding the key so the only way to add the value to the context is
// with NewContext, and the only way to get it out is through FromContext
//
// No one else can accidentaly modify the context.

// same as userCtxKey := 0, using a package level var as context key
var userCtxKey int

type User struct{}

// Returns the value from the Context. Only this package knows the key so it's
// the only one that can retrieve it.
//
// `pkgsite` has a convention of ThingFromContext, so this could also
// be called UserIDFromContext.
func FromContext(ctx context.Context) User {
	// Important to use &userCtxKey and not just userCtxKey as userCtxKey is just a
	// 0 int, which can be overwritten with context.WithValue(0, "something"). But
	// by taking the address of this 0 in memory we avoid being overwritten as no
	// other package has access to this variable to take its address.
	value, ok := ctx.Value(&userCtxKey).(User)
	if !ok {
		return User{}
	}
	return value
}

// Returns a new Context with the value. The key is not known to code outside of
// this package.
//
// `pkgsite` has a convention of NewContextWithThing, so this could also
// be called NewContextWithUser.
func NewContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, &userCtxKey, user)
}
