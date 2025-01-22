package main

import (
	"demo/middleware"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mw := middleware.Chain(
		middleware.Log(logger),
		middleware.Timeout(2*time.Second),
		middleware.RequestSize(1024),
		middleware.Recover(),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("POST /limited", limited)
	mux.HandleFunc("GET /panic", handlePanic)
	mux.HandleFunc("GET /timeout", timeout)
	srv := http.Server{
		Addr:    ":8080",
		Handler: mw(mux),
	}

	fmt.Println("listening on :8080")
	srv.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
}

func limited(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	w.Write([]byte(fmt.Sprintf("size limited request size: %d", len(body))))
}

func handlePanic(w http.ResponseWriter, r *http.Request) {
	panic("panic from the handler")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
}
