package middleware

import "github.com/gin-gonic/gin"

func CommonHeaders(c *gin.Context) {
	accessControllPrefix := "Access-Control-Allow-"
	c.Header(accessControllPrefix+"Origin", "*")
	c.Header(accessControllPrefix+"Credentials", "true")
	c.Header(accessControllPrefix+"Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header(accessControllPrefix+"Headers",
		"Content-Type, UserName-Agent, If-Modified-Since, Cache-Control")
	c.Header("X-Frame-Options", "AMEORIGIN")
	c.Header("Cache-Control", "no-cache, no-store")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
}
