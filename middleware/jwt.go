package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Hulhay/goldfish/shared"
	"github.com/Hulhay/goldfish/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(u usecase.Token) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if len(authHeader) <= 0 {
			response := shared.BuildErrorResponse("Failed", "authorization is empty")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) < 2 {
			response := shared.BuildErrorResponse("Failed", "authorization is empty")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if splitToken[0] != "Bearer" {
			response := shared.BuildErrorResponse("Failed", "authorization is invalid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := u.ValidateToken(splitToken[1])
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("email", claims["email"])
		} else {
			log.Println(err)
			response := shared.BuildErrorResponse("Token is not valid", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
