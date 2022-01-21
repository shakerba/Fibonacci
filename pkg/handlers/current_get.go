package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pex/fibonacci/pkg/config"
	"github.com/pex/fibonacci/pkg/models"
)

func (CurrentHandler *RouteHandler) GetCurrent(c *gin.Context) {
	resp := &models.CurrentResponse{
    Number: config.B,
		Err: nil,
	}
	c.JSON(http.StatusOK, resp)
}
