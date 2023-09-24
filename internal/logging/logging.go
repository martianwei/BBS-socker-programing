package logging

import (
	"log/slog"
	"os"
)

func CreateLogger(logLevel slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	return slog.New(handler)
}
