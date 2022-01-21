package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pex/fibonacci/pkg/handlers"
)

// SetDefaultRoutes wires up routing to various handlers
func SetDefaultRoutes(
	ginEngine *gin.Engine,
	serviceName string,
	version string,
	basePath string,
) {

	router := handlers.RouteHandler{
		ServiceName: serviceName,
		Version:     version,
	}

	// swagger:route GET /current status GetCurrent
	ginEngine.GET("/current", router.GetCurrent)

	// swagger:route GET /next status GetNext
	ginEngine.GET("/next", router.GetNext)

	// swagger:route GET /previous status GetPrevious
	ginEngine.GET("/previous", router.GetPrevious)
}

func setSwaggerRoutes(engine *gin.Engine, swaggerPath string) {
	engine.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger-api/")
	})

	engine.GET("/swagger-api-json/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger-api/swagger.json")
	})

	// In the docker container, the etc directory from GHE repo is copied to /etc
	// directory in local filesystem. Reference that directory here.
	engine.Static("/swagger-api/", swaggerPath)
}
