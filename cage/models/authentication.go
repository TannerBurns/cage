package models

import (
	"database/sql"

	jwt "github.com/dgrijalva/jwt-go"
)

type Authentication struct {
	EmployeeID int
	Decoded    interface{}
}

func (auth *Authentication) Authorize(db *sql.DB, aLevel int) (authorized bool, err error) {
	authorized = false

	dec, ok := auth.Decoded.(jwt.MapClaims)
	if !ok {
		return
	}
	id, ok := dec["employee_id"].(float64)
	if !ok {
		return
	}

	employee := &Employee{Id: int(id)}
	err = employee.GetEmployee(db)
	if err != nil {
		return
	}
	roles, err := employee.GetRoles(db)
	if err != nil {
		return
	}
	for _, r := range roles.Roles {
		authLevel, err := r.GetAccessLevel(db)
		if err != nil {
			break
		}
		if authLevel <= aLevel {
			authorized = true
			auth.EmployeeID = employee.Id
			break
		}
	}
	return
}
