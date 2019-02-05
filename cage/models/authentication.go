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
	access_id, ok := dec["access_id"].(float64)
	if !ok {
		return
	}
	if int(access_id) <= aLevel {
		authorized = true
		return
	}
	return
}
