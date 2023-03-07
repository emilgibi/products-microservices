package main

import (
	"os"

	"net/http"

	"github.com/emilgibi/products-microservices/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	var handlerObj handlers.Handler

	godotenv.Load()

	HOST := os.Getenv("DATABASE_HOST")
	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASS")
	NAME := os.Getenv("DATABASE_NAME")
	PORT := os.Getenv("DATABASE_PORT")

	handlerObj.Connect(HOST, USER, PASS, NAME, PORT)

	dbinstance, _ := handlerObj.DB.DB()
	defer dbinstance.Close()

	router := mux.NewRouter()

	router.HandleFunc("/products", handlerObj.GetProduct).Methods("GET")
	router.HandleFunc("/product/{product_id}", handlerObj.Getuser).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8081", router)
}
