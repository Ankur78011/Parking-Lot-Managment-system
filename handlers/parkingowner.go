package handlers

import (
	"ankur/parkinlot/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// /GET PARKING OWNER //////
func (api *Apihandler) GetParkingOwner() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ctx.Header("Content-type", "application/json")
		OwnersList := api.Storage.GetOwners()
		json.NewEncoder(ctx.Writer).Encode(OwnersList)
	}
}

// ///CREATE PARKING OWNER START///
func (api *Apihandler) CreateParkingOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		var Newuser models.NewParkingOwner
		err := ctx.ShouldBindJSON(&Newuser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Credentials missing"})
			return
		}

		_ = api.Storage.PostOwner(Newuser.Username, Newuser.Email, Newuser.Password)
		user_id := api.Storage.GetUserId(Newuser.Email)
		fmt.Println(user_id)
		token, _ := generateJwt(user_id)
		ctx.JSON(http.StatusOK, gin.H{"token": token})

	}
}

// /jwttoken///
func generateJwt(user_id int) (string, error) {
	var key = []byte("ankur")
	expirationtime := time.Now().Add(time.Minute * 60).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = expirationtime
	tokenString, _ := token.SignedString(key)
	return tokenString, nil
}

// ///DELETE PARKING OWNER START///
func (api *Apihandler) DeleteParkingOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-type", "application/json")
		id, _ := strconv.Atoi(ctx.Param("id"))
		api.Storage.DeleteOwnerFine(id)
		api.Storage.DeleteOwnerPermit(id)
		api.Storage.DeleteOwnerParkingLot(id)
		api.Storage.DeleteOwner(id)
		json.NewEncoder(ctx.Writer).Encode("sucessfully record deleted")
	}
}

// /UPDATE PAKRINGLOT OWNER////
func (api *Apihandler) Updateparkingowner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		var Updateduser models.ParkingOwner
		_ = json.NewDecoder(ctx.Request.Body).Decode(&Updateduser)
		api.Storage.UpdateParkingOwner(Updateduser.Username, id)
		json.NewEncoder(ctx.Writer).Encode("Updated owner")
	}
}
