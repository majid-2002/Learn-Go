package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/majid-2002/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
    r := mux.NewRouter()
    routes.RegisterBookStoreRoutes(r)
    log.Println("Starting server on port :9010...")
    log.Fatal(http.ListenAndServe(":9010", r))
}
