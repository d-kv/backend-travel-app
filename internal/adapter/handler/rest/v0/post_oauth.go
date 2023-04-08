package resthandlerv0

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

func (h *HTTPHandler) PostOAuth(ctx *gin.Context) {
	uID, ok := ctx.Get("user_id")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
	}
	u, err := h.ctrl.GetUser(ctx, uID.(string))
	if err != nil {
		if errors.Is(err, irepository.ErrUserNotFound) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "account not found",
			})
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err, // TODO: do not return raw errors
		})
	}

	ctx.JSON(http.StatusOK, u)
}
