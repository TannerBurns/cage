package controllers

import (
	"encoding/json"
	"math"
	"net/http"

	"../models"
	"github.com/gorilla/context"
)

func (c *Controller) GetRoster(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//connect to database
	postclient := models.PostgresConnection{}
	db, err := postclient.Connect()
	if err != nil {
		error := models.RespError{Error: "Failed to connect, cannot reach database"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}
	defer db.Close()

	auth := models.Authentication{Decoded: context.Get(req, "decoded")}
	ok, err := auth.Authorize(db, 3)
	if err != nil {
		error := models.RespError{Error: "Failed to authorize, error during authorization"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}
	if !ok {
		error := models.RespError{Error: "Failed to authorize, error during authorization. Make sure you have permissions to use this route."}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 401)
		c.Logger.Logging(req, 401)
		return
	}

	roster := models.ManagerRoster{}
	for id, player := range c.Manager.Roster {
		tp := player.PlayerTimer.Elapsed()
		membership := models.Membership{PlayerID: id}
		err = membership.GetMembership(db)
		if err != nil {
			roster.Responses = append(roster.Responses, models.ManagerResp{
				PlayerID:   id,
				Status:     "checked in, failed to retrieve player membership",
				TimePlayed: int(tp),
				AmountOwed: int(math.Round(tp / 360))})
		} else {
			roster.Responses = append(roster.Responses, models.ManagerResp{
				PlayerID:   id,
				Status:     "checked in",
				TimePlayed: membership.PlayTime + int(tp),
				AmountOwed: membership.Amount - int(math.Round(tp/360))})
		}

	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(roster)
	return
}
