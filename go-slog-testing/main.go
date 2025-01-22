package main

import (
	"log/slog"
)

func Default() int {
	slog.Info("called Default()")
	return 0
}

func Foo(log *slog.Logger) int {
	log.Info("called Foo()")
	return 0
}

func main() {}
