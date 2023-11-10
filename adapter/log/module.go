package log

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("Logger",
		fx.Provide(NewLogger),
	)
}
