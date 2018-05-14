package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/api/user/{Name}", getUser).Methods("GET")
	myRouter.HandleFunc("/api/user/{Name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/api/user/", updateUser).Methods("PUT")
	myRouter.HandleFunc("/api/user/", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

var db *gorm.DB
var err error

func main() {
	fmt.Println("Go ORM Tutorial")

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}) // Handle Subsequent requests
	handleRequests()
}
