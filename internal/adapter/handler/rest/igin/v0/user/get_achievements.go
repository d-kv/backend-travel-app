package igin_user_v0

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	igin "github.com/d-kv/backend-travel-app/internal/adapter/handler/rest/igin"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

type userAchievementsProvider interface {
	GetAchievements(ctx context.Context, userUUID string) ([]user.Achievement, error)
}

func (h *UserHandler) getAchievements(ctx *gin.Context) {
	uID, err := igin.UserUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "internal error",
			"description": err.Error(), // FIXME: do not return raw error
		})
		return
	}

	achs, err := h.userCtrl.GetAchievements(ctx, uID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "internal error",
			"description": err.Error(), // FIXME: do not return raw error
		})

		return
	}

	ctx.JSON(http.StatusOK, achs)
}
