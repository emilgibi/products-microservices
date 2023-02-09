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

	HOST := os.Getenv("DB_HOST")
	USER := os.Getenv("USER_NAME")
	PASS := os.Getenv("PASS")

	handlerObj.Connect(HOST, USER, PASS, "postgres", "5432")

	dbinstance, _ := handlerObj.DB.DB()
	defer dbinstance.Close()

	router := mux.NewRouter()

	router.HandleFunc("/products", handlerObj.GetProduct).Methods("GET")
	router.HandleFunc("/product/{product_id}", handlerObj.Getuser).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8081", router)
}
