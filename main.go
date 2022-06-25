package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"rest/database"
	"rest/routes"
)

func main() {
	router := initialSetupAndGetRouter()

	log.Println("Server started at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}

func initialSetupAndGetRouter() *mux.Router {
	err := gotenv.Load()

	if err != nil {
		log.Fatal("Could not load env vars, exiting.")
	}

	database.CreateTablesAndPrePopulate()
	router := routes.GetRouter()

	return router
}
