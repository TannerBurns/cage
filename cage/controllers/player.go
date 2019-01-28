package controllers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

/*
GetPlayer - retrieve information for a player by id. returns 400, 401, 404 for errors
*/
func (c *Controller) GetPlayer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//connect to database
	db, err := c.Session.Connect()
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

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	player := models.Player{Id: id}

	err = player.GetPlayer(db)
	if err != nil {
		error := models.RespError{Error: "Failed to find player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(player)
	return
}

/*
CreatePlayer - create a new player. returns 400, 401, 404 for errors
*/
func (c *Controller) CreatePlayer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//connect to database
	db, err := c.Session.Connect()
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

	var player models.Player
	err = json.NewDecoder(req.Body).Decode(&player)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}

	err = player.CreatePlayer(db)
	if err != nil {
		error := models.RespError{Error: "Failed to create a new player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(player)
	return
}

/*
UpdatePlayer - Update player information by id. returns 400, 401, 404 for errors
*/
func (c *Controller) UpdatePlayer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//connect to database
	db, err := c.Session.Connect()
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

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	var player models.Player
	err = json.NewDecoder(req.Body).Decode(&player)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	player.Id = id
	err = player.UpdatePlayer(db)
	if err != nil {
		error := models.RespError{Error: "Failed to update player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(player)
	return
}

/*
Search - search for a player by id, first last, email, or phone numbers. returns 400, 401, 404 for errors
*/
func (c *Controller) Search(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//connect to database
	db, err := c.Session.Connect()
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

	q := req.URL.Query().Get("q")
	if q == "" {
		error := models.RespError{Error: "Failed to find 'q' variable"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	players := models.PlayerBulk{}
	err = players.Search(db, q)
	if err != nil {
		error := models.RespError{Error: "Failed to find any players matching the search content"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(players)
	return
}

/*
CheckIn - check in a player, returns 400, 401, 404 for errors
*/
func (c *Controller) CheckIn(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//connect to database
	db, err := c.Session.Connect()
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

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	err = c.Manager.CheckIn(id)
	if err != nil {
		error := models.RespError{Error: "Failed to check in player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}

	player := models.Player{Id: id}
	err = player.GetPlayer(db)
	if err != nil {
		error := models.RespError{Error: "Failed to find player to check in"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	membership := models.Membership{PlayerID: id}
	err = membership.GetMembership(db)
	if err != nil {
		json.NewEncoder(w).Encode(models.ManagerResp{
			PlayerID:        id,
			First:           player.First,
			Last:            player.Last,
			Status:          "checked in, error retrieving membership information",
			CheckedInTime:   0,
			TotalTimePlayed: 0,
			AmountOwed:      0})
	} else {
		json.NewEncoder(w).Encode(models.ManagerResp{
			PlayerID:        id,
			First:           player.First,
			Last:            player.Last,
			Status:          "checked in",
			CheckedInTime:   0,
			TotalTimePlayed: membership.PlayTime,
			AmountOwed:      membership.Amount})
	}

	return
}

/*
CheckOut - check out player, returns 400, 401, 404 for errors
*/
func (c *Controller) CheckOut(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//connect to database
	db, err := c.Session.Connect()
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

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	player := models.Player{Id: id}
	err = player.GetPlayer(db)
	if err != nil {
		error := models.RespError{Error: "Failed to find player to check out"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}

	tp := c.Manager.Roster[id].PlayerTimer.Elapsed()
	c.Manager.CheckOut(db, id)
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	membership := models.Membership{PlayerID: id}
	err = membership.GetMembership(db)
	if err != nil {
		json.NewEncoder(w).Encode(models.ManagerResp{
			PlayerID:        id,
			First:           player.First,
			Last:            player.Last,
			Status:          "checked out, failed to retrieve player membership",
			CheckedInTime:   int(tp),
			TotalTimePlayed: 0,
			AmountOwed:      int(math.Round(tp / 360))})
	} else {
		json.NewEncoder(w).Encode(models.ManagerResp{
			PlayerID:        id,
			First:           player.First,
			Last:            player.Last,
			Status:          "checked out",
			CheckedInTime:   int(tp),
			TotalTimePlayed: membership.PlayTime + int(tp),
			AmountOwed:      membership.Amount - int(math.Round(tp/360))})
	}
	return
}
