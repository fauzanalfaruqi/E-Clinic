package middleware

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/dto/json"
	"avengers-clinic/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			json.NewResponseUnauthorized(c, "Invalid token", "02", "01")
			c.Abort()
			return
		}

		tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
		token, err := utils.VerifyJWT(tokenString)
		if err != nil {
			json.NewResponseError(c, err.Error(), "02", "03")
			c.Abort()
			return
		}

		if !token.Valid {
			json.NewResponseForbidden(c, "Forbidden", "02", "04")
			c.Abort()
			return
		}
		claims := token.Claims.(*dto.JWTClams)

		validRole := false
		if len(roles) > 0 {
			for _, role := range roles {
				if claims.Role == role {
					validRole = true
					break
				}
			}
		}

		if !validRole {
			json.NewResponseForbidden(c, "Forbidden", "02", "05")
			c.Abort()
			return
		}
		c.Next()
	}
}