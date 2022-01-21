package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pex/fibonacci/pkg/config"
)

// NewEngine creates an instance of the gin engine
func NewEngine(
	config *config.ServerConfig,
	setHandlerRoutes SetHandlerRoutesFunc,
) *gin.Engine {
	// Create the new Gin engine and setup middleware handler chain

	ge := gin.New()
	ge.Use(RequestIDMiddleware())
	ge.Use(HostnameMiddleware())
	ge.Use(gin.Recovery())

	// root path to return OK response code
	ge.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Initialize routes and add to routing group
	setHandlerRoutes(ge, config.ServiceName, config.Version, config.BasePath)
	// setSwaggerRoutes(ge, config.SwaggerPath)
	return ge
}
