package bootstrap

import (
	"github.com/Pauloricardo2019/teste_fazpay/adapter/config"
	"github.com/Pauloricardo2019/teste_fazpay/adapter/database"
	"github.com/Pauloricardo2019/teste_fazpay/adapter/web"
	"github.com/Pauloricardo2019/teste_fazpay/internal/controller"
	"github.com/Pauloricardo2019/teste_fazpay/internal/facade"
	"github.com/Pauloricardo2019/teste_fazpay/internal/middleware"
	"github.com/Pauloricardo2019/teste_fazpay/internal/repository"
	"github.com/Pauloricardo2019/teste_fazpay/internal/service"
	"go.uber.org/fx"
)

func GetModule() fx.Option {

	return fx.Module(
		"Bootstrap",
		config.GetModule(),
		database.GetModule(),
		repository.GetModule(),
		service.GetModule(),
		facade.GetModule(),
		controller.GetModule(),
		middleware.GetModule(),
		fx.Provide(web.NewRestServer),
	)
}
