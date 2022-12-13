package controllers

import (
	"crmBackend/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

var customerDb = models.NewCustomerDb()

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := customerDb.GetDb()

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Printf("error in GetAllCustomer: %s", err)
		return
	}
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)

	id, ok := params["id"]

	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("int convert error %s", err)
	}
	var resp models.DetailResp

	if ok {
		resp = models.DetailResp{Status: 1, Msg: id}

		customer := customerDb.GetCustomer(intId)

		json.NewEncoder(w).Encode(customer)
		return

	} else {
		resp = models.DetailResp{Status: 0, Msg: "customer id not provided"}
	}

	json.NewEncoder(w).Encode(resp)
}

// This "create" handler function's goal is to take user input via a JSON request body (i.e., payload)
// As a response to calls made to its associated route, this function returns the new state of the recently-updated dictionary
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate Content-Type in the response header
	w.Header().Set("Content-Type", "application/json")

	// Create (but not yet assign values to) for the new entry
	var newEntry map[string]string

	// Read the HTTP request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Encode the request body into a Golang value so that we can work with the data
	json.Unmarshal(reqBody, &newEntry)

	// Iterate over the "newEntry" map
	//for k, v := range newEntry {
	//	if _, ok := dictionary[k]; ok {
	//		// If the key already exists (i.e., "dictionary" work already exists) in the map, update the HTTP status with a "407 Conflict" message
	//		// In such a case, we do not update the original "dictionay" map at all
	//		w.WriteHeader(http.StatusConflict)
	//	} else {
	//		// If the key doesn't exist, add it to the "dictionary" map and return a "201 Created" in the header
	//		dictionary[k] = v
	//		w.WriteHeader(http.StatusCreated)
	//	}
	//}

	// Regardless of successful resource creation or not, return the current state of the "dictionary" map
	resp := models.DetailResp{Status: 1, Msg: "CreateCustomer"}

	json.NewEncoder(w).Encode(resp)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := models.DetailResp{Status: 1, Msg: "UpdateCustomer"}

	json.NewEncoder(w).Encode(resp)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := models.DetailResp{Status: 1, Msg: "DeleteCustomer"}

	json.NewEncoder(w).Encode(resp)
}
