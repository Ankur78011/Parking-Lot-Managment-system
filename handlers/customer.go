package handlers

import (
	"ankur/parkinlot/models"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

// //GET CUSTOMER///
func (api *Apihandler) GetCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Content-type", "application/json")
		customerDetail := api.Storage.GetCustomer()
		json.NewEncoder(ctx.Writer).Encode(customerDetail)

	}
}

// /CREATE CUSTOMER////
func (api *Apihandler) CreateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		var Newuser models.ParkingOwner
		_ = json.NewDecoder(ctx.Request.Body).Decode(&Newuser)
		result := api.Storage.PostCustomer(Newuser.Username)
		json.NewEncoder(ctx.Writer).Encode(result)

	}
}

// /DELETE CUSTOMER///
func (api *Apihandler) DeleteCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-type", "application/json")
		id, _ := strconv.Atoi(ctx.Param("id"))
		api.Storage.Deletecustomer(id)
		json.NewEncoder(ctx.Writer).Encode("sucessfully record deleted")
	}
}

// //UPDATE CUSTOMER/////
func (api *Apihandler) UpdateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		var Updateduser models.ParkingOwner
		_ = json.NewDecoder(ctx.Request.Body).Decode(&Updateduser)
		api.Storage.UpdateCustomer(Updateduser.Username, id)
		json.NewEncoder(ctx.Writer).Encode("Updated owner")
	}
}
