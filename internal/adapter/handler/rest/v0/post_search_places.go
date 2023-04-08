package resthandlerv0

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

func (h *HTTPHandler) PostSearchPlaces(ctx *gin.Context) {
	llStr, ok := ctx.GetQuery("ll")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing ll parameter",
		})
	}

	ll, err := util.NewLatLngFromString(llStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid ll parameter",
		})
	}

	places, err := h.ctrl.GetPlaces(ctx, ll)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err, // TODO: do not return raw errors
		})
	}
	if len(places) == 0 {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			// TODO: return StatusNotFound?
			// Maybe it makes sense to return StatusNotFound
			// but google returns StatusOK for empty search queries, for instance
			"message": "no places with such criteria",
		})
	}
	ctx.JSON(http.StatusOK, places)
}
