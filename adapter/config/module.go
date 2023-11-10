package config

import "go.uber.org/fx"

func GetModule() fx.Option {

	return fx.Module(
		"Configuration",
		fx.Provide(
			GetConfig,
		),
	)
}
