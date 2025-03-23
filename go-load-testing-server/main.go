package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

const responseSize = 500 * 1024 // 500KB

var content = randomString(responseSize - 100)

func handler(w http.ResponseWriter, r *http.Request) {
	content := randomString(responseSize - 100) // leave space for HTML tags

	html := fmt.Sprintf("<!DOCTYPE html><html><head><title>Random Page</title></head><body><h1>%s</h1></body></html>", content)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
