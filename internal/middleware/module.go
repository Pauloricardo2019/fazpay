package middleware

import "go.uber.org/fx"

func GetModule() fx.Option {

	return fx.Module(
		"Middleware",
		fx.Provide(
			NewAuthMiddleware,
			NewLogMiddleware,
		),
	)

}
