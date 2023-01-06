package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vick1208/jwt-go/controllers/authcontroller"
	"github.com/vick1208/jwt-go/controllers/productcontroller"
	"github.com/vick1208/jwt-go/middlewares"
	"github.com/vick1208/jwt-go/models"
)

func main() {

	models.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods(http.MethodPost)
	r.HandleFunc("/register", authcontroller.Register).Methods(http.MethodPost)
	r.HandleFunc("/logout", authcontroller.Logout).Methods(http.MethodGet)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods(http.MethodGet)
	api.Use(middlewares.JWTMiddle)

	log.Fatal(http.ListenAndServe(":4000", r))
}
