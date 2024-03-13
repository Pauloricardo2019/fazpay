package repository

import (
	"go.uber.org/fx"
)

func GetModule() fx.Option {

	return fx.Module(
		"Repository",
		fx.Provide(
			NewBaseRepository,
			NewTokenRepository,
			NewUserRepository,
			NewMigrator,
		),
	)

}
