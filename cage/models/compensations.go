package models

import "database/sql"

type Compensations struct {
	EmployeeID int
	PlayerID   int
	Amount     int
	CompType   string
	Notes      string
}

func (comp *Compensations) CreateCompensation(db *sql.DB) (err error) {
	var compID int
	query := `INSERT INTO compensations (employee_id, player_id, amount, type, notes) 
				VALUES($1, $2, $3, $4, $5) RETURNING id;`

	err = db.QueryRow(query, comp.EmployeeID, comp.PlayerID, comp.Amount, comp.CompType, comp.Notes).Scan(&compID)
	if err != nil {
		return
	}

	if comp.CompType == "PlayTime" {
		membership := Membership{PlayerID: comp.PlayerID, Amount: comp.Amount, PlayTime: (int(comp.Amount) / 10) * 60 * 60}
		err = membership.AddPlayTime(db, "Unknown")
		if err != nil {
			return
		}
	}
	return
}
