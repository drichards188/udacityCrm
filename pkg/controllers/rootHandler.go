package controllers

import (
	"crmBackend/pkg/models"
	"encoding/json"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := models.DetailResp{Status: 1, Msg: "Welcome to the customer API. Active routes are /customers. The customer object needed is {Id: int, Name: string, Role: string, active: bool} Create: /customers Get Single: /customers/:id Get All: /customers Update Single: /customers/:id Delete Single: /customers/:id"}

	json.NewEncoder(w).Encode(resp)

}
