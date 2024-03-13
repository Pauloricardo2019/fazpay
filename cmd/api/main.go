package main

import (
	"context"
	"github.com/Pauloricardo2019/teste_fazpay/adapter/web"
	"github.com/Pauloricardo2019/teste_fazpay/internal/bootstrap"
	repositoryIntf "github.com/Pauloricardo2019/teste_fazpay/internal/repository/interface"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
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

		fx.Invoke(func(server web.ServerRest, db *gorm.DB, migrator repositoryIntf.Migrator, lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {

					if err := migrator.ExecuteMigrations(ctx); err != nil {
						log.Fatal(err.Error())
					}

					go server.StartListener()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					err := server.StopListener(ctx)
					return err
				},
			})
		}),
	).Run()
}
