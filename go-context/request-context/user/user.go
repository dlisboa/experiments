package user

import "context"

// Use unexported variables
//
// Here we're hiding the key so the only way to add the value to the context is
// with NewContext, and the only way to get it out is through FromContext
//
// No one else can accidentaly modify the context.

type contextKey string

// could also be:
// type userIDKey struct{}
// using a string is easier if you care to print it
var userIDKey = contextKey("userID")

// Returns the value from the Context. Only this package knows the key so it's
// the only one that can retrieve it.
//
// `pkgsite` has a convention of ThingFromContext, so this could also
// be called UserIDFromContext.
func FromContext(ctx context.Context) int {
	value, ok := ctx.Value(userIDKey).(int)
	if !ok {
		return 0
	}
	return value
}

// Returns a new Context with the value. The key is not known to code outside of
// this package.
//
// `pkgsite` has a convention of NewContextWithThing, so this could also
// be called NewContextWithUserID.
func NewContext(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}
