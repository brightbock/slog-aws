package tocloudwatch

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/brightbock/slogcloud/cwlogger"
	"github.com/brightbock/slogcloud/jsonlogger"
	"log/slog"
	"os"
	"time"
)

type Timestamp struct {
	RFC3339Nano string `json:"timestamp"`
}

type ExtractTimeWriter struct {
	CWLogger *cwlogger.Logger
}

type Config struct {
	LogGroupName  string
	Client        *cloudwatchlogs.Client
	ErrorReporter func(error)
	Closer        func()
}

func (w *ExtractTimeWriter) Write(p []byte) (n int, err error) {
	var timeFromJSON Timestamp
	err = json.Unmarshal(p, &timeFromJSON)
	if err != nil {
		return 0, err
	}
	t, err := time.Parse(time.RFC3339Nano, timeFromJSON.RFC3339Nano)
	w.CWLogger.Log(t, string(bytes.TrimSpace(p)))
	return len(p), nil
}

func Logger(config *Config) (*slog.Logger, error) {
	cwl, err := cwlogger.New(&cwlogger.Config{
		LogGroupName:  config.LogGroupName,
		Client:        config.Client,
		ErrorReporter: config.ErrorReporter,
	})
	if err != nil {
		return nil, err
	}
	config.Closer = func() {
		cwl.Close()
	}
	extractTimeWriter := &ExtractTimeWriter{
		CWLogger: cwl,
	}
	logger := jsonlogger.Logger(
		extractTimeWriter,
		os.Getenv("LOG_LEVEL"),
	)
	return logger, nil
}

func LogToCloudwatch(config *Config) error {
	logger, err := Logger(config)
	if err != nil {
		return err
	}
	slog.SetDefault(logger)
	originalCloser := config.Closer
	config.Closer = func() {
		defer originalCloser()
		time.Sleep(1 * time.Second)
		slog.Debug("Closing Cloudwatch Logging Handle")
		logger := jsonlogger.Logger(
			os.Stdout,
			os.Getenv("LOG_LEVEL"),
		)
		slog.SetDefault(logger)
	}
	return nil
}
