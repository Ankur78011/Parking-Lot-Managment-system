package models

type ParkingOwner struct {
	Id       int    `json:"id"`
	UserType string `json:"user_type"`
	Username string `json:"username"`
}
type NewParkingOwner struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
