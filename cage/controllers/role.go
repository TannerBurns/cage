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
AddRole - add a role for an employee by id. returns 400, 401, 404 for errors
*/
func (c *Controller) AddRole(w http.ResponseWriter, req *http.Request) {
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

	var role models.Role
	err = json.NewDecoder(req.Body).Decode(&role)
	if err != nil {
		error := models.RespError{Error: "Failed to parse request. Please make sure request is valid format"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}

	auth := models.Authentication{Decoded: context.Get(req, "decoded")}
	var ok bool
	if role.RoleID == 1 {
		ok, err = auth.Authorize(db, 1)
	} else {
		ok, err = auth.Authorize(db, 2)
	}
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

	employee := models.Employee{Id: id}
	err = employee.GetEmployee(db)
	if err != nil {
		error := models.RespError{Error: "Failed to find employee"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	role.EmployeeID = employee.Id

	err = role.AddRole(db)
	if err != nil {
		error := models.RespError{Error: "Failed to create role for employee"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(role)
	return
}

/*
GetRoles - get all roles for an employee by id. returns 400, 401, 404 for errors
*/
func (c *Controller) GetRoles(w http.ResponseWriter, req *http.Request) {
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

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		error := models.RespError{Error: "Id is required in route"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 400)
		c.Logger.Logging(req, 400)
		return
	}

	employee := models.Employee{Id: id}
	err = employee.GetEmployee(db)
	if err != nil {
		error := models.RespError{Error: "Failed to get employee"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	roles := models.Roles{EmployeeID: employee.Id}

	err = roles.GetRoles(db)
	if err != nil {
		error := models.RespError{Error: "Faied to get roles for employee"}
		resp, _ := json.Marshal(error)
		http.Error(w, string(resp), 404)
		c.Logger.Logging(req, 404)
		return
	}
	w.WriteHeader(http.StatusOK)
	c.Logger.Logging(req, 200)
	json.NewEncoder(w).Encode(roles)
	return
}
