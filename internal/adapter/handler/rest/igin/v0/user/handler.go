package igin_user_v0

import (
	"github.com/gin-gonic/gin"
)

type (
	userCtrl interface {
		userOAuthAuthorizer

		userAchievementsUpdater
		userAchievementsProvider
	}
)

// UserHandler defines a HTTP handler.
type UserHandler struct {
	userCtrl
	eng *gin.Engine
}

// New is a default UserHandler ctor.
func New(uCtrl userCtrl, eng *gin.Engine) *UserHandler {
	h := &UserHandler{
		userCtrl: uCtrl,
		eng:      eng,
	}
	h.registerRoutes()
	return h
}

func (h *UserHandler) registerRoutes() {
	h.eng.POST("/api/v0/auth/oauth", h.postAuthOAuth)

	h.eng.GET("/api/v0/account/achievements", h.getAchievements)
	h.eng.POST("/api/v0/account/achievement", h.postAchievement)
}

func (h *UserHandler) Run(addr, port string) error {
	return h.eng.Run(addr + ":" + port)
}
