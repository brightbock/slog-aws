package helpers

import (
	"log/slog"
	"os"
)

func FatalError(s string, v ...any) {
	slog.Error(s, v...)
	os.Exit(1)
}

