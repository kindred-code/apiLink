package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	end "mpolitakis.LinkApi/Endpoints"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/details/{profileId}", end.GetDetails).Methods("GET")
	router.HandleFunc("/details/{profileId}", end.PostDetails).Methods("POST")
	router.HandleFunc("/profile/{profileId}", end.GetProfileById).Methods("GET")
	router.HandleFunc("/profile", end.GetAllProfiles).Methods("GET")
	router.HandleFunc("/profile", end.PostProfile).Methods("POST")
	router.HandleFunc("/photo/", end.PostPhoto).Methods("POST")
	http.ListenAndServe(":8080", router)
}
