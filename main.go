package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	routing "mpolitakis.LinkApi/Routing"
)

func main() {
	router := mux.NewRouter()
	routing.Routing(router)
	http.ListenAndServe(":8080", router)
}
