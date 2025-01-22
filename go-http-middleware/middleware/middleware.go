// Taken from: https://github.com/golang/telemetry/blob/04cd7b/godev/internal/middleware/middleware.go
// With small modifications.

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package middleware implements a simple middleware pattern for http handlers,
// along with implementations for some common middlewares.
package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

// A Middleware is a func that wraps an http.Handler.
type Middleware func(http.Handler) http.Handler

// Chain creates a new Middleware that applies a sequence of Middlewares, so
// that they execute in the given order when handling an http request.
//
// In other words, Chain(m1, m2)(handler) = m1(m2(handler))
//
// A similar pattern is used in e.g. github.com/justinas/alice:
// https://github.com/justinas/alice/blob/ce87934/chain.go#L45
func Chain(middlewares ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := range middlewares {
			h = middlewares[len(middlewares)-1-i](h)
		}
		return h
	}
}

// Log is a middleware that logs request start, end, duration, and status.
func Log(logger *slog.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ctx := r.Context()
			l := logger.With(
				slog.String("method", r.Method),
				slog.String("uri", r.RequestURI),
				// TODO(hyangah): set trace context from X-Cloud-Trace-Context
			)
			l.InfoContext(ctx, "request start")
			w2 := &statusRecorder{w, 200}
			h.ServeHTTP(w2, r)
			level := slog.LevelInfo
			msg := "request end"
			switch w2.status / 100 {
			case 5:
				level = slog.LevelError // 5XX error
				msg = "request error"
			case 4:
				level = slog.LevelWarn // 4XX error
				msg = "request rejected"
			}
			l.Log(ctx, level, msg,
				slog.Int("status", w2.status),
				slog.Int64("duration_µs", time.Since(start).Microseconds()),
			)
		})
	}
}

// Recover is a middleware that recovers from panics in the delegate
// handler and prints a stack trace.
func Recover() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					slog.Error(r.RequestURI, "err", fmt.Errorf(`panic("%s")`, err))
					fmt.Println(string(debug.Stack()))
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}

// RequestSize limits the size of incoming request bodies, in bytes.
func RequestSize(n int64) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, n)
			h.ServeHTTP(w, r)
		})
	}
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

// Timeout returns a new Middleware that times out each request after the given
// duration.
func Timeout(d time.Duration) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.TimeoutHandler(h, d, "request timed out").ServeHTTP(w, r)
		})
	}
}
