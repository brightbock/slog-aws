package slog_aws

import (
	"log/slog"
	"os"
)

var (
	AwsLambdaLogLevel slog.Level
)

func LambdaLogger() *slog.Logger {
	err := AwsLambdaLogLevel.UnmarshalText([]byte(
		os.Getenv("AWS_LAMBDA_LOG_LEVEL"),
	))
	if err != nil {
		AwsLambdaLogLevel = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: AwsLambdaLogLevel,
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
