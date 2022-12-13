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

	json.NewEncoder(w).Encode(resp)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)

	id, ok := params["id"]

	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("int convert error %s", err)
		json.NewEncoder(w).Encode(models.DetailResp{Status: 0, Msg: "internal server error"})
		return
	}
	var resp models.DetailResp

	if ok {
		resp = models.DetailResp{Status: 1, Msg: id}

		customer, err := customerDb.GetCustomer(intId)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Printf("GetCustomer error %s", err)
			json.NewEncoder(w).Encode(models.DetailResp{Status: 0, Msg: "customer not found"})
			return
		}

		json.NewEncoder(w).Encode(customer)
		return

	} else {
		w.WriteHeader(http.StatusBadRequest)
		resp = models.DetailResp{Status: 0, Msg: "customer id not provided"}
		return
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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("unmarshal error %s", err)
		json.NewEncoder(w).Encode(models.DetailResp{Status: 0, Msg: "internal server error"})
		return
	}
	customerDb.AddCustomer(newCustomer)

	resp := models.DetailResp{Status: 1, Msg: "Customer Created"}

	json.NewEncoder(w).Encode(resp)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newCustomer models.Customer
	err := json.Unmarshal(reqBody, &newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("unmarshal error %s", err)
		json.NewEncoder(w).Encode(models.DetailResp{Status: 0, Msg: "internal server error"})
		return
	}
	_, err = customerDb.UpdateCustomer(newCustomer)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Printf("UpdateCustomer error %s", err)
		json.NewEncoder(w).Encode(models.DetailResp{Status: 0, Msg: "customer not found"})
		return
	}

	resp := models.DetailResp{Status: 1, Msg: "Customer Updated"}

	json.NewEncoder(w).Encode(resp)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)

	id, ok := params["id"]

	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("int convert error %s", err)
		json.NewEncoder(w).Encode(models.DetailResp{Status: 0, Msg: "internal server error"})
		return
	}
	var resp models.DetailResp

	if ok {
		resp = models.DetailResp{Status: 1, Msg: id}

		_, err = customerDb.DeleteCustomer(intId)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Printf("DeleteCustomer error %s", err)
			json.NewEncoder(w).Encode(models.DetailResp{Status: 0, Msg: "customer not found"})
			return
		}

		json.NewEncoder(w).Encode(models.DetailResp{Status: 1, Msg: "Customer deleted"})
		return

	} else {
		w.WriteHeader(http.StatusNotFound)
		resp = models.DetailResp{Status: 0, Msg: "customer id not provided"}
		json.NewEncoder(w).Encode(models.DetailResp{Status: 0, Msg: "customer id not provided"})
		return
	}

	json.NewEncoder(w).Encode(resp)
}
