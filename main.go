package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	end "mpolitakis.LinkApi/Endpoints"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profile", end.GetAllProfiles).Methods("GET")
	router.HandleFunc("/profile/{id}", end.GetProfileById).Methods("GET")
	router.HandleFunc("/profile", end.PostProfile).Methods("POST")
	router.HandleFunc("/photo", end.PostPhoto).Methods("POST")
	router.HandleFunc("/details/{id}", end.PostDetails).Methods("POST")
	http.ListenAndServe(":8080", router)
}
