package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/middlewares"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/controllers/authcontroller"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/controllers/photocontroller"
	"github.com/Rezapahlevi3108/task-5-pbi-btpns-reza-pahlevi-kurniawan/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/photos", photocontroller.CreatePhoto).Methods("POST")
	api.HandleFunc("/photos", photocontroller.GetPhotos).Methods("GET")
	api.HandleFunc("/photos/{photoId}", photocontroller.UpdatePhoto).Methods("PUT")
	api.HandleFunc("/photos/{photoId}", photocontroller.DeletePhoto).Methods("DELETE")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}