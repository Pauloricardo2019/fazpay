package log_test

import (
	"context"
	"kickoff/adapter/log"
	"kickoff/internal/model"
	"testing"
)

func TestLogInfo(t *testing.T) {
	cfg := &model.Config{
		Environment: "dev",
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "kickoff-request-id", "new value")

	logger := log.NewLogger(cfg, nil)
	logger.Info(ctx, "Test with value",
		"value", 1,
	)

	logger.Info(ctx, "Test without value")
}
