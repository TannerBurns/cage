package models

import (
	"math"
)

type Manager struct {
	Roster []CheckedInPlayers
}

type CheckedInPlayers struct {
	ID        int
	timer     Timer
	PlayTime  int
	ErrorChan chan error
}

func (check *CheckedInPlayers) Play() (err error) {
	check.ErrorChan = make(chan error)

	go func() {
		/*
			err = membership.GetPlayTime(db)
			if err != nil {
				//DO SOMETHING WITH THIS ERROR
				fmt.Println(err)
				//return
			}*/
		// make sure player has playtime before check in
		/*check.PlayTime = membership.PlayTime
		if check.PlayTime > 0 {
			check.timer.Start()
		}*/
		check.timer.Start()
		defer func() {

		}()
		for {
			select {
			default:
				// TODO: CHECK IF THE PLAYTIME == ELAPSED TIME === TIME IS UP!
				/*if check.PlayTime == int(check.timer.Elapsed()) {
					fmt.Println("Player is out of playtime!")
					return
				}*/
			case <-check.ErrorChan:
				check.timer.Stop()
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
				membership.PlayTime += int(math.Round(check.timer.TotalElapsed))
				membership.Amount = membership.PlayTime / 360
				err = membership.UpdatePlayTime(db)
				return
			}
		}
	}()

	return
}

func (manager *Manager) CheckIn(PlayerID int) (err error) {
	checkIn := CheckedInPlayers{ID: PlayerID}
	err = checkIn.Play()
	manager.Roster = append(manager.Roster, checkIn)
	return
}

func (manager *Manager) CheckOut(PlayerID int) {
	for ind := range manager.Roster {
		if manager.Roster[ind].ID == PlayerID {
			close(manager.Roster[ind].ErrorChan)
			if ind == 0 {
				manager.Roster = append(manager.Roster[:ind+1], manager.Roster[ind+1:]...)
			} else {
				manager.Roster = append(manager.Roster[:ind], manager.Roster[ind+1:]...)
			}
		}
	}
}
