package main

import (
	"log"
	"net/http"

	"github.com/customerapi/my_api"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/customers", my_api.GetCustomers).Methods("GET")
	router.HandleFunc("/customers", my_api.PostCustomer).Methods("POST")
	router.HandleFunc("/customers", my_api.PutCustomer).Methods("PUT")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", my_api.DeleteCustomer).Methods("DELETE")

	log.Println("Server started on :80")
	log.Fatal(http.ListenAndServe(":8080", router))
}
