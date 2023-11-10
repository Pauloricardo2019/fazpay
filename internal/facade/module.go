package facade

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("Facade",
		fx.Provide(
			NewSecurityFacade,
			NewUserFacade,
		),
	)
}
