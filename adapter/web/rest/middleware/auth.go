package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	middlewareIntf "kickoff/adapter/web/rest/middleware/interface"
	"kickoff/dto"
	facade "kickoff/internal/facade/interface"
	"net/http"
)

type AuthMiddleware struct {
	securityFacade facade.SecurityFacade
}

func NewAuthMiddleware(securityFacade facade.SecurityFacade) middlewareIntf.AuthMiddleware {
	return &AuthMiddleware{
		securityFacade: securityFacade,
	}
}

func (a *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := context.Background()

		const bearerSchema = "Bearer "

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := header[len(bearerSchema):]

		if header != (bearerSchema + token) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenDTO := &dto.ValidateTokenRequest{Value: token}

		response, err := a.securityFacade.ValidateToken(ctx, tokenDTO)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				dto.Error{
					Message: err.Error(),
				},
			)
			return
		}

		if response == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user_id", response.UserID)
	}
}
