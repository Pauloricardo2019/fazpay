package main

import (
	"context"
	"github.com/Pauloricardo2019/teste_fazpay/adapter/config"
	"github.com/Pauloricardo2019/teste_fazpay/adapter/database"
	"github.com/Pauloricardo2019/teste_fazpay/internal/repository"
	repositoryIntf "github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	"go.uber.org/fx"
	"os"
)

func main() {
	app := fx.New(
		config.GetModule(),
		database.GetModule(),
		repository.GetModule(),

		fx.Invoke(func(migrator repositoryIntf.Migrator) {
			ctx := context.Background()

			err := migrator.ExecuteMigrations(ctx)
			if err != nil {
				panic(err)
			}
			os.Exit(0)
		}),
	)

	app.Run()
}
