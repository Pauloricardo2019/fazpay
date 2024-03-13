package controllerIntf

import "github.com/gin-gonic/gin"

type HealthCheckController interface {
	HealthCheck(c *gin.Context)
}

type UserController interface {
	CreateUser(c *gin.Context)
	GetByIdUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type LoginController interface {
	Login(c *gin.Context)
}
