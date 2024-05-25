package main

import (
	"CRUD_API/router"
	"CRUD_API/store"
	"fmt"
	"log"
	"net/http"
)

//creating dummy db

func main() {
	fmt.Println("Golang CRUD API using Gorilla mux")

	r := router.Router()
	store.InitializeMovies()

	fmt.Println("Creting the server on post 4000")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal(err, "Unable to serve")
		return
	}
	fmt.Println("Seccessfully serving on port 4000")
}
