package main

import (
	"log"
	"net/http"

	"github.com/KrMrityunjay/GO-BOOKSTORE/pkg/routes"
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))

}


/* 
*/