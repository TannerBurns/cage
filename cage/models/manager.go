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
	ErrorChan   chan string
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

func (check *CheckedInPlayers) Play() (err error) {
	check.ErrorChan = make(chan string)
	go func() {
		check.PlayerTimer.Start()
		for {
			select {
			case <-check.ErrorChan:
				check.PlayerTimer.Stop()
				postclient := PostgresConnection{}
				db, err := postclient.Connect()
				if err != nil {
					return
				}
				defer db.Close()
				membership := Membership{PlayerID: check.ID}
				err = membership.GetPlayTime(db)
				if err != nil {
					return
				}
				membership.PlayTime = int(check.PlayerTimer.TotalElapsed)
				membership.Amount = int(math.Round(check.PlayerTimer.TotalElapsed / 360))
				err = membership.UpdatePlayTime(db)
				return
			}
		}
	}()
	return
}

func (manager *Manager) CheckIn(PlayerID int) (err error) {
	checkIn := CheckedInPlayers{ID: PlayerID}
	checkIn.PlayerTimer = NewTimer()
	err = checkIn.Play()
	manager.Roster[PlayerID] = &checkIn
	return
}

func (manager *Manager) CheckOut(PlayerID int) {
	if manager.Roster[PlayerID].Games != nil {
		if int(manager.Roster[PlayerID].Games.GameTimer.Elapsed()) > 0 {
			manager.LeaveGame(PlayerID)
		}
	}
	close(manager.Roster[PlayerID].ErrorChan)
	delete(manager.Roster, PlayerID)
}

func (manager *Manager) JoinGame(PlayerID int, game Game) {
	manager.Roster[PlayerID].Games.GameTimer = NewTimer()
	manager.Roster[PlayerID].Games.GameTimer.Start()
	manager.Roster[PlayerID].Games.Game = game
	// write to database
}

func (manager *Manager) LeaveGame(PlayerID int) (err error) {
	manager.Roster[PlayerID].Games.GameTimer.Stop()
	postclient := PostgresConnection{}
	db, err := postclient.Connect()
	if err != nil {
		return
	}
	defer db.Close()
	err = manager.Roster[PlayerID].CreateTransaction(db)
	if err != nil {
		return
	}
	return
}

func (manager *Manager) MoveGame(PlayerID int, game Game) (err error) {
	err = manager.LeaveGame(PlayerID)
	if err != nil {
		return
	}
	manager.JoinGame(PlayerID, game)
	return
}

func (check *CheckedInPlayers) CreateTransaction(db *sql.DB) (err error) {
	query := `INSERT INTO playertransactions (player_id, game_id, time_played)
	VALUES($1, $2, $3)`

	_, err = db.Query(query, check.ID, check.Games.Game.GameID, int(check.Games.GameTimer.TotalElapsed))
	return
}
