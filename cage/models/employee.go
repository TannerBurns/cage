package models

import (
	"database/sql"
)

type Employee struct {
	Id       int    `json:"id"`
	First    string `json:"first"`
	Last     string `json:"last"`
	Alias    string `json:"alias"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Birthday string `json:"birthday"`
	Joined   string `json:"joined"`
	Address  string `json:"address"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Region   string `json:"region"`
	Notes    string `json:"notes"`
	Picture  string `json:"picture"`
}

func (employee *Employee) CreateEmployee(db *sql.DB) (err error) {
	query := `INSERT INTO employees (first, last, alias, email, phone, birthday, address, city, state, zip, region,
	          notes, picture) 
			  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
              RETURNING id;`

	err = db.QueryRow(query, employee.First, employee.Last, employee.Alias,
		employee.Email, employee.Phone, employee.Birthday,
		employee.Address, employee.City, employee.State, employee.Zip,
		employee.Region, employee.Notes, employee.Picture).Scan(&employee.Id)
	if err != nil {
		return
	}
	return
}

func (employee *Employee) GetEmployee(db *sql.DB) (err error) {
	query := `SELECT * FROM employees WHERE id=$1`

	err = db.QueryRow(query, employee.Id).Scan(&employee.Id, &employee.First, &employee.Last, &employee.Alias,
		&employee.Email, &employee.Phone, &employee.Birthday, &employee.Joined,
		&employee.Address, &employee.City, &employee.State, &employee.Zip,
		&employee.Region, &employee.Notes, &employee.Picture)
	if err != nil {
		return
	}
	return
}

func (employee *Employee) GetEmployeeByFullName(db *sql.DB) (err error) {
	query := `SELECT * FROM employees WHERE first=$1 AND last=$2`

	err = db.QueryRow(query, employee.First, employee.Last).Scan(&employee.Id, &employee.First, &employee.Last, &employee.Alias,
		&employee.Email, &employee.Phone, &employee.Birthday, &employee.Joined,
		&employee.Address, &employee.City, &employee.State, &employee.Zip,
		&employee.Region, &employee.Notes, &employee.Picture)
	if err != nil {
		return
	}
	return
}

func (employee *Employee) GetRoles(db *sql.DB) (roles Roles, err error) {
	roles.EmployeeID = employee.Id
	err = roles.GetRoles(db)
	if err != nil {
		return
	}
	return
}

func (employee *Employee) UpdateEmployee(db *sql.DB) (err error) {
	query := `UPDATE employees
				SET first=$1,
					last=$2,
					alias=$3,
					email=$4,
				    phone=$5,
					birthday=$6,
					address=$7,
					city=$8,
					state=$9,
					zip=$10,
					region=$11,
					notes=$12,
					picture=$13
				WHERE id=$14`
	_, err = db.Query(query, employee.First, employee.Last, employee.Alias,
		employee.Email, employee.Phone, employee.Birthday,
		employee.Address, employee.City, employee.State, employee.Zip,
		employee.Region, employee.Notes, employee.Picture, employee.Id)
	if err != nil {
		return
	}
	return
}
