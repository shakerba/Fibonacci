package server

import (
	"os"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestIDMiddleware ...
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a new request ID if one is not already set on
		// the incoming header.
		var reqID string
		if reqID = c.Request.Header.Get("X-Request-ID"); reqID == "" {
			reqID = Prefix + uuid.NewV4().String()
		}
		c.Set("reqID", reqID)
		c.Writer.Header().Add("X-Request-ID", reqID)
	}
}

// HostnameMiddleware ...
func HostnameMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, _ := os.Hostname()
		c.Writer.Header().Add("X-Hostname", name)
		c.Next()
	}
}

// LoggingMiddleware ...
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Start timer
		start := time.Now()

		// Log the incoming request
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		agent :=  c.Request.UserAgent()
		reqID := c.GetString("reqID")

		fmt.Printf("incoming request, REQ-ID: %v, CLIENT_IP: %v, USER_AGENT: %v, METHOD: %v, PATH: %v\n", reqID, clientIP, agent, method, path)
		// Measure time and process request
		c.Next()
		latency := time.Since(start)

		fmt.Printf("request complete, REQ-ID: %v, CLIENT_IP: %v, USER_AGENT: %v, METHOD: %v, PATH: %v, STATUS: %v, LATENCY: %v\n", reqID, clientIP, agent, method, path, c.Writer.Status(), latency.Seconds())
}
}
