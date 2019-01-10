package main

import (
	"net/http"
	"os"

	"./logger"
	"./routes"

	"github.com/gorilla/handlers"
)

func main() {
	logs := &logger.Logger{}
	logs.InitLogging("House - API", os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr)

	//host := "localhost"
	port := "5000"

	router := routes.NewRouter() // create routes

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	logs.Fatal.Println(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
