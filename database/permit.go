package database

import (
	"ankur/parkinlot/models"
	"fmt"
)

func (m *Storage) GetPermits(owner_id int) []models.Permit {
	sqlQuery := `SELECT *FROM permit WHERE parkinglot_id IN (SELECT id FROM parkinglot WHERE owner_id=$1)`
	rows, err := m.Db.Query(sqlQuery, owner_id)
	if err != nil {
		panic(err)
	}
	var permits []models.Permit
	for rows.Next() {
		var per models.Permit
		err := rows.Scan(&per.Id, &per.Start_date, &per.End_date, &per.Car_number, &per.Customer_id, &per.Parkinglot_id)
		if err != nil {
			panic(err)
		}
		permits = append(permits, per)
	}
	return permits
}
func (m *Storage) Insertpermit(start_Date string, end_Date string, car_number string, uid int, parkinglot_id int) {
	sqlStatment := `INSERT INTO permit(start_date,end_date,car_number,customer_id,parkinglot_id) values ($1,$2,$3,$4,$5)`
	_, err := m.Db.Exec(sqlStatment, start_Date, end_Date, car_number, uid, parkinglot_id)
	if err != nil {
		fmt.Println("LINE36")
		panic(err)
	}
}
func (m *Storage) DeletePermit(id int, user_id int) {
	sqlstatment := `DELETE FROM Permit where parkinglot_id IN (SELECT id from parkinglot WHERE owner_id=$1) AND id=$2`
	_, err := m.Db.Query(sqlstatment, user_id, id)
	if err != nil {
		panic(err)
	}
}
func (m *Storage) UpdatePermit(start_date string, end_date string, car_number string, customer_id int, parkinglot_id int, id int) {
	sqlstatmentL := `UPDATE permit set start_date=$1,end_date=$2,car_number=$3,customer_id=$4,parkinglot_id=$5 where id=$6`
	_, err := m.Db.Exec(sqlstatmentL, start_date, end_date, car_number, customer_id, parkinglot_id, id)
	if err != nil {
		panic(err)
	}
}
func (m *Storage) DeleteOwnerPermit(id int) {
	sql := `DELETE FROM permit WHERE parkinglot_id IN (SELECT id from parkinglot WHERE owner_id=$1)`
	_, err := m.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
}
