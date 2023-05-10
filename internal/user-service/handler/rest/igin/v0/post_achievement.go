package igin_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/internal/user-service/handler/rest/igin"
	ctrl_v0 "github.com/d-kv/backend-travel-app/pkg/user-service/controller/v0"
	"github.com/d-kv/backend-travel-app/pkg/user-service/model"
)

type userAchievementsUpdater interface {
	AddAchievement(ctx context.Context, achievement *model.Achievement, userUUID string) error
}

func (h *UserHandler) postAchievement(ctx *gin.Context) {
	uID, err := igin.UserUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "internal error",
			"description": err.Error(), // FIXME: do not return raw error
		})
		return
	}

	var ach model.Achievement
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
		if errors.Is(err, ctrl_v0.ErrAchievementAlreadyExists) {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "user already has this achievement",
			})
		} else if errors.Is(err, ctrl_v0.ErrBadAchievement) {
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
