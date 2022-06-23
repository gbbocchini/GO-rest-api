package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"rest/controllers"
)

func main() {
	gotenv.Load()
	controller := controllers.Controller{}
	router := mux.NewRouter()
	router.HandleFunc("/books", controller.GetBooks()).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook()).Methods("GET")
	router.HandleFunc("/books", controller.AddBook()).Methods("POST")
	router.HandleFunc("/books/{id}", controller.UpdateBook()).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook()).Methods("DELETE")
	log.Println("Starting server at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
