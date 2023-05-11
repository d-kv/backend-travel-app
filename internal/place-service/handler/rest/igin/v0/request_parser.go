package igin_v0 //nolint:revive,stylecheck // using underscore in package name for better readability

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/util"
)

const (
	qLat   = "lat"
	qLng   = "lng"
	qMinD  = "min_d"
	qMaxD  = "max_d"
	qSkipN = "skip_n"
	qResN  = "result_n"
)

func parseCategories(ctx *gin.Context) (*category.Category, error) {
	if ctx.Request.ContentLength == 0 {
		return category.NewAnyCategory(), nil
	}

	var req requestBody
	err := ctx.BindJSON(&req)
	if err != nil {
		return nil, errInvalidBody
	}

	ctgs := &category.Category{
		Main: req.Category.Main,
		Sub:  req.Category.Sub,
	}

	return ctgs, nil
}

func parseLatLng(ctx *gin.Context) (*util.LatLng, error) {
	latStr, hasLat := ctx.GetQuery(qLat)
	lngStr, hasLng := ctx.GetQuery(qLng)

	hasBoth := hasLat && hasLng

	if !hasBoth {
		return nil, errLatLngCoupling
	}

	lat, err := strconv.ParseFloat(strings.TrimSpace(latStr), bitSize)
	if err != nil {
		return nil, err
	}

	lng, err := strconv.ParseFloat(strings.TrimSpace(lngStr), bitSize)
	if err != nil {
		return nil, err
	}

	ll, err := util.NewLatLng(lat, lng)
	if err != nil {
		return nil, errInvalidLatLng
	}

	return ll, nil
}

func parseGeoToken(ctx *gin.Context) (*util.GeoToken, error) {
	ll, err := parseLatLng(ctx)
	if err != nil {
		return nil, err
	}

	minD := int64(util.DefaultMinDistance)
	maxD := int64(util.DefaultMaxDistance)

	minDStr, hasMinD := ctx.GetQuery(qMinD)
	maxDStr, hasMaxD := ctx.GetQuery(qMaxD)

	hasBoth := hasMaxD && hasMinD
	hasOnlyOne := hasMaxD != hasMinD // equals hasMaxD ^ hasMinD
	// min_d & min_d must either be both present or both absent

	// hasMinD=True,	hasMaxD=True 		hasMinD^hasMaxD=False
	// hasMinD=True, 	hasMaxD=False 	hasMinD^hasMaxD=True
	// hasMinD=False, hasMaxD=True 		hasMinD^hasMaxD=True
	// hasMinD=False, hasMaxD=False 	hasMinD^hasMaxD=False

	if hasOnlyOne {
		return nil, errMinDMaxDCoupling
	}

	if hasBoth {
		var err error

		minD, err = strconv.ParseInt(strings.TrimSpace(minDStr), base, bitSize)
		if err != nil || minD < 0 {
			return nil, errInvalidMinD
		}

		maxD, err = strconv.ParseInt(strings.TrimSpace(maxDStr), base, bitSize)
		if err != nil || maxD <= 0 {
			return nil, errInvalidMaxD
		}

		if maxD < minD {
			return nil, errMaxDSmallerThanMinD
		}
	}

	geoT := util.NewGeoToken(ll,
		util.WithMin(minD),
		util.WithMax(maxD),
	)

	return geoT, nil
}

func parsePaginationToken(ctx *gin.Context) (*util.PaginationToken, error) {
	skipNStr, hasSkipN := ctx.GetQuery(qSkipN)
	resNStr, hasResN := ctx.GetQuery(qResN)

	resN := int64(util.DefaultResN)
	skipN := int64(util.DefaultSkipN)

	var err error
	if hasResN {
		resN, err = strconv.ParseInt(strings.TrimSpace(resNStr), base, bitSize)
		if err != nil || resN <= 0 {
			return nil, errInvalidSkipN
		}
	}

	if hasSkipN {
		skipN, err = strconv.ParseInt(strings.TrimSpace(skipNStr), base, bitSize)
		if err != nil || skipN < 0 {
			return nil, errInvalidResN
		}
	}

	pT := util.NewPaginationToken(
		skipN, resN,
	)

	return pT, nil
}
