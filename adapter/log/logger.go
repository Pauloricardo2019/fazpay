package log

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	logIntf "kickoff/adapter/log/interface"
	"kickoff/internal/model"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func createLogger(logLevel zapcore.Level) *zap.Logger {
	var logger *zap.Logger
	var err error

	switch logLevel {
	case zapcore.InfoLevel:
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err.Error())
	}

	return logger
}

func NewLogger(config *model.Config, lc fx.Lifecycle) logIntf.Logger {
	var logger *zap.Logger

	env := config.Environment
	if env == "dev" {
		logger = createLogger(zapcore.DebugLevel)
	} else {
		logger = createLogger(zapcore.InfoLevel)
	}

	sugaredLogger := logger.Sugar()

	instance := &Logger{
		logger: sugaredLogger,
	}

	if lc != nil {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				_ = instance.Sync()
				return nil
			},
		})
	}

	return instance
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}

func (l *Logger) Debug(ctx context.Context, message string, args ...interface{}) {
	args = l.enrichData(ctx, args...)
	l.logger.Debugw(message, args...)
}

func (l *Logger) Info(ctx context.Context, message string, args ...interface{}) {
	args = l.enrichData(ctx, args...)
	l.logger.Infow(message, args...)
}

func (l *Logger) Warn(ctx context.Context, message string, args ...interface{}) {
	args = l.enrichData(ctx, args...)
	l.logger.Warnw(message, args...)
}

func (l *Logger) Error(ctx context.Context, message string, args ...interface{}) {
	args = l.enrichData(ctx, args...)
	l.logger.Errorw(message, args...)
}

func (l *Logger) Fatal(ctx context.Context, message string, args ...interface{}) {
	args = l.enrichData(ctx, args...)
	l.logger.Fatalw(message, args...)
}

func (l *Logger) enrichData(ctx context.Context, args ...interface{}) []interface{} {
	requestID := ctx.Value("kickoff-request-id")
	if args == nil {
		args = make([]interface{}, 0)
	}
	if requestID != nil {
		args = append(args, "kickoff-request-id", requestID)
	}

	return args
}
