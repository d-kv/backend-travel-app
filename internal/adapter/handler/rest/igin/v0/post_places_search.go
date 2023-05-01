package iginv0

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/place"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/place/category"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/query"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

const defaultResN = 50

type placeSearcher interface {
	SearchPlaces(ctx context.Context,
		geoQ *query.Geo,
		mCats []category.MainCategory, sCats []category.SubCategory,
		skipN int64, resN int64) ([]place.Place, error)
}

type reqBody struct {
	Category category.Category `json:"category"`
	SeenUUID []string          `json:"seen_uuid"`
}

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type resPlace struct {
	UUID        string `json:"uuid"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Phone       string `json:"phone,omitempty"`
	URL         string `json:"url,omitempty"`

	Category category.Category `json:"category"`
	Location *location         `json:"location"`

	Lifetime  time.Duration `json:"lifetime"`
	CreatedAt time.Time     `json:"created_at"`
}

func toResPlace(p *place.Place) *resPlace {
	return &resPlace{
		UUID:        p.UUID,
		Address:     p.Address,
		Name:        p.Name,
		Description: p.Description,
		Phone:       p.Phone,
		URL:         p.URL,
		Category:    p.Category,
		Location: &location{
			Latitude:  p.Latitude(),
			Longitude: p.Longitude(),
		},
		Lifetime:  p.Lifetime,
		CreatedAt: p.CreatedAt,
	}
}

func toResBody(places []place.Place) []resPlace {
	ret := make([]resPlace, 0, len(places))

	for i := range places {
		ret = append(ret, *toResPlace(&places[i]))
	}

	return ret
}

func (h *HTTPHandler) postPlacesSearch(ctx *gin.Context) {
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
	minD := int64(query.DefaultMinDistance)
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
	maxD := int64(query.DefaultMaxDistance)
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

	gQ := query.New(ll,
		query.WithMin(minD),
		query.WithMax(maxD),
	)

	var mCtgs []category.MainCategory
	var sCtgs []category.SubCategory
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
			"description": err, // FIXME: do not return raw error
		})

		return
	}

	ctx.JSON(http.StatusOK, toResBody(places))
}
