package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/pkg/domain/model/util"
)

func Versioner(reqV util.Version) gin.HandlerFunc {
	return func(c *gin.Context) {
		appVStr, ok := c.GetQuery("ver")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing version",
			})
		}

		appV, err := util.NewVersionFromString(appVStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid version format",
			})
		}

		if appV.Less(reqV) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid version",
			})
		}

		c.Next()
	}
}
