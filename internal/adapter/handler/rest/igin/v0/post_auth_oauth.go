package iginv0

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	user_ctrl_v0 "github.com/d-kv/backend-travel-app/pkg/app/controller/v0/user"
)

type userOAuthAuthorizer interface {
	AuthorizeOAuthUser(ctx context.Context, accessToken, refreshToken string) (userUUID string, err error)
}

func (h *HTTPHandler) postAuthOAuth(ctx *gin.Context) {
	aT, ok := ctx.GetQuery("access_token")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing access_token parameter",
		})
		return
	}
	rT, ok := ctx.GetQuery("refresh_token")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing refresh_token parameter",
		})
		return
	}

	uID, err := h.userCtrl.AuthorizeOAuthUser(ctx, aT, rT)
	if err != nil {
		if errors.Is(err, user_ctrl_v0.ErrUserIsBlocked) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "user account is blocked",
			})
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "internal error",
			"description": err.Error(), // FIXME: do not return raw error
		})

		return
	}

	ctx.JSON(http.StatusOK, uID)
}
