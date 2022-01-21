package handlers

const (
	HealthStatusOK = "OK"
)

type RouteHandler struct {
	ServiceName string
	Version     string
}
