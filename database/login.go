package database

import "golang.org/x/crypto/bcrypt"

func (m *Storage) Login(email string, password string) (string, bool) {

	sql := `SELECT password from userr WHERE email=$1`
	row := m.Db.QueryRow(sql, email)
	var hashedpassword string
	err := row.Scan(&hashedpassword)
	if err != nil {
		return "username or psssword inncorrect", false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	if err != nil {
		return "user not allowed", false
	}
	return "logged in", true
}
