package middleware

import (
	middlewareIntf "github.com/Pauloricardo2019/teste_fazpay/internal/middleware/interface"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type trackMiddleware struct{}

func NewTrackMiddleware() middlewareIntf.TrackMiddleware {
	return &trackMiddleware{}
}

func (t trackMiddleware) TrackRequest() gin.HandlerFunc {
	return func(c *gin.Context) {

		urlRequest := c.Request.URL.String()
		c.Set("url_request", urlRequest)

		methodRequest := c.Request.Method
		c.Set("method_request", methodRequest)

		loggerUUID := uuid.New().String()
		c.Set("request_id", loggerUUID)

		c.Next()

	}
}
