package handlers

import (
	"ankur/parkinlot/models"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

// //GET FINE//
func (api *Apihandler) Getfine() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-type", "application/json")
		user_id := int(ctx.GetFloat64("user_id"))
		Fines := api.Storage.GetFine(user_id)
		json.NewEncoder(ctx.Writer).Encode(Fines)
	}
}

// CREATE FINE//
func (api *Apihandler) CreateFine() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		var Newfine models.Fine
		_ = json.NewDecoder(ctx.Request.Body).Decode(&Newfine)
		res := api.Storage.Postfine(Newfine.ParkingLot_id, Newfine.Car_number, Newfine.Car_Owner, Newfine.Amount)
		json.NewEncoder(ctx.Writer).Encode(res)
	}
}

// DELETE FINE//
func (api *Apihandler) DeleteFine() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.Deletefine(id)
		json.NewEncoder(ctx.Writer).Encode("Succesfully deleted")

	}
}

// UPDATEFINE//
func (api *Apihandler) UpdateFine() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-type", "application/json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		var updatevalue models.Fine
		_ = json.NewDecoder(ctx.Request.Body).Decode(&updatevalue)
		api.Storage.Updatefine(updatevalue.ParkingLot_id, updatevalue.Car_number, updatevalue.Car_Owner, updatevalue.Amount, id)
		json.NewEncoder(ctx.Writer).Encode("Updated successfully")
	}
}
