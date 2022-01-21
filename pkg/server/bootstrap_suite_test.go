package server

import (
	"net/http"
	"net/http/httptest"

	"testing"

	. "github.com/onsi/ginkgo"
	. "github.ibm.com/genctl/test"

	"github.com/gin-gonic/gin"
	"github.com/onsi/gomega/gbytes"
	"github.com/pex/fibonacci/pkg/config"

)

var (
	testEngine *gin.Engine
	logBuffer  *gbytes.Buffer
)

func TestServer(t *testing.T) {
	TestPackage(t, "Server Suite")
}

func createGinEngine() *gin.Engine {
	// Create the in-memory logger
	logBuffer = gbytes.NewBuffer()

	// Create a config
	config :=  &config.ServerConfig{}

	gin.DefaultWriter = GinkgoWriter
	gin.DefaultErrorWriter = GinkgoWriter
	gin.SetMode(gin.ReleaseMode)

	// Set up API server
	return NewEngine(config, SetDefaultRoutes)
}

func performRequest(engine http.Handler, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	resp := httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	return resp
}
