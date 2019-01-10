package models

import (
	"database/sql"
	"errors"
)

type Login struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employee_id"`
	Username   string `json:"username"`
}

func (login *Login) CreateLogin(db *sql.DB, password string) (err error) {
	query := `INSER INTO logins (employee_id, username, password)
			  VALUES ($1, $2, crypt($3, gen_salt('bf')));`

	_, err = db.Query(query, login.EmployeeID, login.Username, password)
	if err != nil {
		return
	}
	return
}

func (login *Login) UpdatePassword(db *sql.DB, oldpassword string, newpassword string, newpasswordcheck string) (err error) {
	query := `UPDATE logins
				SET password=crypt($1, gen_salt('bf'))
				WHERE username=$2 AND password=crypt(oldpassword, password);`

	if newpassword == newpasswordcheck {
		_, err = db.Query(query, newpassword, login.Username)
		if err != nil {
			return
		}
	} else {
		err = errors.New("New passwords do not match")
		return
	}
	return
}

func (login *Login) ValidateLogin(db *sql.DB, password string) (err error) {
	query := `SELECT id, employee_id FROM logins WHERE username=$1 AND password=crypt($2, password);`

	err = db.QueryRow(query, login.Username, password).Scan(&login.ID, &login.EmployeeID)
	return
}
