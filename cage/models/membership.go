package models

import (
	"database/sql"
	"math"
)

type Membership struct {
	ID           int    `json:"id"`
	PlayerID     int    `json:"player_id"`
	EmployeeID   int    `json:"employee_id"`
	Created      string `json:"created"`
	Active       bool   `json:"active"`
	ActiveDate   string `json:"active_date"`
	DeactiveDate string `json:"deactive_date"`
	Amount       int    `json:"amount"`
	PlayTime     int    `json:"playtime"`
	Notes        string `json:"notes"`
}

type MemberTransaction struct {
	MemberID   int
	PlayerID   int
	EmployeeID int
	Created    string
	Type       string
	Amount     int
}

func (membership *Membership) CreateMembership(db *sql.DB, TransactionType string) (err error) {
	query := `INSERT INTO memberships (player_id, employee_id, active, 
                         active_date, deactive_date, amount, playtime, notes)
              VALUES($1, $2, $3, $4, $5, $6, $7, $8)
			  RETURNING id;`

	err = db.QueryRow(query, membership.PlayerID, membership.EmployeeID,
		membership.Active, membership.ActiveDate, membership.DeactiveDate,
		membership.Amount, membership.PlayTime, membership.Notes).Scan(&membership.ID)
	if err != nil {
		return
	}

	query = `INSERT INTO membertransactions (membership_id, player_id, employee_id, type, amount)
			  VALUES($1, $2, $3, $4, $5)`

	_, err = db.Query(query, membership.ID, membership.PlayerID, membership.EmployeeID,
		TransactionType, membership.Amount)

	if err != nil {
		return
	}
	return
}

func (membership *Membership) AddPlayTime(db *sql.DB, TransactionType string) (err error) {
	query := `UPDATE memberships 
				SET amount = amount + $1,
					playtime = playtime - $2
				WHERE player_id=$3
				RETURNING id`
	err = db.QueryRow(query, membership.Amount, membership.PlayTime, membership.PlayerID).Scan(&membership.ID)
	if err != nil {
		return
	}
	query = `INSERT INTO membertransactions (membership_id, player_id, employee_id, type, amount)
			  VALUES($1, $2, $3, $4, $5)`

	_, err = db.Query(query, membership.ID, membership.PlayerID, membership.EmployeeID,
		TransactionType, membership.Amount)
	return
}

func (membership *Membership) UpdatePlayTime(db *sql.DB) (err error) {
	query := `UPDATE memberships 
				SET playtime = playtime + $1
				WHERE player_id=$2
				RETURNING playtime`
	err = db.QueryRow(query, membership.PlayTime, membership.PlayerID).Scan(&membership.PlayTime)
	if err != nil {
		return
	}
	query = `UPDATE memberships 
				SET amount = $1
				WHERE player_id=$2
				RETURNING id`
	err = db.QueryRow(query, 0-int(math.Round(float64(membership.PlayTime)/360)), membership.PlayerID).Scan(&membership.ID)
	if err != nil {
		return
	}
	return
}

func (membership *Membership) GetMembership(db *sql.DB) (err error) {
	query := `SELECT * FROM memberships WHERE player_id=$1`
	err = db.QueryRow(query, membership.PlayerID).Scan(&membership.ID, &membership.PlayerID, &membership.EmployeeID, &membership.Created,
		&membership.Active, &membership.ActiveDate, &membership.DeactiveDate,
		&membership.Amount, &membership.PlayTime, &membership.Notes)
	return
}

func (membership *Membership) GetPlayTime(db *sql.DB) (err error) {
	query := `SELECT playtime FROM memberships WHERE player_id=$1`
	err = db.QueryRow(query, membership.PlayerID).Scan(&membership.PlayTime)
	return
}
