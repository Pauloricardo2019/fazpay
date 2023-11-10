package main

import (
	"context"
	"go.uber.org/fx"
	"kickoff/adapter/config"
	"kickoff/adapter/database/repository"
	repositoryIntf "kickoff/adapter/database/repository/interface"
	"kickoff/adapter/log"
	logIntf "kickoff/adapter/log/interface"
	"os"
)

func main() {
	app := fx.New(
		config.GetModule(),
		log.GetModule(),
		repository.GetModule(),
		fx.Invoke(func(migrator repositoryIntf.Migrator, logger logIntf.Logger) {
			ctx := context.Background()
			logger.Info(ctx, "Starting migration program")

			err := migrator.ExecuteMigrations(ctx)
			if err != nil {
				logger.Error(ctx, "error running migrations",
					"error", err.Error())
			}
			os.Exit(0)
		}),
	)

	app.Run()
}
