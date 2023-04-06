package resthandlerv0

import (
	"github.com/gin-gonic/gin"

	icontrollerv0 "github.com/d-kv/backend-travel-app/pkg/app/icontroller/v0"
)

// HTTPHandler defines a HTTP handler.
type HTTPHandler struct {
	ctrl icontrollerv0.ControllerI
	eng  *gin.Engine
}

// New is a default HTTPHandler ctor.
func New(ctrl icontrollerv0.ControllerI, r *gin.Engine) *HTTPHandler {
	h := &HTTPHandler{
		ctrl: ctrl,
		eng:  r,
	}
	h.registerRoutes()
	return h
}

// TODO: sync API endpoints with the client
func (h *HTTPHandler) registerRoutes() {
	h.eng.POST("/api/v0/auth/oauth", h.PostOAuth)
	h.eng.POST("/api/v0/places/search", h.PostSearchPlaces)
}

func (h *HTTPHandler) Run(addr, port string) error {
	return h.eng.Run(addr + ":" + port)
}
