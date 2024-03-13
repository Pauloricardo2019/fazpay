package database

import (
	"github.com/Pauloricardo2019/teste_fazpay/adapter/database/sql"
	"go.uber.org/fx"
)

func GetModule() fx.Option {

	return fx.Module(
		"Database",
		fx.Provide(
			sql.NewDb,
		),
	)

}
