package models

import (
	"database/sql"
	"math"
)

type Manager struct {
	Roster map[int]*CheckedInPlayers
}

type CheckedInPlayers struct {
	ID          int
	Games       *Games
	PlayerTimer *Timer
}

type Games struct {
	GameTimer *Timer
	Game      Game
}

func NewManager() *Manager {
	manager := &Manager{}
	manager.Roster = make(map[int]*CheckedInPlayers)
	return manager
}

func (manager *Manager) CheckIn(PlayerID int) (err error) {
	checkIn := CheckedInPlayers{ID: PlayerID}
	checkIn.PlayerTimer = NewTimer()
	checkIn.PlayerTimer.Start()
	manager.Roster[PlayerID] = &checkIn
	return
}

func (manager *Manager) CheckOut(db *sql.DB, PlayerID int) (err error) {
	manager.Roster[PlayerID].PlayerTimer.Stop()
	if manager.Roster[PlayerID].Games != nil {
		if int(manager.Roster[PlayerID].Games.GameTimer.Elapsed()) > 0 {
			manager.LeaveGame(db, PlayerID)
		}
	}
	membership := Membership{PlayerID: PlayerID}
	err = membership.GetPlayTime(db)
	if err != nil {
		return
	}
	membership.PlayTime = int(manager.Roster[PlayerID].PlayerTimer.TotalElapsed)
	membership.Amount = int(math.Round(manager.Roster[PlayerID].PlayerTimer.TotalElapsed / 360))
	err = membership.UpdatePlayTime(db)
	delete(manager.Roster, PlayerID)
	return
}

func (manager *Manager) JoinGame(PlayerID int, game Game) {
	manager.Roster[PlayerID].Games.GameTimer = NewTimer()
	manager.Roster[PlayerID].Games.GameTimer.Start()
	manager.Roster[PlayerID].Games.Game = game
	// write to database
}

func (manager *Manager) LeaveGame(db *sql.DB, PlayerID int) (err error) {
	manager.Roster[PlayerID].Games.GameTimer.Stop()
	err = manager.Roster[PlayerID].CreateTransaction(db)
	if err != nil {
		return
	}
	manager.Roster[PlayerID].Games = nil
	return
}

func (manager *Manager) MoveGame(db *sql.DB, PlayerID int, game Game) (err error) {
	err = manager.LeaveGame(db, PlayerID)
	if err != nil {
		return
	}
	manager.JoinGame(PlayerID, game)
	return
}

func (check *CheckedInPlayers) CreateTransaction(db *sql.DB) (err error) {
	query := `INSERT INTO playertransactions (player_id, game_id, category, time_played)
	VALUES($1, $2, $3, $4)`

	_, err = db.Query(query, check.ID, check.Games.Game.GameID, check.Games.Game.Category, int(check.Games.GameTimer.TotalElapsed))
	return
}
