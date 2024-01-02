package slogcloud

import (
	"github.com/brightbock/slogcloud/jsonlogger"
	"log/slog"
	"os"
)

func Logger() *slog.Logger {
	logger := jsonlogger.Logger(
		os.Stdout,
		os.Getenv("AWS_LAMBDA_LOG_LEVEL"),
	)
	return logger
}
