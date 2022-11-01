package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zgeorgievBG/go_bookstore/pkg/router"
)

func main() {
	r := mux.NewRouter()
	router.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
	fmt.Printf("Server Started")
}