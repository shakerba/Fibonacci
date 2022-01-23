package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pex/fibonacci/pkg/config"
	"github.com/pex/fibonacci/pkg/models"
)

func (PreviousHandler *RouteHandler) GetPrevious(c *gin.Context) {
	var resp *models.PreviousResponse

	resp = &models.PreviousResponse{
		Number: config.A,
		Err: nil,
	}

config.Lock.Lock()

if config.A != 0 {
	config.C = config.B
	config.B = config.A
	config.A = config.C - config.B
} else if config.B == 1 && config.C == 1 {
	config.A = 0
	config.B = 0
	config.C = 0
	resp.Number = config.A
}
config.Lock.Unlock()

	c.JSON(http.StatusOK, resp)
}
