package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.Handle("/api/users", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(allUsers))).Methods("GET")
	r.Handle("/api/user/{Name}", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(getUser))).Methods("GET")
	r.Handle("/api/user/{Name}", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(deleteUser))).Methods("DELETE")
	r.Handle("/api/user/", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(updateUser))).Methods("PUT")
	r.Handle("/api/user/", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(newUser))).Methods("POST")

	r.HandleFunc("/api/user/{Name}", deleteUser).Methods("DELETE")
	r.HandleFunc("/api/user/", updateUser).Methods("PUT")
	r.HandleFunc("/api/user/", newUser).Methods("POST")
	log.Println("Go server starting on port 8081")

	log.Fatal(http.ListenAndServe(":8081", r))
}

var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}) // Handle Subsequent requests
	handleRequests()
}
