package main

import (
	"crmBackend/pkg/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Instantiate a new router
	router := mux.NewRouter()
	// Register handler functions to the same path -- but with different methods
	// E.g., only a GET request to /dictionary can invoke the "getDictionary" handler function
	//router.HandleFunc("/", controllers.GetRoot).Methods("GET")
	router.HandleFunc("/customers", controllers.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", controllers.GetCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Server is starting on port 3000...")
	// Pass the customer router into ListenAndServe
	http.ListenAndServe(":3000", router)
}
