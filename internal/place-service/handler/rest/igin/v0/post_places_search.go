package igin_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
)

const defaultResN = 50

type placeSearcher interface {
	SearchPlaces(ctx context.Context,
		geoQ *util.GeoQuery,
		mCats []category.Main, sCats []category.Sub,
		skipN int64, resN int64) ([]model.Place, error)
}

type reqBody struct {
	Category category.Category `json:"category"`
	SeenUUID []string          `json:"seen_uuid"`
}

func (h *PlaceHandler) postPlacesSearch(ctx *gin.Context) {
	llStr, ok := ctx.GetQuery("ll")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing ll parameter",
		})
		return
	}
	ll, err := util.NewLatLngFromString(llStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid ll parameter",
		})
		return
	}

	minDStr, ok := ctx.GetQuery("min_d")
	minD := int64(util.DefaultMinDistance)
	if ok {
		minD, err = strconv.ParseInt(strings.TrimSpace(minDStr), base, bitSize)
		if err != nil || minD < 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid min_d parameter",
			})
			return
		}
	}

	maxDStr, ok := ctx.GetQuery("max_d")
	maxD := int64(util.DefaultMaxDistance)
	if ok {
		maxD, err = strconv.ParseInt(strings.TrimSpace(maxDStr), base, bitSize)
		if err != nil || maxD <= 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid max_d parameter",
			})
			return
		}
	}

	if maxD < minD {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid max_d & min_d parameters",
		})
		return
	}

	resNStr, ok := ctx.GetQuery("result_n")
	resN := int64(defaultResN)
	if ok {
		resN, err = strconv.ParseInt(strings.TrimSpace(resNStr), base, bitSize)
		if err != nil || resN <= 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid result_n parameter",
			})
			return
		}
	}

	skipNStr, ok := ctx.GetQuery("skip_n")
	skipN := int64(0)
	if ok {
		skipN, err = strconv.ParseInt(strings.TrimSpace(skipNStr), base, bitSize)
		if err != nil || skipN < 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid skip_n parameter",
			})
			return
		}
	}

	gQ := util.NewGeoQuery(ll,
		util.WithMin(minD),
		util.WithMax(maxD),
	)

	var mCtgs []category.Main
	var sCtgs []category.Sub
	if ctx.Request.ContentLength != 0 {
		var req reqBody
		err = ctx.BindJSON(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}

		mCtgs = req.Category.Main
		sCtgs = req.Category.Sub
	}

	places, err := h.placeCtrl.SearchPlaces(ctx, gQ, mCtgs, sCtgs, skipN, resN)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "internal error",
			"description": err.Error(), // FIXME: do not return raw error
		})

		return
	}

	ctx.JSON(http.StatusOK, places)
}
