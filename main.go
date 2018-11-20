package main

// Run this program
// go run main.go database.go config.go

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user", newUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	dbConn, err := gorm.Open("mysql", dataSourceName+dbName)
	if err != nil {
		log.Fatal(err)
	}
	db = dbConn

	fmt.Println("Database Connected Successfully")

	// Handle Subsequent requests
	handleRequests()
}
