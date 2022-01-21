package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pex/fibonacci/pkg/config"
	"github.com/pex/fibonacci/pkg/models"
)

func (NextHandler *RouteHandler) GetNext(c *gin.Context) {
	resp := &models.NextResponse{Number: config.C,}

config.Lock.Lock()

if config.B == 0 {
	config.B = 1
	config.C = 1

	resp.Number = config.C
} else {
	config.A = config.B
	config.B = config.C
	config.C = config.A + config.B
}

config.Lock.Unlock()

	c.JSON(http.StatusOK, resp)
}
