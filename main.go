package main

import (
	"ankur/parkinlot/handlers"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "pklot"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("CONNNECTED!!!")
	router := gin.Default()
	apihandler := handlers.NewApiHandler(db)
	/////Login/////
	router.POST("/login", apihandler.Login())
	////parkinglot owner//
	router.GET("/parkingowner", apihandler.GetParkingOwner())
	router.POST("/parkingowner", apihandler.CreateParkingOwner())
	router.DELETE("/parkingowner/:id", apihandler.DeleteParkingOwner())
	/////customer/////
	router.GET("/customer", apihandler.GetCustomer())
	router.POST("/customer", apihandler.CreateCustomer())
	router.DELETE("/customer/:id", apihandler.DeleteCustomer())
	// ////////Parkinglot////////
	router.GET("/parkinglot", apihandler.Token(), apihandler.GetParkingLot())
	router.POST("/parkinglot", apihandler.Token(), apihandler.CreateParkingLot())
	router.DELETE("/parkinglot/:id", apihandler.Token(), apihandler.DeleteParkingLot())
	router.PUT("/parkinglot/:id", apihandler.Token(), apihandler.UpdateParkingLot())
	//////////Permit/////
	router.GET("/permit", apihandler.Token(), apihandler.GetPermit())
	router.POST("/permit", apihandler.Token(), apihandler.CreatePermit())
	router.DELETE("/permit/:id", apihandler.Token(), apihandler.DeletePermit())
	router.PUT("/permit/:id", apihandler.Token(), apihandler.UpdatePermit())
	// //////fine./////////
	router.GET("/fine", apihandler.Token(), apihandler.Getfine())
	router.POST("/fine", apihandler.Token(), apihandler.CreateFine())
	router.PUT("/fine/:id", apihandler.Token(), apihandler.UpdateFine())
	router.DELETE("/fine/:id", apihandler.Token(), apihandler.DeleteFine())
	// ///
	router.Run(":8000")
}
