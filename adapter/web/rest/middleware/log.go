package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	logIntf "kickoff/adapter/log/interface"
	middlewareIntf "kickoff/adapter/web/rest/middleware/interface"
	"strings"
)

type LogMiddleware struct {
	logger logIntf.Logger
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func NewLogMiddleware(logger logIntf.Logger) middlewareIntf.LogMiddleware {
	return &LogMiddleware{
		logger: logger,
	}
}

func (l *LogMiddleware) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		l.internalLogRequest(c)

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		if strings.HasPrefix(c.Request.RequestURI, "/swagger") {
			return
		}

		l.logger.Info(c.Request.Context(), "LogMiddleware: Response",
			"status", w.Status(),
			"body", w.body.String(),
			"headers", w.Header(),
			"keys", c.Keys)
	}
}

func (l *LogMiddleware) internalLogRequest(c *gin.Context) {

	if strings.HasPrefix(c.Request.RequestURI, "/swagger") {
		return
	}

	var requestBody string
	if c.Request.Body != nil {
		// if we have a body, let's log it
		buf, _ := io.ReadAll(c.Request.Body)
		reader := io.NopCloser(bytes.NewBuffer(buf))
		backupReader := io.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because reader will be read.

		requestBody = l.readBody(reader)
		c.Request.Body = backupReader
	}

	l.logger.Info(c.Request.Context(), "LogMiddleware: Request",
		"body", requestBody,
		"requestURI", c.Request.RequestURI,
		"host", c.Request.Host,
		"method", c.Request.Method,
		"headers", c.Request.Header)
}

func (l *LogMiddleware) readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	result := buf.String()
	return result
}
