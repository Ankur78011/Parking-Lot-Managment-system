package database

import (
	"ankur/parkinlot/models"
)

func (m *Storage) Getparkinglot(userId int) []models.Parkinglot {
	sqlQuery := `SELECT *FROM parkinglot WHERE owner_id=$1`
	rows, err := m.Db.Query(sqlQuery, userId)
	if err != nil {
		panic(err)
	}
	var Parkinglot []models.Parkinglot
	for rows.Next() {
		parkinglot := models.Parkinglot{}
		err := rows.Scan(&parkinglot.Id, &parkinglot.Name, &parkinglot.Address, &parkinglot.OwnerId)
		if err != nil {
			panic(err)
		}
		Parkinglot = append(Parkinglot, parkinglot)
	}
	return Parkinglot
}

func (m *Storage) InserParkinglot(name string, address string, id int) int64 {
	sqlStatment := `INSERT INTO parkinglot (name,address,owner_id) values($1,$2,$3) `
	result, err := m.Db.Exec(sqlStatment, name, address, id)
	if err != nil {
		panic(err)
	}
	newId, err := result.LastInsertId()
	return newId
}

func (m *Storage) DeleteParkinglot(id int, user_id int) {
	sqlfine := `DELETE FROM fine WHERE parkinglot_id IN (SELECT id FROM parkinglot WHERE owner_id=$1)`
	_, err := m.Db.Exec(sqlfine, user_id)
	if err != nil {
		panic(err)
	}
	sqlQuery := `DELETE FROM permit WHERE parkinglot_id IN (SELECT id from parkinglot WHERE owner_id=$1)`
	_, err = m.Db.Exec(sqlQuery, user_id)
	if err != nil {
		panic(err)
	}
	sqlStatment := `DELETE FROM parkinglot WHERE id=$1 AND owner_id=$2`
	_, err = m.Db.Exec(sqlStatment, id, user_id)
	if err != nil {
		panic(err)
	}
}
func (m *Storage) UpdateParkinglot(name string, address string, owner_id int, id int, user_id int) {
	sqlStatment := `UPDATE parkinglot SET name=$1,address=$2,owner_id=$3 WHERE id=$4 AND owner_id=$5`
	_, err := m.Db.Exec(sqlStatment, name, address, owner_id, id, user_id)
	if err != nil {
		panic(err)
	}
}

func (m *Storage) DeleteOwnerParkingLot(id int) {
	sqlStatment := `DELETE FROM parkinglot where owner_id=$1`
	_, err := m.Db.Exec(sqlStatment, id)
	if err != nil {
		panic(err)
	}
}
