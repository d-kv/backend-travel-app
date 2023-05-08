package igin_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
)

type placeSearcher interface {
	SearchPlaces(ctx context.Context,
		geoQ *util.GeoToken,
		mCats []category.Main, sCats []category.Sub,
		skipN int64, resN int64) ([]model.Place, error)
}

type requestBody struct {
	Category category.Category `json:"category"`
	SeenUUID []string          `json:"seen_uuid"`
}

func (h *PlaceHandler) postPlacesSearch(ctx *gin.Context) {
	geoT, err := parseGeoToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	paginT, err := parsePaginationToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctgs, err := parseCategories(ctx)
	if err != nil && !errors.Is(err, errEmptyBody) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	places, err := h.placeCtrl.SearchPlaces(ctx, geoT, ctgs.Main, ctgs.Sub, paginT.SkipN, paginT.ResN)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "internal error",
			"description": err.Error(), // FIXME: do not return raw error
		})
		return
	}

	ctx.JSON(http.StatusOK, places)
}
