package web

import (
	"context"
	"errors"
	"fmt"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	controller "kickoff/adapter/web/rest/controller/interface"
	middlewareIntf "kickoff/adapter/web/rest/middleware/interface"
	"kickoff/docs"
	"kickoff/internal/model"
	"net/http"
	"time"
)

type GlobalControllers struct {
	fx.In

	UserController        controller.UserController
	HealthCheckController controller.HealthCheckController
	LoginController       controller.LoginController
}

type Middlewares struct {
	fx.In

	AuthMiddleware     middlewareIntf.AuthMiddleware
	LogMiddleware      middlewareIntf.LogMiddleware
	TrackingMiddleware middlewareIntf.TrackingMiddleware
}

type ServerRest struct {
	httpServer  *http.Server
	Engine      *gin.Engine
	config      *model.Config
	controllers *GlobalControllers
	middlewares *Middlewares
}

func UseCORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		origin := c.GetHeader("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin, apiKey")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}

func NewRestServer(
	cfg *model.Config,
	controllers GlobalControllers,
	middlewares Middlewares) ServerRest {

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(UseCORS())

	docs.SwaggerInfo.Title = "kickoff - API"
	docs.SwaggerInfo.Description = "API para comunicação com o sistema kickoff"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	docs.SwaggerInfo.BasePath = cfg.BasePath

	server := ServerRest{
		Engine:      engine,
		config:      cfg,
		controllers: &controllers,
		middlewares: &middlewares,
	}

	server.registerRoutes()

	return server
}

func (s *ServerRest) registerRoutes() {

	if s.config.EnableSentry {
		s.Engine.Use(sentryGin.New(sentryGin.Options{
			WaitForDelivery: false,
		}))
	}

	basePath := s.Engine.Group(s.config.BasePath, s.middlewares.TrackingMiddleware.TrackRequest(), s.middlewares.LogMiddleware.LogRequest())
	{
		basePath.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		basePath.GET("/health", s.controllers.HealthCheckController.HealthCheck)

		routeV1 := basePath.Group("/v1")
		{
			userGroup := routeV1.Group("user")
			{
				userGroup.POST("/", s.controllers.UserController.CreateUser)
				userGroup.GET("/:id", s.middlewares.AuthMiddleware.Auth(), s.controllers.UserController.GetByIdUser)
				userGroup.PUT("/:id", s.middlewares.AuthMiddleware.Auth(), s.controllers.UserController.UpdateUser)
				userGroup.DELETE("/:id", s.middlewares.AuthMiddleware.Auth(), s.controllers.UserController.DeleteUser)
			}

			authGroup := routeV1.Group("auth")
			{
				authGroup.POST("/login", s.controllers.LoginController.Login)
			}

		}
	}

}

func (s *ServerRest) StartListener() {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.RestPort),
		Handler: s.Engine,
	}

	fmt.Println("Listening on port", s.config.RestPort)
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err.Error())
	}
}

func (s *ServerRest) StopListener(ctx context.Context) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctxWithTimeout)
	if err != nil {
		return err
	}

	fmt.Println("http server was gracefully stopped")
	return nil
}
