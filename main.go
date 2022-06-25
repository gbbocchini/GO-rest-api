package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"rest/controllers"
	"rest/database"
	"rest/middlewares"
)

func main() {
	router, userController, booksController := start()

	router.HandleFunc("/auth/login", userController.LoginUser()).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", userController.SignupUser()).Methods(http.MethodPost)

	router.HandleFunc("/books", middlewares.CheckAuth(booksController.GetBooks())).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", middlewares.CheckAuth(booksController.GetBook())).Methods(http.MethodGet)
	router.HandleFunc("/books", middlewares.CheckAuth(booksController.AddBook())).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", middlewares.CheckAuth(booksController.UpdateBook())).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", middlewares.CheckAuth(booksController.RemoveBook())).Methods(http.MethodDelete)

	log.Println("Server started at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}

func start() (*mux.Router, controllers.UserController, controllers.BookController) {
	err := gotenv.Load()

	if err != nil {
		panic("Could not load env vars, exiting.")
	}

	database.CreateTablesAndPrePopulate()
	router := mux.NewRouter()
	booksController := controllers.BookController{}
	userController := controllers.UserController{}

	return router, userController, booksController
}
