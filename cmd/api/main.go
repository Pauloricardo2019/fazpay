package main

import (
	"context"
	"go.uber.org/fx"
	"gorm.io/gorm"
	repositoryIntf "kickoff/adapter/database/repository/interface"
	logIntf "kickoff/adapter/log/interface"
	"kickoff/adapter/web"
	"kickoff/internal/bootstrap"
)

// @contact.name				API Support
// @contact.url					http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url					http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService				http://swagger.io/terms/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	fx.New(
		bootstrap.GetModule(),

		fx.Invoke(func(logger logIntf.Logger, server web.ServerRest, db *gorm.DB, migrator repositoryIntf.Migrator, lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					logger.Info(ctx, "Starting migrator...")
					err := migrator.ExecuteMigrations(ctx)
					if err != nil {
						return err
					}
					logger.Info(ctx, "Migrator has executed successfully")

					logger.Info(ctx, "Starting server...")
					// As the startListener is blocking, we need to start it in a separated goroutine
					go server.StartListener()

					logger.Info(ctx, "Server started.")
					return nil
				},
				OnStop: func(ctx context.Context) error {
					logger.Info(ctx, "Stopping server...")
					err := server.StopListener(ctx)
					return err
				},
			})
		}),
	).Run()
}
