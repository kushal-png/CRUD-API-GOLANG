package router

import (
	"CRUD_API/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/movies", controller.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controller.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.DeleteMovie).Methods("DELETE")
    
	return r
}
