package main

import (
	"log"
	"net/http"

	// "os"

	"go-database/config"
	"go-database/internal/database"
	"go-database/internal/handler"
)

func main() {
	// load env
	config.LoadEnv()

	// connect db
	database.ConnectDb()

	// request handler
	http.HandleFunc("/users", handler.GetListUserHandler)
	http.HandleFunc("/users/", handler.GetDetailUserHandler)

	log.Println("Server starting on " + config.GetEnv("APP_PORT"))
	log.Fatal(http.ListenAndServe(config.GetEnv("APP_PORT"), nil))
}