package models

import (
	"database/sql"
	"fmt"
)

type Manager struct {
	Roster []CheckedInPlayers
}

type CheckedInPlayers struct {
	Id        int
	timer     Timer
	PlayTime  int
	ErrorChan chan error
}

func (check *CheckedInPlayers) Play(db *sql.DB) (err error) {
	check.ErrorChan = make(chan error)

	go func(db *sql.DB) {
		// TODO: GET PlayTime from membership
		fmt.Println("Starting")
		membership := Membership{PlayerID: check.Id}
		err := membership.GetPlayTime(db)
		if err != nil {
			//DO SOMETHING WITH THIS ERROR
			fmt.Println(err)
			//return
		}
		// make sure player has playtime before check in
		/*check.PlayTime = membership.PlayTime
		if check.PlayTime > 0 {
			check.timer.Start()
		}*/
		check.timer.Start()
		defer func() {
			// TODO: do teardown work

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
				// TODO: CLEANUP WHEN STOPPING
				fmt.Println("Stopping")
				check.timer.Stop()
				fmt.Println(check.Id, check.timer.TotalElapsed)
				return
			}
		}
	}(db)
	return
}

func (manager *Manager) CheckIn(db *sql.DB, PlayerID int) (err error) {
	checkIn := CheckedInPlayers{Id: PlayerID}
	err = checkIn.Play(db)
	manager.Roster = append(manager.Roster, checkIn)
	return
}

func (manager *Manager) CheckOut(PlayerID int) {
	for ind := range manager.Roster {
		if manager.Roster[ind].Id == PlayerID {
			close(manager.Roster[ind].ErrorChan)
			if ind == 0 {
				manager.Roster = append(manager.Roster[:ind+1], manager.Roster[ind+1:]...)
			} else {
				manager.Roster = append(manager.Roster[:ind], manager.Roster[ind+1:]...)
			}
		}
	}
}
