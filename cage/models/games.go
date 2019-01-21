package models

import "database/sql"

type Game struct {
	ID           int     `json:"id"`
	Created      string  `json:"created"`
	GameID       int     `json:"game_id"`
	Name         string  `json:"name"`
	GameCategory string  `json:"category"`
	MaxPlayers   int     `json:"max_players"`
	Minimum      float32 `json:"minimum"`
	Maximum      float32 `json:"maximum"`
	Interval     int     `json:"interval"`
	Rules        string  `json:"rules"`
	Notes        string  `json:"notes"`
}

func (gt *Game) CreateTable(db *sql.DB) (err error) {
	query := `INSERT INTO games (game_id, name, category, max_players, 
				minimum, maximum, interval, rules, notes) 
			  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) 
              RETURNING id;`

	err = db.QueryRow(query, gt.GameID, gt.Name,
		gt.GameCategory, gt.MaxPlayers, gt.Minimum, gt.Maximum,
		gt.Interval, gt.Rules, gt.Notes).Scan(&gt.ID)
	if err != nil {
		return
	}
	return
}

func (gt *Game) UpdateTable(db *sql.DB) (err error) {
	query := `UPDATE games
				SET game_id=$1,
					name=$2,
					category=$3,
				    max_players=$4,
					minimum=$5,
					maximum=$6,
					interval=$7,
					rules=$8,
					notes=$9,
				WHERE id=$10`
	_, err = db.Query(query, gt.GameID, gt.Name,
		gt.GameCategory, gt.MaxPlayers, gt.Minimum, gt.Maximum,
		gt.Interval, gt.Rules, gt.Notes, gt.ID)
	if err != nil {
		return
	}
	return
}

func (gt *Game) GetTable(db *sql.DB) (err error) {
	query := `SELECT * FROM games WHERE id=$1`

	err = db.QueryRow(query, gt.ID).Scan(&gt.ID, &gt.Created, &gt.GameID, &gt.Name,
		&gt.GameCategory, &gt.MaxPlayers, &gt.Minimum, &gt.Maximum,
		&gt.Interval, &gt.Rules, &gt.Notes)

	if err != nil {
		return
	}
	return
}
