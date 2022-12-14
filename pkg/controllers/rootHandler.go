package controllers

import (
	"fmt"
	"net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	fmt.Fprintf(w, "<h1>Welcome to the customer API. Active routes are /customers.</h1>")
	fmt.Fprintf(w, "<p>The customer object needed is {Id: int, Name: string, Role: string, active: bool} </p>")
	fmt.Fprintf(w, "<p>Create: /customers Get Single: /customers/:id Get All: /customers Update Single: /customers/:id</p>")
	fmt.Fprintf(w, "<p>Delete Single: /customers/:id</p>")
}
