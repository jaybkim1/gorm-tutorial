package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Account struct {
	ID    int64
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {

	var accounts []Account
	db.Find(&accounts)
	json.NewEncoder(w).Encode(accounts)
	fmt.Println(accounts)
}

func newUser(w http.ResponseWriter, r *http.Request) {

	var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	db.Create(&account)

	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	var account Account
	db.Where("name = ?", name).Find(&account)
	db.Delete(&account)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var account Account
	db.Where("name = ?", name).Find(&account)

	account.Email = email

	db.Save(&account)
	fmt.Fprintf(w, "Successfully Updated User")
}
