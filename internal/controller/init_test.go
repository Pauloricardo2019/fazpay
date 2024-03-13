package controller_test

import (
	"github.com/Pauloricardo2019/teste_fazpay/adapter/log"
	"github.com/Pauloricardo2019/teste_fazpay/adapter/web"
	"github.com/Pauloricardo2019/teste_fazpay/adapter/web/rest/controller"
	"github.com/Pauloricardo2019/teste_fazpay/adapter/web/rest/middleware"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	"github.com/Pauloricardo2019/teste_fazpay/mocks/facade"
	"github.com/golang/mock/gomock"
	"testing"
)

type Facade struct {
	UserFacadeMock     *facade.MockUserFacade
	SecurityFacadeMock *facade.MockSecurityFacade
}

func setupTestRouter(t *testing.T) (*web.ServerRest, Facade) {
	t.Helper()
	ctrl := gomock.NewController(t)

	facades := Facade{
		UserFacadeMock:     facade.NewMockUserFacade(ctrl),
		SecurityFacadeMock: facade.NewMockSecurityFacade(ctrl),
	}

	cfg := &model.Config{
		EnableSentry: false,
	}
	logger := log.NewLogger(cfg, nil)

	serverRest := web.NewRestServer(
		cfg,
		web.GlobalControllers{
			UserController:        controller.NewUserController(facades.UserFacadeMock),
			LoginController:       controller.NewLoginController(facades.SecurityFacadeMock),
			HealthCheckController: controller.NewHealthCheckController(logger),
		},
		web.Middlewares{
			AuthMiddleware:     middleware.NewAuthMiddleware(facades.SecurityFacadeMock),
			LogMiddleware:      middleware.NewLogMiddleware(logger),
			TrackingMiddleware: middleware.NewTrackingMiddleware(logger),
		},
	)

	return &serverRest, facades
}
