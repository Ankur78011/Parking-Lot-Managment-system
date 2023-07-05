package models

type Fine struct {
	Id            int    `json:"id"`
	ParkingLot_id int    `json:"parking_lot_id"`
	Car_number    string `json:"car_number"`
	Car_Owner     string `json:"car_owner"`
	Amount        int    `json:"amount"`
}
