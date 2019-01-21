package models

import (
	"database/sql"
	"strconv"
	"strings"
)

type PlayerBulk struct {
	Players []PlayerCard `json:"players"`
}

type PlayerCard struct {
	Player     Player     `json:"player"`
	Membership Membership `json:"membership"`
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

func (players *PlayerBulk) Search(db *sql.DB, q string) (err error) {
	var rows *sql.Rows
	i, err := strconv.Atoi(q)
	if err != nil {
		if strings.Contains(q, "@") && strings.Contains(q, ".") {
			// email
			query := `SELECT * FROM players WHERE email=$1 ORDER BY id DESC`
			rows, err = db.Query(query, q)
			if err != nil {
				return err
			}
		} else if len(strings.Split(q, " ")) > 1 {
			// name
			query := `SELECT * FROM players WHERE first=$1 AND last=$2 ORDER BY id DESC`
			rows, err = db.Query(query, strings.Split(q, " ")[0], strings.Split(q, " ")[1])
			if err != nil {
				return err
			}
		}
	} else {
		dcount := 0
		j := i
		for j != 0 {
			j /= 10
			dcount++
		}
		if dcount >= 7 {
			// phone number
			query := `SELECT * FROM players WHERE phone=$1 ORDER BY id DESC`
			rows, err = db.Query(query, i)
			if err != nil {
				return err
			}
		} else {
			// id
			query := `SELECT * FROM players WHERE id=$1`
			rows, err = db.Query(query, i)
			if err != nil {
				return err
			}

		}
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
		membership := Membership{PlayerID: player.Id}
		err = membership.GetMembership(db)
		if err != nil {
			return
		}
		playercard := PlayerCard{Player: player, Membership: membership}
		players.Players = append(players.Players, playercard)
	}
	return
}
