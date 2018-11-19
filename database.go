package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var accounts []Account
	db.Find(&accounts)
	fmt.Println("{}", accounts)

	json.NewEncoder(w).Encode(accounts)

}

func newUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(email)

	db.Create(&Account{Name: name, Email: email})
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
