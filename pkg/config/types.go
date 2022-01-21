package config

const (
	ServiceName   = "fibonacci-server"
)

// ServerConfig
type ServerConfig struct {
	ServiceName  string
	Version      string
	SwaggerPath  string
	BasePath     string
}
