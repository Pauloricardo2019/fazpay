package service

import "go.uber.org/fx"

func GetModule() fx.Option {

	return fx.Module("Service",
		fx.Provide(
			NewTokenService,
			NewUserService,
		),
	)

}
