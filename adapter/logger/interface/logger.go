package loggerIntf

import "context"

type LoggerInterface interface {
	LoggerInfo(ctx context.Context, infoMessage, layer string)
	LoggerWarn(ctx context.Context, warnMessage, layer string)
	LoggerError(ctx context.Context, errorMessage error, layer string)
}
