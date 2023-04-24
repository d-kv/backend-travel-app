package iginv0

import (
	"github.com/gin-gonic/gin"
)

type (
	userCtrl interface {
		userOAuthAuthorizer

		userAchievementsUpdater
		userAchievementsProvider
	}

	placeCtrl interface {
		placeSearcher
	}
)

// HTTPHandler defines a HTTP handler.
type HTTPHandler struct {
	userCtrl
	placeCtrl
	eng *gin.Engine
}

// New is a default HTTPHandler ctor.
func New(uCtrl userCtrl, pCtrl placeCtrl, eng *gin.Engine) *HTTPHandler {
	h := &HTTPHandler{
		userCtrl:  uCtrl,
		placeCtrl: pCtrl,
		eng:       eng,
	}
	h.registerRoutes()
	return h
}

func (h *HTTPHandler) registerRoutes() {
	h.eng.POST("/api/v0/auth/oauth", h.postAuthOAuth)

	h.eng.POST("/api/v0/places/search", h.postPlacesSearch)

	h.eng.GET("/api/v0/account/achievements", h.getAchievements)
	h.eng.POST("/api/v0/account/achievement", h.postAchievement)
}

func (h *HTTPHandler) Run(addr, port string) error {
	return h.eng.Run(addr + ":" + port)
}
