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

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate Content-Type in the response header
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newCustomer models.Customer
	err := json.Unmarshal(reqBody, &newCustomer)
	if err != nil {
		fmt.Printf("unmarshal error %s", err)
	}
	customerDb.AddCustomer(newCustomer)

	resp := models.DetailResp{Status: 1, Msg: "Customer Created"}

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
