package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"../models"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

/*
CreateMembership - create a new membership for a player. returns 400, 401, 404 for errors
*/
func (c *Controller) CreateMembership(w http.ResponseWriter, req *http.Request) {
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

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	transtype := req.URL.Query().Get("type")
	if transtype == "" {
		transtype = "Unknown"
	}

	var membership models.Membership
	err = json.NewDecoder(req.Body).Decode(&membership)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
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
	membership.EmployeeID = auth.EmployeeID
	membership.PlayerID = player.Id
	membership.PlayTime = (membership.Amount / 10) * 3600
	membership.ActiveDate = time.Now().Format(time.RFC3339)
	membership.DeactiveDate = time.Now().AddDate(20, 0, 0).Format(time.RFC3339)
	membership.Active = true

	err = membership.CreateMembership(db, transtype)
	if err != nil {
		error := models.RespError{Error: "Failed to create membership for player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(membership)
	return
}

/*
GetMembership - get the membership infomration of a player by player id. returns 400, 401, 404 for errors
*/
func (c *Controller) GetMembership(w http.ResponseWriter, req *http.Request) {
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
	membership := models.Membership{PlayerID: player.Id}

	err = membership.GetMembership(db)
	if err != nil {
		error := models.RespError{Error: "Failed to find membership for player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(membership)
	return
}

/*
AddPlayTime - add play time to an existing membership for a player. returns 400, 401, 404 for errors
*/
func (c *Controller) AddPlayTime(w http.ResponseWriter, req *http.Request) {
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

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id variable is required"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}
	amount, err := strconv.Atoi(req.URL.Query().Get("amount"))
	if err != nil {
		error := models.RespError{Error: "Amount variable is required"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	transtype := req.URL.Query().Get("type")
	if transtype == "" {
		transtype = "Unknown"
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
	membership := models.Membership{PlayerID: player.Id, EmployeeID: auth.EmployeeID, Amount: amount,
		PlayTime: amount * 360}

	err = membership.AddPlayTime(db, transtype)
	if err != nil {
		error := models.RespError{Error: "Failed to add play time to membership"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	err = membership.GetMembership(db)
	if err != nil {
		error := models.RespError{Error: "Failed to refetch membership"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(membership)
	return
}
