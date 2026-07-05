// Package logger configures structured application logging.
package logger

import (
	"io"
	"log/slog"
)

// New creates a text logger suitable for CLI output.
func New(w io.Writer, verbose bool) *slog.Logger {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: level}))
}
