package logger

import (
	"context"
	loggerIntf "github.com/Pauloricardo2019/teste_fazpay/adapter/logger/interface"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	"go.uber.org/zap"
	"time"
)

var Logger *zap.Logger

func init() {
	Logger, _ = zap.NewProduction()
}

type logger struct {
	cfg    *model.Config
	logger *zap.Logger
}

func NewZapLogger(cfg *model.Config) loggerIntf.LoggerInterface {
	return &logger{
		cfg:    cfg,
		logger: Logger,
	}
}

func (l *logger) LoggerInfo(ctx context.Context, infoMessage, layer string) {
	defer l.logger.Sync()

	l.logger.Info("Info",
		zap.String("environment", l.cfg.Environment),
		zap.String("layer", layer),
		zap.String("request_id", ctx.Value("request_id").(string)),
		zap.String("method_request", ctx.Value("method_request").(string)),
		zap.String("url", ctx.Value("request_url").(string)),
		zap.String("message_info", infoMessage),
		zap.Duration("time", time.Second),
	)

}

func (l *logger) LoggerWarn(ctx context.Context, warnMessage, layer string) {
	defer l.logger.Sync()

	l.logger.Warn("Warn",
		zap.String("environment", l.cfg.Environment),
		zap.String("layer", layer),
		zap.String("request_id", ctx.Value("request_id").(string)),
		zap.String("method_request", ctx.Value("method_request").(string)),
		zap.String("url", ctx.Value("request_url").(string)),
		zap.String("warn_message", warnMessage),
		zap.Duration("time", time.Second),
	)
}

func (l *logger) LoggerError(ctx context.Context, errorMessage error, layer string) {
	defer l.logger.Sync()

	l.logger.Error("Error",
		zap.String("environment", l.cfg.Environment),
		zap.String("layer", layer),
		zap.String("request_id", ctx.Value("request_id").(string)),
		zap.String("method_request", ctx.Value("method_request").(string)),
		zap.String("url", ctx.Value("request_url").(string)),
		zap.Error(errorMessage),
		zap.Duration("time", time.Second),
	)
}
