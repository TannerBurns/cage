package models

import "database/sql"

type Game struct {
	ID         int     `json:"id"`
	Active     bool    `json:"active"`
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
	query := `INSERT INTO games (active, game_id, name, category, max_players, 
				minimum, maximum, interval, rules, notes) 
			  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
              RETURNING id;`

	err = db.QueryRow(query, gt.Active, gt.GameID, gt.Name,
		gt.Category, gt.MaxPlayers, gt.Minimum, gt.Maximum,
		gt.Interval, gt.Rules, gt.Notes).Scan(&gt.ID)
	if err != nil {
		return
	}
	return
}

func (gt *Game) UpdateGame(db *sql.DB) (err error) {
	query := `UPDATE games
				SET active=$1,
					game_id=$2,
					name=$3,
					category=$4,
				    max_players=$5,
					minimum=$6,
					maximum=$7,
					interval=$8,
					rules=$9,
					notes=$10
				WHERE game_id=$11
				RETURNING id;`
	err = db.QueryRow(query, gt.Active, gt.GameID, gt.Name,
		gt.Category, gt.MaxPlayers, gt.Minimum, gt.Maximum,
		gt.Interval, gt.Rules, gt.Notes, gt.GameID).Scan(&gt.ID)
	if err != nil {
		return
	}
	return
}

func (gt *Game) GetGame(db *sql.DB) (err error) {
	query := `SELECT * FROM games WHERE game_id=$1`

	err = db.QueryRow(query, gt.GameID).Scan(&gt.ID, &gt.Active, &gt.Created, &gt.GameID, &gt.Name,
		&gt.Category, &gt.MaxPlayers, &gt.Minimum, &gt.Maximum,
		&gt.Interval, &gt.Rules, &gt.Notes)

	if err != nil {
		return
	}
	return
}

func (gt *Game) SetState(db *sql.DB) (err error) {
	query := `UPDATE games SET active=$1 WHERE game_id=$2 RETURNING id;`

	err = db.QueryRow(query, gt.Active, gt.GameID).Scan(&gt.ID)
	if err != nil {
		return
	}
	return
}
