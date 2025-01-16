package user

import "context"

// Use unexported variables
//
// Here we're hiding the key so the only way to add the value to the context is
// with IntoContext, and the only way to get it out is through FromContext
//
// No one else can accidentaly modify the context.

type contextKey string

var userIDKey = contextKey("userID")

// Returns the value from the Context. Only this package knows the key so it's
// the only one that can retrieve it.
func FromContext(ctx context.Context) int {
	value, ok := ctx.Value(userIDKey).(int)
	if !ok {
		return 0
	}
	return value
}

// Returns a new Context with the value. The key is not known to code outside of
// this package.
func IntoContext(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}
