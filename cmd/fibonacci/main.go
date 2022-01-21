// @APIVersion 1.0
// @APITitle Fibonnaci API
// @APIDescription Public API for Fibonacci Service

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pex/fibonacci/pkg/config"
	"github.com/pex/fibonacci/pkg/server"
)

var (
	httpAddr     = flag.String("httpAddr", ":8080", "HTTP listen address")
	swaggerPath  = flag.String("swaggerPath", "/etc/swagger", "Path to swagger asset root")
	version      = flag.String("version", "dev", "Version of the fibonacci")
	basePath     = flag.String("basePath", "", "base path of this rest server")
)

var (
	httpServer       *http.Server
)

func main() {

	flag.Parse()

	ctx := context.Background()

	gin.SetMode(gin.ReleaseMode)
	// Create the server configuration
	serverConfig := &config.ServerConfig{
			ServiceName:  "fibonacci",
			Version:      *version,
			SwaggerPath:  *swaggerPath,
			BasePath:     *basePath,
	}

for { //Restart server if you receive these process codes
	// Start  HTTP server
	go startHTTPServer(context.Background(), serverConfig, *httpAddr)

	//Wait till SIGTERM comes

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // block until signal received
	fmt.Println("SIGTERM received, stopping servers")
	stopServer(ctx)
}

}

func stopServer(ctx context.Context) {
	var wg sync.WaitGroup

	//HTTP
	if httpServer != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := httpServer.Shutdown(ctx)

			if err != nil {
				fmt.Printf("Failed to shutdown server %s\n", err.Error())
			}

		}()
	}

	// Wait
	wg.Wait()
}

func startHTTPServer(ctx context.Context, serverConfig *config.ServerConfig, addr string) {
	fmt.Printf("http(non-secure) fibonacci server starting at %v", *httpAddr)
	ginEngine := server.NewEngine(serverConfig, server.SetDefaultRoutes)
	httpServer = &http.Server{
		Addr:    addr,
		Handler: ginEngine,
	}

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Failed to start http server %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("http fibonacci server stopping")
}
