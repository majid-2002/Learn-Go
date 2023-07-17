package routes

import (
	"github.com/gorilla/mux"
	"github.com/majid-2002/go-bookstore/pkg/controllers"
	"net/http"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/api/books", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{bookId}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{bookId}", controllers.DeleteBook).Methods(http.MethodDelete)
}
 