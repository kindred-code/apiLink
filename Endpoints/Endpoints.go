package endpoints

import (
	"encoding/json"
	"net/http"

	"context"
	"fmt"

	details "mpolitakis.LinkApi/Data/Details"

	"os"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	db "mpolitakis.LinkApi/Connections"
	ph "mpolitakis.LinkApi/Data/Photo"
	us "mpolitakis.LinkApi/Data/Profile"
)

func PostDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileId := params["profileId"]
	var details = new(details.Details) // Assuming you have a struct called details.Details
	conn := db.Connections()
	err := json.NewDecoder(r.Body).Decode(&details)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Save the details to the database or perform any necessary operations with the data
	rows, err := conn.Query(fmt.Sprintf("Select * from details where Id = %s;", profileId))
	if err != nil {
		fmt.Fprintf(os.Stderr, "No user: %v\n", err)
		os.Exit(1)
	}
	err = rows.Scan(profileId, &details.Gender, &details.Bio, &details.Interests, &details.Location, &details.LastActiveTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
		os.Exit(1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(details)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(details)
}
func GetProfileById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileId := params["profileId"]
	conn := db.Connections()

	rows, err := conn.Query(fmt.Sprintf("Select * from profile where Id = %s;", profileId))

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	var user us.Profile
	err = rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// GetProfile returns all users from the database in json format.
func GetAllProfiles(w http.ResponseWriter, r *http.Request) {

	conn := db.Connections()
	var users = []us.Profile{}
	rows, err := conn.Query("Select * from profile;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var user us.Profile
		err = rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}
		users = append(users, user)
	}

	defer conn.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// PostProfile adds a new user to the database, given the json body of the POST request.
func PostProfile(w http.ResponseWriter, r *http.Request) {

	var u = new(us.Profile)
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	conn := db.Connections()

	_, err := conn.ExecContext(context.Background(), db.BuildSql(u))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// PostPhoto adds a new photo to the database, given the json body of the POST request.
func PostPhoto(w http.ResponseWriter, r *http.Request) {
	var photo = new(ph.Photo)
	if err := json.NewDecoder(r.Body).Decode(&photo); err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}
	conn := db.Connections()
	_, err := conn.ExecContext(context.Background(), db.BuildSqlPhoto(photo))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photo)
}
