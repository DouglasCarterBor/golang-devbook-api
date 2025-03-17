package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
			log.Fatal(err)
	}

	var user models.User 
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewRepositoryUsers(db)
	usuarioID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(fmt.Appendf(nil, "ID inserted: %d", usuarioID))

}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for users!"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user!"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user!"))
}