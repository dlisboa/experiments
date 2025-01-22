package main

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func main() {
	basic := slog.New(tint.NewHandler(os.Stdout, nil))
	basic.Info("this is info", "int", 1, "string", "a string")
	basic.Error("this is error", "int", 1, "string", "a string")

	// set global logger with custom options. See tint.Options.
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			AddSource:  true,
			Level:      slog.LevelDebug,
			TimeFormat: time.RFC3339,
		}),
	))

	slog.Info("this is info", "int", 1, "string", "a string")
	slog.Error("this is error", "int", 1, "string", "a string")

	// prints all values of type error in red
	errlogger := slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if err, ok := a.Value.Any().(error); ok {
					aErr := tint.Err(err)
					aErr.Key = a.Key
					return aErr
				}
				return a
			},
		}),
	)
	errlogger.Error("this is error", "err", errors.New("an error"))
}
