package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"onecv_tech/controller"
)

func main() {
	router := mux.NewRouter()

	// API exposed endpoints
	router.HandleFunc("/api/register", controller.RegisterStudentsToTeacher).Methods("POST")
	router.HandleFunc("/api/commonstudents", controller.GetCommonStudentsOfTeachers).Methods("GET")
	router.HandleFunc("/api/suspend", controller.SuspendStudent).Methods("POST")
	router.HandleFunc("/api/retrievefornotifications", controller.RetrieveRecipientList).Methods("POST")
	
	http.Handle("/", router)
	fmt.Println("Connected to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}