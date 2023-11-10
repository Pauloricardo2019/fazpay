package middleware

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	logIntf "kickoff/adapter/log/interface"
	middlewareIntf "kickoff/adapter/web/rest/middleware/interface"
)

type TrackingMiddleware struct {
	logger logIntf.Logger
}

func NewTrackingMiddleware(logger logIntf.Logger) middlewareIntf.TrackingMiddleware {
	return &TrackingMiddleware{
		logger: logger,
	}
}

func (l *TrackingMiddleware) TrackRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		newUUID, _ := uuid.NewV4()
		ctx := context.WithValue(c.Request.Context(), "kickoff-request-id", newUUID.String())
		c.Request = c.Request.WithContext(ctx)

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		w.Header().Add("kickoff-request-id", newUUID.String())
		c.Next()
	}
}
