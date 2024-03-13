package middleware

import (
	"bytes"
	middlewareIntf "github.com/Pauloricardo2019/teste_fazpay/internal/middleware/interface"
	"github.com/gin-gonic/gin"
)

type LogMiddleware struct{}

func (l LogMiddleware) LogRequest() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func NewLogMiddleware() middlewareIntf.LogMiddleware {
	return &LogMiddleware{}
}
