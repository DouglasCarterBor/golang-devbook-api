package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new user!"))
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