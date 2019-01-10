package controllers

import (
	"encoding/json"
	"net/http"
)

/*
Status - strucutre or server status
*/
type Status struct {
	Status  string
	Name    string
	Version string
}

/*
Index - return if server is live
*/
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Status{Status: "OK", Name: "The House", Version: "1.0.0"})
	return
}
