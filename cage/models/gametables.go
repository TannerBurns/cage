package models

import "database/sql"

type GameTables struct {
	ID         int     `json:"id"`
	Created    string  `json:"created"`
	GameID     int     `json:"game_id"`
	Name       string  `json:"game_name"`
	GameType   string  `json:"game_typer"`
	MaxPlayers int     `json:"max_gts"`
	Minimum    float32 `json:"minimum"`
	Maximum    float32 `json:"maximum"`
	Interval   int     `json:"interval"`
	Rules      string  `json:"rules"`
	Notes      string  `json:"notes"`
}

func (gt *GameTables) CreateTable(db *sql.DB) (err error) {
	query := `INSERT INTO gametables (created, game_id, game_name, game_type, max_gts, 
				minimum, maximum, interval, rules, notes) 
			  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
              RETURNING id;`

	err = db.QueryRow(query, gt.Created, gt.GameID, gt.Name,
		gt.GameType, gt.MaxPlayers, gt.Minimum, gt.Maximum,
		gt.Interval, gt.Rules, gt.Notes).Scan(&gt.ID)
	if err != nil {
		return
	}
	return
}

func (gt *GameTables) UpdateTable(db *sql.DB) (err error) {
	query := `UPDATE gametables
				SET created=$1,
					game_id=$2,
					game_name=$3,
					game_type=$4,
				    max_gts=$5,
					minimum=$6,
					maximum=$7,
					interval=$8,
					rules=$9,
					notes=$10,
				WHERE id=$11`
	_, err = db.Query(query, gt.Created, gt.GameID, gt.Name,
		gt.GameType, gt.MaxPlayers, gt.Minimum, gt.Maximum,
		gt.Interval, gt.Rules, gt.Notes, gt.ID)
	if err != nil {
		return
	}
	return
}

func (gt *GameTables) GetTable(db *sql.DB) (err error) {
	query := `SELECT * FROM gametables WHERE id=$1`

	err = db.QueryRow(query, gt.ID).Scan(&gt.ID, &gt.Created, &gt.GameID, &gt.Name,
		&gt.GameType, &gt.MaxPlayers, &gt.Minimum, &gt.Maximum,
		&gt.Interval, &gt.Rules, &gt.Notes)

	if err != nil {
		return
	}
	return
}
