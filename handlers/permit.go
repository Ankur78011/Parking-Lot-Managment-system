package handlers

import (
	"ankur/parkinlot/models"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

// //GET PERMIT//
func (api *Apihandler) GetPermit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		owner_id := int(ctx.GetFloat64("user_id"))
		ctx.Header("Content-Type", "application/json")
		permits := api.Storage.GetPermits(owner_id)
		json.NewEncoder(ctx.Writer).Encode(permits)
	}
}

// /CREATE PERMIT///
func (api *Apihandler) CreatePermit() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Content-Type", "application/json")
		var Insertpermit models.NewPermit
		_ = json.NewDecoder(ctx.Request.Body).Decode(&Insertpermit)
		newId := api.Storage.CheckCustomer(Insertpermit.Customer_name)
		var uid int
		uid = newId
		///// IF CUSTOMER DOSEN'T EXISTS MAKE HE OR SHE USER AS A CUSTOMER AND INSERT IT INTO PERMIT//
		if uid == 0 {
			latestId := api.Storage.InsertNewCustomer(Insertpermit.Customer_name)
			uid = latestId
		}
		// INSERT DIRECTLY INTO PERMIT IF THE USER EXISTS//
		api.Storage.Insertpermit(Insertpermit.Start_date, Insertpermit.End_date, Insertpermit.Car_number, uid, Insertpermit.Parkinglot_id)
		json.NewEncoder(ctx.Writer).Encode("Inserted")

	}
}

// DELETE PERMIT//
func (api *Apihandler) DeletePermit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := int(ctx.GetFloat64("user_id"))
		ctx.Header("Content-type", "application/json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		api.Storage.DeletePermit(id, user_id)

	}
}

// //UPDATE PERMIT//
func (api *Apihandler) UpdatePermit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-type", "application/json")
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		var updatedinfo models.Permit
		_ = json.NewDecoder(ctx.Request.Body).Decode(&updatedinfo)
		api.Storage.UpdatePermit(updatedinfo.Start_date, updatedinfo.End_date, updatedinfo.Car_number, updatedinfo.Customer_id, updatedinfo.Parkinglot_id, id)
		json.NewEncoder(ctx.Writer).Encode("uppdate")
	}
}
