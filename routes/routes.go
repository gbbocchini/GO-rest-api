package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest/controllers"
	"rest/middlewares"
)

func GetRouter() *mux.Router {
	booksController := controllers.BookController{}
	userController := controllers.UserController{}
	router := mux.NewRouter()

	router.HandleFunc("/auth/login", userController.LoginUser()).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", userController.SignupUser()).Methods(http.MethodPost)

	router.HandleFunc("/books", middlewares.CheckAuth(booksController.GetBooks())).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", middlewares.CheckAuth(booksController.GetBook())).Methods(http.MethodGet)
	router.HandleFunc("/books", middlewares.CheckAuth(booksController.AddBook())).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", middlewares.CheckAuth(booksController.UpdateBook())).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", middlewares.CheckAuth(booksController.RemoveBook())).Methods(http.MethodDelete)

	return router
}
