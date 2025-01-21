package main

import (
	"context"
	"fmt"
	"net/http"
)

// Use a custom type to avoid key collision
//
// For better protection create a different package and never export the key,
// using methods like `NewContextWithThing/ThingFromContext` to manipulate the
// context.
//
// See user/user.go for an example.

type contextKey int

// If there were other context keys increase the integer value so no two keys
// have the same value. See user/user.go for a slightly easier solution to
// multiple keys.
var userCtxKey contextKey = 0

func main() {
	http.Handle("GET /", middleware(http.HandlerFunc(handler)))
	fmt.Println("make request to :8080/")
	http.ListenAndServe(":8080", nil)
}

func middleware(h http.Handler) http.Handler {
	type User struct{}

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), userCtxKey, User{})
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userCtxKey)
	w.Write([]byte(fmt.Sprintf("user: %#+v", user)))
}
