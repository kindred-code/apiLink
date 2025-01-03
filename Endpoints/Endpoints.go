package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func GetDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileId := params["profileId"]
	var details details.Details // Assuming you have a struct called details.Details
	conn := db.Connections()

	// Save the details to the database or perform any necessary operations with the data
	err := conn.QueryRowContext(context.Background(), (fmt.Sprintf("Select * from details where profile.id = %s;", profileId))).Scan(&details.ProfileId, &details.Gender, &details.Bio, &details.Location, &details.LastActiveTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No user: %v\n", err)
		os.Exit(1)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(details)
}
func GetProfileById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	profileId := params["profileId"]
	conn := db.Connections()
	var user us.Profile
	err := conn.QueryRow(fmt.Sprintf("Select * from profile where profile.id = %s;", profileId)).Scan(&user.Id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
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

func PostDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	s := params["profileId"]
	profileId, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	var u = new(details.Details)
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	conn := db.Connections()

	_, err = conn.ExecContext(context.Background(), db.BuildSqlDetails(u, profileId))
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
