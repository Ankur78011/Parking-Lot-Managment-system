package database

import (
	"ankur/parkinlot/models"

	"golang.org/x/crypto/bcrypt"
)

func (m *Storage) GetOwners() []models.ParkingOwner {
	rows, err := m.Db.Query("Select id,usertype,username from userr where usertype='owner' ORDER BY id")
	if err != nil {
		panic(err)
	}
	var UserDetails []models.ParkingOwner
	for rows.Next() {
		user := models.ParkingOwner{}
		err := rows.Scan(&user.Id, &user.UserType, &user.Username)
		if err != nil {
			panic(err)
		}
		UserDetails = append(UserDetails, user)
	}
	return UserDetails
}
func (m *Storage) PostOwner(newname string, newemial string, newpassword string) int64 {
	hasspassword, err := bcrypt.GenerateFromPassword([]byte(newpassword), 8)
	sqlStatment := `INSERT INTO userr (usertype,username,email,password) values ('owner',$1,$2,$3)`
	result, err := m.Db.Exec(sqlStatment, newname, newemial, hasspassword)
	if err != nil {
		panic(err)
	}
	owner, err := result.LastInsertId()
	return owner
}
func (m *Storage) DeleteOwner(id int) {
	sqlStatment_2 := `DELETE FROM userr where id=$1 and usertype='owner'`
	_, err := m.Db.Exec(sqlStatment_2, id)
	if err != nil {
		panic(err)
	}
}
func (m *Storage) GetCustomer() []models.ParkingOwner {
	rows, err := m.Db.Query("Select id,usertype,username from userr where usertype='customer'")
	if err != nil {
		panic(err)
	}
	var UserDetails []models.ParkingOwner
	for rows.Next() {
		user := models.ParkingOwner{}
		err := rows.Scan(&user.Id, &user.UserType, &user.Username)
		if err != nil {
			panic(err)
		}
		UserDetails = append(UserDetails, user)
	}
	return UserDetails
}

func (m *Storage) PostCustomer(name string) int64 {
	sqlStatment := `INSERT INTO userr (usertype,username) values ('customer',$1)`
	result, err := m.Db.Exec(sqlStatment, name)
	if err != nil {
		panic(err)
	}
	newId, err := result.LastInsertId()
	return newId
}
func (m *Storage) Deletecustomer(id int) {
	sqlStatment := `DELETE FROM permit where customer_id=$1`
	_, err := m.Db.Exec(sqlStatment, id)
	if err != nil {
		panic(err)
	}
	sqlStatment_2 := `DELETE FROM userr where id=$1 and usertype='customer'`
	_, err = m.Db.Exec(sqlStatment_2, id)
	if err != nil {
		panic(err)
	}
}
func (m *Storage) CheckUserType(id int) {
	sqlquery := `SELECT usertype from userr WHERE id=$1 and usertype='owner'`
	row := m.Db.QueryRow(sqlquery, id)
	var utype string
	err := row.Scan(&utype)
	if err != nil {
		panic(err)

	}
}

func (m *Storage) CheckCustomer(name string) int {
	sqlQuery := `SELECT id FROM userr WHERE usertype='customer' AND username=$1 `
	var newId int
	row := m.Db.QueryRow(sqlQuery, name)
	err := row.Scan(&newId)
	if err != nil {
		return 0
	}
	return newId

}
func (m *Storage) InsertNewCustomer(name string) int {
	sqlQuery := `INSERT INTO userr (usertype,username) values('customer',$1) RETURNING id`
	res := m.Db.QueryRow(sqlQuery, name)
	var newId int
	err := res.Scan(&newId)
	if err != nil {
		panic(err)
	}
	return newId

}
func (m *Storage) UpdateCustomer(name string, id int) {
	sqlQuery := `UPDATE userr SET usertype='customer',username=$1 WHERE id=$2`
	_, err := m.Db.Exec(sqlQuery, name, id)
	if err != nil {
		panic(err)
	}
}
func (m *Storage) UpdateParkingOwner(name string, id int) {
	sqlQuery := `UPDATE userr SET usertype='owner',username=$1 WHERE id=$2`
	_, err := m.Db.Exec(sqlQuery, name, id)
	if err != nil {
		panic(err)
	}
}

func (m *Storage) GetUserId(emial string) int {
	sqlQuery := `SELECT id FROM userr WHERE email=$1`
	var id int
	row := m.Db.QueryRow(sqlQuery, emial)
	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}
	return id

}

func (m *Storage) ValidateUser(id float64) bool {
	sql := `SELECT username FROM userr WHERE id=$1`
	var username string
	row := m.Db.QueryRow(sql, id)
	err := row.Scan(&username)
	if err != nil {
		return false
	}
	return true
}
