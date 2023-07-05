package models

type Permit struct {
	Id            int    `json:"id"`
	Start_date    string `json:"start_date"`
	End_date      string `json:"end_date"`
	Car_number    string `json:"car_number"`
	Customer_id   int    `json:"customer_id"`
	Parkinglot_id int    `json:"parking_lot"`
}

type NewPermit struct {
	Start_date    string `json:"start_date"`
	End_date      string `json:"end_date"`
	Car_number    string `json:"car_number"`
	Customer_name string `json:"customer_name"`
	Parkinglot_id int    `json:"parking_lot"`
}
