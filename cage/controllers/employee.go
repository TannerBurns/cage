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
CreateEmployee - creates an employee, returns 400, 401, 404 for errors
*/
func (c *Controller) CreateEmployee(w http.ResponseWriter, req *http.Request) {
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
	ok, err := auth.Authorize(db, 2)
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

	var employee models.Employee
	err = json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}

	err = employee.CreateEmployee(db)
	if err != nil {
		error := models.RespError{Error: "Failed to create a new employee"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
	return
}

/*
GetEmployee - read an employee by id, returns 400, 401, 404 for errors
*/
func (c *Controller) GetEmployee(w http.ResponseWriter, req *http.Request) {
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
	ok, err := auth.Authorize(db, 2)
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

	employee := models.Employee{Id: id}
	err = employee.GetEmployee(db)
	if err != nil {
		error := models.RespError{Error: "Failed to find employee"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
	return
}

/*
UpdateEmployee - update an employee by id, returns 400, 401, 404 for errors
*/
func (c *Controller) UpdateEmployee(w http.ResponseWriter, req *http.Request) {
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
	ok, err := auth.Authorize(db, 2)
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

	var employee models.Employee
	err = json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	employee.Id = id
	err = employee.UpdateEmployee(db)
	if err != nil {
		error := models.RespError{Error: "Failed to update employee"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)

	return
}
