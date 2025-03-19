package controllers

import (
	"api/src/answers"
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
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User 
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusCreated, user)

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