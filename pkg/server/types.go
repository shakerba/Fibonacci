package server

import (
	"github.com/gin-gonic/gin"
)

const (
	Prefix = "Phantom-"
)

// SetHandlerRoutesFunc interface to define a generic method to set routes
type SetHandlerRoutesFunc func(
	r *gin.Engine, serviceName, version, basePath string)
