package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Account defines an account that is registered to use our service
type Account struct {
	ID    int64
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var accounts []Account
	db.Find(&accounts)
	err := json.NewEncoder(w).Encode(accounts)
	if err != nil {
		panic(err)
	}
}

func newUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	db.Create(&account)

	json, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var account Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		panic(err)
	}

	id := account.ID
	name := account.Name
	email := account.Email

	db.First(&account, id)

	account.Name = name
	account.Email = email
	db.Save(&account)

	fmt.Fprintf(w, "Successfully Updated User")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var account Account
	db.Where("id = ?", id).Find(&account).Delete(&account)

	fmt.Fprintf(w, "Successfully Deleted User")
}
