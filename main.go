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

type Account struct {
	gorm.Model

	Name  string
	Email string
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")

	db, err := gorm.Open("mysql", dataSourceName+dbName)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Account{})

	// Handle Subsequent requests
	handleRequests()
}
