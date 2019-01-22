package models

import "database/sql"

type Game struct {
	ID         int     `json:"id"`
	Created    string  `json:"created"`
	GameID     string  `json:"game_id"`
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	MaxPlayers int     `json:"max_players"`
	Minimum    float32 `json:"minimum"`
	Maximum    float32 `json:"maximum"`
	Interval   int     `json:"interval"`
	Rules      string  `json:"rules"`
	Notes      string  `json:"notes"`
}

func (gt *Game) CreateGame(db *sql.DB) (err error) {
	query := `INSERT INTO games (game_id, name, category, max_players, 
				minimum, maximum, interval, rules, notes) 
			  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) 
              RETURNING id;`

	err = db.QueryRow(query, gt.GameID, gt.Name,
		gt.Category, gt.MaxPlayers, gt.Minimum, gt.Maximum,
		gt.Interval, gt.Rules, gt.Notes).Scan(&gt.ID)
	if err != nil {
		return
	}
	return
}

func (gt *Game) UpdateGame(db *sql.DB) (err error) {
	query := `UPDATE games
				SET game_id=$1,
					name=$2,
					category=$3,
				    max_players=$4,
					minimum=$5,
					maximum=$6,
					interval=$7,
					rules=$8,
					notes=$9
				WHERE game_id=$10
				RETURNING id;`
	err = db.QueryRow(query, gt.GameID, gt.Name,
		gt.Category, gt.MaxPlayers, gt.Minimum, gt.Maximum,
		gt.Interval, gt.Rules, gt.Notes, gt.GameID).Scan(&gt.ID)
	if err != nil {
		return
	}
	return
}

func (gt *Game) GetGame(db *sql.DB) (err error) {
	query := `SELECT * FROM games WHERE game_id=$1`

	err = db.QueryRow(query, gt.GameID).Scan(&gt.ID, &gt.Created, &gt.GameID, &gt.Name,
		&gt.Category, &gt.MaxPlayers, &gt.Minimum, &gt.Maximum,
		&gt.Interval, &gt.Rules, &gt.Notes)

	if err != nil {
		return
	}
	return
}
