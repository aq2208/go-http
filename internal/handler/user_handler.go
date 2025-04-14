package handler

import (
	"encoding/json"
	"go-database/internal/repository"
	"log"
	"net/http"
	"strconv"
)

func GetListUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		users, err := repository.GetAllUsers()
		if err != nil {
			log.Fatal(err)
			http.NotFound(w, r)
		}

		json.NewEncoder(w).Encode(users)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetDetailUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userId, err := strconv.Atoi(r.URL.Path[len("/users/"):])
		users, err := repository.GetUserById(userId)
		if err != nil {
			log.Fatal(err)
			http.NotFound(w, r)
		}

		json.NewEncoder(w).Encode(users)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	
}