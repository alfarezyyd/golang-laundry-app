package middleware

import (
	"net/http"
)

type GatewayMiddleware struct {
	Handler       http.Handler
	allMiddleware []func(next http.Handler) http.Handler
}

func NewGatewayMiddleware(handler http.Handler) *GatewayMiddleware {
	return &GatewayMiddleware{Handler: handler}
}

func (gatewayMiddleware *GatewayMiddleware) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	gatewayMiddleware.allMiddleware = append(gatewayMiddleware.allMiddleware, next)
}

func (gatewayMiddleware *GatewayMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	currentRoute := gatewayMiddleware.Handler
	for _, nextMiddleware := range gatewayMiddleware.allMiddleware {
		currentRoute = nextMiddleware(currentRoute)
	}

	currentRoute.ServeHTTP(w, r)
}
