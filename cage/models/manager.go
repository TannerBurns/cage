package models

import (
	"math"
)

type Manager struct {
	Roster map[int]CheckedInPlayers
}

type CheckedInPlayers struct {
	ID          int
	Game        Game
	PlayerTimer *Timer
	ErrorChan   chan string
}

func NewManager() *Manager {
	manager := &Manager{}
	manager.Roster = make(map[int]CheckedInPlayers)
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
	manager.Roster[PlayerID] = checkIn
	return
}

func (manager *Manager) CheckOut(PlayerID int) {
	close(manager.Roster[PlayerID].ErrorChan)
	delete(manager.Roster, PlayerID)
}
