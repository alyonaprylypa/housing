package main

import (
	"Project/controllers"
	"Project/driver"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB


func main() {
	db=driver.ConnectDB()
	defer db.Close()
	router := mux.NewRouter()
	controller:=controllers.Controller{}
	router.HandleFunc("/flats", controller.GetFlats(db)).Methods("GET")
	router.HandleFunc("/flats/{id}", controller.GetFlat(db)).Methods("GET")
	router.HandleFunc("/login",controller.LoginPageHandler(db)).Methods("GET")
	router.HandleFunc("/login",controller.LoginHandler(db)).Methods("POST")
	router.HandleFunc("/register",controller.RegisterPageHandler(db)).Methods("GET")
	router.HandleFunc("/register",controller.RegisterHandler(db)).Methods("POST")
	//router.HandleFunc("/index/{name}",controller.GetUserPage(db)).Methods("GET")


	log.Fatal(http.ListenAndServe(":8000", router))
}
