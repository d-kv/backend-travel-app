package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/d-kv/backend-travel-app/pkg/adapter/igateway"
	"github.com/d-kv/backend-travel-app/pkg/domain/model/user"
	"github.com/d-kv/backend-travel-app/pkg/infra/irepository"
)

func Authorizer(uStore irepository.UserI, oAuthGateway igateway.OAuthProviderI) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		oAuthAToken, ok := ctx.GetQuery("access_token")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid access token",
			})
		}

		oAuthID, err := oAuthGateway.GetUserID(ctx, oAuthAToken)
		if err != nil {
			if errors.Is(err, igateway.ErrTokenIsExpired) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "access token is expired",
				})
			}

			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"error": "oAuth provider error",
			})
		}

		u, err := uStore.GetByOAuthID(ctx, oAuthID)
		if err != nil {
			if errors.Is(err, irepository.ErrUserNotFound) { // Create a new user
				u = user.New(
					user.WithOAuthAToken(oAuthAToken),
					user.WithOAuthID(oAuthID),
				)

				err = uStore.Create(ctx, u)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": "account creation error",
					})
				}
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "invalid account data",
			})
		}

		ctx.Set("user_id", u.UUID)

		ctx.Next()
	}
}
