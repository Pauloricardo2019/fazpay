package controller

import (
	"go.uber.org/fx"
)

func GetModule() fx.Option {
	return fx.Module(
		"Controllers",
		fx.Provide(
			NewHealthCheckController,
			NewLoginController,
			NewUserController,
		),
	)
}
