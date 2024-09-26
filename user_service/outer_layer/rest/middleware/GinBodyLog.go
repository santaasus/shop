package middleware

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(byte []byte) (int, error) {
	w.body.Write(byte)
	return w.ResponseWriter.Write(byte)
}

func GinBodyLogMiddleware(c *gin.Context) {
	bwr := &responseBodyWriter{
		ResponseWriter: c.Writer,
		body:           &bytes.Buffer{},
	}
	c.Writer = bwr
	c.Next()

	bodyInfo := map[string]any{
		"full_path":     c.FullPath(),
		"uri_path":      c.Request.RequestURI,
		"request_body":  c.Request.Body,
		"status_code":   c.Writer.Status(),
		"response_body": bwr.body.String(),
		"errors":        c.Errors.Errors(),
		"created_at":    time.Now().Format(time.DateTime),
	}

	fmt.Printf("server data info %v\n", bodyInfo)
}
