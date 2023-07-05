package database

import (
	"ankur/parkinlot/models"
)

func (m *Storage) GetFine(user_id int) []models.Fine {
	var Fines []models.Fine
	sqlQuery := `SELECT * FROM fine WHERE parkinglot_id IN (SELECT id FROM parkinglot WHERE owner_id=$1)`
	rows, err := m.Db.Query(sqlQuery, user_id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var fineinfo models.Fine
		err := rows.Scan(&fineinfo.Id, &fineinfo.ParkingLot_id, &fineinfo.Car_number, &fineinfo.Car_Owner, &fineinfo.Amount)
		if err != nil {
			panic(err)
		}
		Fines = append(Fines, fineinfo)
	}
	return Fines
}

func (m *Storage) Postfine(parkinglot_id int, car_number string, car_owner string, amount int) int64 {
	sqlStatment := `INSERT INTO  fine (parkinglot_id,car_number,car_owner,amount) values($1,$2,$3,$4)`
	res, err := m.Db.Exec(sqlStatment, parkinglot_id, car_number, car_owner, amount)
	if err != nil {
		panic(err)
	}
	newId, err := res.LastInsertId()
	return newId
}

func (m *Storage) Updatefine(parkinglot_id int, car_number string, car_owner string, amount int, id int) {
	sqlStatment := `UPDATE fine SET parkinglot_id=$1,car_number=$2,car_owner=$3,amount=$4 where id=$5`
	_, err := m.Db.Exec(sqlStatment, parkinglot_id, car_number, car_owner, amount, id)
	if err != nil {
		panic(err)
	}
}

func (m *Storage) Deletefine(id int) {
	sqlStatment := `DELETE FROM fine where id=$1`
	_, err := m.Db.Exec(sqlStatment, id)
	if err != nil {
		panic(err)
	}
}

func (m *Storage) DeleteOwnerFine(id int) {
	sql := `DELETE FROM fine where parkinglot_id IN (SELECT id from parkinglot WHERE owner_id=$1)`
	_, err := m.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
}
