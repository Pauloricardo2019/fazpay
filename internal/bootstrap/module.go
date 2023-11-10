package bootstrap

import (
	"go.uber.org/fx"
	"kickoff/adapter/config"
	"kickoff/adapter/database/repository"
	"kickoff/adapter/log"
	"kickoff/adapter/web"
	"kickoff/adapter/web/rest/controller"
	"kickoff/adapter/web/rest/middleware"
	"kickoff/internal/facade"
	"kickoff/internal/service"
)

func GetModule() fx.Option {

	return fx.Module(
		"Bootstrap",
		config.GetModule(),
		log.GetModule(),
		repository.GetModule(),
		service.GetModule(),
		facade.GetModule(),
		controller.GetModule(),
		middleware.GetModule(),
		fx.Provide(web.NewRestServer),
	)
}
