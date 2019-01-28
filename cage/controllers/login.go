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
Creds - temporary holder for username, password
*/
type Creds struct {
	Username string
	Password string
}

/*
CreateLogin - creates an employee login, returns 400, 401, 404 for errors
*/
func (c *Controller) CreateLogin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

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
	ok, err := auth.Authorize(db, 2)
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

	var creds Creds
	err = json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}

	login := models.Login{EmployeeID: id, Username: creds.Username}
	err = login.CreateLogin(db, creds.Password)
	if err != nil {
		error := models.RespError{Error: "Failed to create a new login for employee"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(login)
	return
}
