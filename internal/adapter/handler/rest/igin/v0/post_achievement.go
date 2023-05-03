package iginv0

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	igin "github.com/d-kv/backend-travel-app/internal/adapter/handler/rest/igin"
	user_ctrl_v0 "github.com/d-kv/backend-travel-app/pkg/app/controller/v0/user"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
)

type userAchievementsUpdater interface {
	AddAchievement(ctx context.Context, achievement *user.Achievement, userUUID string) error
}

func (h *HTTPHandler) postAchievement(ctx *gin.Context) {
	uID, err := igin.UserUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "internal error",
			"description": err.Error(), // FIXME: do not return raw error
		})
		return
	}

	var ach user.Achievement
	err = ctx.BindJSON(&ach)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "missing achievement object",
		})
		return
	}

	err = h.userCtrl.AddAchievement(ctx, &ach, uID)
	if err != nil {
		//nolint:gocritic // switch on an error will fail on wrapped errors
		if errors.Is(err, user_ctrl_v0.ErrAchievementAlreadyExists) {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "user already has this achievement",
			})
		} else if errors.Is(err, user_ctrl_v0.ErrBadAchievement) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "undefined achievement",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":       "internal error",
				"description": err.Error(), // FIXME: do not return raw error
			})
		}

		return
	}

	ctx.Status(http.StatusAccepted)
}
