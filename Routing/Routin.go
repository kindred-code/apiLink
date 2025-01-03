package routing

import (
	"github.com/gorilla/mux"
	end "mpolitakis.LinkApi/Endpoints"
)

func Routing(router *mux.Router) {

	router.HandleFunc("/details/{profileId}", end.GetDetails).Methods("GET")
	router.HandleFunc("/details/{profileId}", end.PostDetails).Methods("POST")
	router.HandleFunc("/profile/{profileId}", end.GetProfileById).Methods("GET")
	router.HandleFunc("/profile", end.GetAllProfiles).Methods("GET")
	router.HandleFunc("/profile", end.PostProfile).Methods("POST")
	router.HandleFunc("/photo/", end.PostPhoto).Methods("POST")
	router.HandleFunc("/photo/{profileId}", end.GetPhoto).Methods("GET")

}
