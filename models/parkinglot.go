package models

type Parkinglot struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	OwnerId int    `json:"ownerid"`
}
