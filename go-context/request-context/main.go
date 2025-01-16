package main

import (
	"context"
	"fmt"
	"net/http"
)

// Use a custom type to avoid key collision.
//
// If it was simply a string anyone would be able to overwrite it.
//
// For better protection create a different package and never export the key,
// using methods like `NewContextWithThing/ThingFromContext` to manipulate the
// context.
//
// See user/user.go for an example.
type contextKey string

// could also be:
// type userIDKey struct{}
// using a string is easier if you care to print it
var userIDKey = contextKey("userID")

func main() {
	http.Handle("GET /", middleware(http.HandlerFunc(handler)))
	fmt.Println("make request to :8080/")
	http.ListenAndServe(":8080", nil)
}

func middleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), userIDKey, 123)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func handler(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(userIDKey)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}
