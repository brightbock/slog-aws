package jsonlogger

import (
	"io"
	"log/slog"
)

func Logger(w io.Writer, level string) *slog.Logger {
	var LogLevel slog.Level
	err := LogLevel.UnmarshalText([]byte(level))
	if err != nil {
		LogLevel = slog.LevelInfo
	}
	logger := slog.New(slog.NewJSONHandler(
		w,
		&slog.HandlerOptions{
			Level: LogLevel,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// Customize the name of the time key
				if a.Key == slog.TimeKey {
					a.Key = "timestamp"
				}
				return a
			},
		},
	))
	return logger
}
