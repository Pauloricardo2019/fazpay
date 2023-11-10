package testing

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"kickoff/adapter/web"
	"kickoff/internal/bootstrap"
)

var serverRest web.ServerRest

func GetServer() (*gin.Engine, error) {
	return serverRest.Engine, nil
}

func StartServer() error {
	err := fx.
		New(
			bootstrap.GetModule(),
			fx.Invoke(func(server web.ServerRest) {
				serverRest = server
			}),
		).
		Start(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func ShutdownServer() error {
	err := fx.
		New().
		Stop(context.Background())

	if err != nil {
		return err
	}

	return nil
}
