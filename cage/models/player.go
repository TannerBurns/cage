package models

import (
	"database/sql"
)

type Players struct {
	Players []Player `json:"players"`
}

type Player struct {
	Id           int    `json:"id"`
	First        string `json:"first"`
	Last         string `json:"last"`
	Alias        string `json:"alias"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Birthday     string `json:"birthday"`
	Joined       string `json:"joined"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Region       string `json:"region"`
	Referral     bool   `json:"referral"`
	ReferredBy   string `json:"referred_by"`
	ReferredType string `json:"referred_type"`
	Banned       bool   `json:"banned"`
	Notes        string `json:"notes"`
	Picture      string `json:"picture"`
}

func (player *Player) CreatePlayer(db *sql.DB) (err error) {
	query := `INSERT INTO players (first, last, alias, email, phone, birthday, address, city, state, zip, region,
	          referral, referredby, referredtype, banned, notes, picture) 
			  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) 
              RETURNING id;`
	err = db.QueryRow(query, player.First, player.Last, player.Alias,
		player.Email, player.Phone, player.Birthday,
		player.Address, player.City, player.State, player.Zip,
		player.Region, player.Referral, player.ReferredBy, player.ReferredType,
		player.Banned, player.Notes, player.Picture).Scan(&player.Id)
	if err != nil {
		return
	}
	return
}

func (player *Player) UpdatePlayer(db *sql.DB) (err error) {
	query := `UPDATE players
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
					referral=$12,
					referredby=$13,
					referredtype=$14,
					banned=$15,
					notes=$16,
					picture=$17
				WHERE id=$18`
	_, err = db.Query(query, player.First, player.Last, player.Alias,
		player.Email, player.Phone, player.Birthday,
		player.Address, player.City, player.State, player.Zip,
		player.Region, player.Referral, player.ReferredBy, player.ReferredType,
		player.Banned, player.Notes, player.Picture, player.Id)
	if err != nil {
		return
	}
	return
}

func (player *Player) GetPlayer(db *sql.DB) (err error) {
	query := `SELECT * FROM players WHERE id=$1`

	err = db.QueryRow(query, player.Id).Scan(&player.Id, &player.First, &player.Last, &player.Alias,
		&player.Email, &player.Phone, &player.Birthday, &player.Joined,
		&player.Address, &player.City, &player.State, &player.Zip,
		&player.Region, &player.Referral, &player.ReferredBy, &player.ReferredType,
		&player.Banned, &player.Notes, &player.Picture)
	if err != nil {
		return
	}
	return
}

func (players *Players) Search(db *sql.DB, first string, last string) (err error) {
	query := `SELECT * FROM players WHERE first=$1 AND last=$2 ORDER BY id DESC`

	rows, err := db.Query(query, first, last)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		player := Player{}
		err = rows.Scan(&player.Id, &player.First, &player.Last, &player.Alias,
			&player.Email, &player.Phone, &player.Birthday, &player.Joined,
			&player.Address, &player.City, &player.State, &player.Zip,
			&player.Region, &player.Referral, &player.ReferredBy, &player.ReferredType,
			&player.Banned, &player.Notes, &player.Picture)
		if err != nil {
			return
		}
		players.Players = append(players.Players, player)
	}
	return
}
