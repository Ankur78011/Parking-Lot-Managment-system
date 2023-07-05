package handlers

import (
	"ankur/parkinlot/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Apihandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-type", "application/json")
		var LoginInfo models.Login
		_ = json.NewDecoder(ctx.Request.Body).Decode(&LoginInfo)
		mess, result := api.Storage.Login(LoginInfo.Email, LoginInfo.Password)
		if result {
			user_id := api.Storage.GetUserId(LoginInfo.Email)
			fmt.Println(user_id)
			token, _ := generateJwt(user_id)
			ctx.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": mess})

	}
}
