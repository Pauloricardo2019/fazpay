package middlewareIntf

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	Auth() gin.HandlerFunc
}

type LogMiddleware interface {
	LogRequest() gin.HandlerFunc
}

type TrackingMiddleware interface {
	TrackRequest() gin.HandlerFunc
}
