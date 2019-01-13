package main

import (
	"fmt"
	"net/http"
	"os"

	"./models"
	"./routes"

	"github.com/gorilla/handlers"
)

func main() {
	logs := &models.Logger{}
	logs.InitLogging("House - API", os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr, os.Stdout)

	//host := "localhost"
	port := "5000"

	router, conlogger := routes.NewRouter() // create routes

	f, err := os.OpenFile("connections.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to initialize logger")
	}
	defer f.Close()
	conlogger.Log.SetOutput(f)

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	logs.Fatal.Println(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
