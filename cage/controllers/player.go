package controllers

import (
	"encoding/json"
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
	postclient := models.PostgresConnection{}
	db, err := postclient.Connect()
	if err != nil {
		error := models.RespError{Error: "Failed to connect, cannot reach database"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}
	defer db.Close()

	auth := models.Authentication{Decoded: context.Get(req, "decoded")}
	ok, err := auth.Authorize(db, 3)
	if err != nil {
		error := models.RespError{Error: "Failed to authorize, error during authorization"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}
	if !ok {
		error := models.RespError{Error: "Failed to authorize, error during authorization. Make sure you have permissions to use this route."}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 401)
		return
	}

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}

	player := models.Player{Id: id}

	err = player.GetPlayer(db)
	if err != nil {
		error := models.RespError{Error: "Failed to find player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	w.WriteHeader(http.StatusOK)
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
	postclient := models.PostgresConnection{}
	db, err := postclient.Connect()
	if err != nil {
		error := models.RespError{Error: "Failed to connect, cannot reach database"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}
	defer db.Close()

	auth := models.Authentication{Decoded: context.Get(req, "decoded")}
	ok, err := auth.Authorize(db, 3)
	if err != nil {
		error := models.RespError{Error: "Failed to authorize, error during authorization"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}
	if !ok {
		error := models.RespError{Error: "Failed to authorize, error during authorization. Make sure you have permissions to use this route."}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 401)
		return
	}

	var player models.Player
	err = json.NewDecoder(req.Body).Decode(&player)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}

	err = player.CreatePlayer(db)
	if err != nil {
		error := models.RespError{Error: "Failed to create a new player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	w.WriteHeader(http.StatusOK)
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
	postclient := models.PostgresConnection{}
	db, err := postclient.Connect()
	if err != nil {
		error := models.RespError{Error: "Failed to connect, cannot reach database"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}
	defer db.Close()

	auth := models.Authentication{Decoded: context.Get(req, "decoded")}
	ok, err := auth.Authorize(db, 3)
	if err != nil {
		error := models.RespError{Error: "Failed to authorize, error during authorization"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}
	if !ok {
		error := models.RespError{Error: "Failed to authorize, error during authorization. Make sure you have permissions to use this route."}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 401)
		return
	}

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}

	var player models.Player
	err = json.NewDecoder(req.Body).Decode(&player)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	player.Id = id
	err = player.UpdatePlayer(db)
	if err != nil {
		error := models.RespError{Error: "Failed to update player"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
	return
}

func (c *Controller) Search(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//connect to database
	postclient := models.PostgresConnection{}
	db, err := postclient.Connect()
	if err != nil {
		error := models.RespError{Error: "Failed to connect, cannot reach database"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}
	defer db.Close()

	auth := models.Authentication{Decoded: context.Get(req, "decoded")}
	ok, err := auth.Authorize(db, 3)
	if err != nil {
		error := models.RespError{Error: "Failed to authorize, error during authorization"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}
	if !ok {
		error := models.RespError{Error: "Failed to authorize, error during authorization. Make sure you have permissions to use this route."}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 401)
		return
	}

	first := req.URL.Query().Get("first")
	if first == "" {
		error := models.RespError{Error: "Failed to find 'first' variable"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}

	last := req.URL.Query().Get("last")
	if last == "" {
		error := models.RespError{Error: "Failed to find 'last' variable"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		return
	}

	players := models.Players{}
	err = players.Search(db, first, last)
	if err != nil {
		error := models.RespError{Error: "Failed to find any players matching the search content"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
	return
}
