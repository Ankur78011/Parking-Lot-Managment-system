package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (api *Apihandler) Token() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		tkn := ctx.GetHeader("token")
		token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing  method")
			}
			return []byte("ankur"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		expireTime := int(claims["exp"].(float64))
		if expireTime < int(time.Now().Unix()) {
			ctx.AbortWithStatus(http.StatusGatewayTimeout)
			return
		}
		if ok && token.Valid {

			value, present := claims["user_id"]
			if !present {
				fmt.Println("no value nil")
			}
			convertedValue := value.(float64)

			fmt.Println(convertedValue)
			result := api.Storage.ValidateUser(convertedValue)
			if !result {
				fmt.Println("user is not authenticated")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			fmt.Println("user is authenticated")
			ctx.Set("user_id", claims["user_id"].(float64))
			ctx.Next()

		}

	}
}
