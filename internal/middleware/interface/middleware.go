package middlewareIntf

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	Auth() gin.HandlerFunc
}

type TrackMiddleware interface {
	TrackRequest() gin.HandlerFunc
}
