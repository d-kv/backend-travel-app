package igin_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"github.com/gin-gonic/gin"
)

type (
	placeCtrl interface {
		placeSearcher
	}
)

// PlaceHandler defines a HTTP handler.
type PlaceHandler struct {
	placeCtrl
	eng *gin.Engine
}

// New is a default PlaceHandler ctor.
func New(pCtrl placeCtrl, eng *gin.Engine) *PlaceHandler {
	h := &PlaceHandler{
		placeCtrl: pCtrl,
		eng:       eng,
	}
	h.registerRoutes()
	return h
}

func (h *PlaceHandler) registerRoutes() {
	h.eng.POST("/api/v0/places/search", h.postPlacesSearch)
}

func (h *PlaceHandler) Run(addr, port string) error {
	return h.eng.Run(addr + ":" + port)
}
