package handlers

import (
	"ankur/parkinlot/models"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET PARKINGLOT//
func (api *Apihandler) GetParkingLot() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := int(ctx.GetFloat64("user_id"))
		fmt.Println(userId, "dkndkndkn")
		ctx.Header("Content-Type", "application/json")
		ParkinglotList := api.Storage.Getparkinglot(userId)
		json.NewEncoder(ctx.Writer).Encode(ParkinglotList)

	}
}

// CREATE PARKINGLOT//
func (api *Apihandler) CreateParkingLot() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		var Newparkinglot models.Parkinglot
		_ = json.NewDecoder(ctx.Request.Body).Decode(&Newparkinglot)
		api.Storage.CheckUserType(Newparkinglot.OwnerId)
		result := api.Storage.InserParkinglot(Newparkinglot.Name, Newparkinglot.Address, Newparkinglot.OwnerId)
		json.NewEncoder(ctx.Writer).Encode(result)

	}
}

// DELETE PARKINGLOT
func (api *Apihandler) DeleteParkingLot() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := int(ctx.GetFloat64("user_id"))
		ctx.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.DeleteParkinglot(id, user_id)

	}
}

// UPDATE PARKING LOT//
func (api *Apihandler) UpdateParkingLot() gin.HandlerFunc {
	return (func(ctx *gin.Context) {
		user_id := int(ctx.GetFloat64("user_id"))
		ctx.Header("Content_type", "application/json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		var Updatedvalue models.Parkinglot

		_ = json.NewDecoder(ctx.Request.Body).Decode(&Updatedvalue)

		api.Storage.CheckUserType(Updatedvalue.OwnerId)

		api.Storage.UpdateParkinglot(Updatedvalue.Name, Updatedvalue.Address, Updatedvalue.OwnerId, id, user_id)

	})
}
