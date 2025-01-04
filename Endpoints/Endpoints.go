package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"context"
	"fmt"

	details "mpolitakis.LinkApi/Data/Details"

	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	db "mpolitakis.LinkApi/Connections"
	ph "mpolitakis.LinkApi/Data/Photo"
	us "mpolitakis.LinkApi/Data/Profile"
)

func GetDetails(c *gin.Context) {
	profileId := c.Params.ByName("profileId")

	var details details.Details // Assuming you have a struct called details.Details
	conn := db.Connections()
	defer conn.Close()
	// Save the details to the database or perform any necessary operations with the data
	err := conn.QueryRowContext(context.Background(), (fmt.Sprintf("Select * from details where profile.id = %s;", profileId))).Scan(&details.ProfileId, &details.Gender, &details.Bio, &details.Location, &details.LastActiveTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No user: %v\n", err)
		os.Exit(1)
	}

	c.Writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(c.Writer).Encode(details)
}
func GetProfileById(c *gin.Context) {
	profileId := c.Params.ByName("profileId")
	conn := db.Connections()
	var user us.Profile
	err := conn.QueryRow(fmt.Sprintf("Select * from profile where profile.id = %s;", profileId)).Scan(&user.Id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	c.Writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(c.Writer).Encode(user)

}

// GetProfile returns all users from the database in json format.
func GetAllProfiles(c *gin.Context) {

	conn := db.Connections()
	var users = []us.Profile{}
	rows, err := conn.Query("Select * from profile;")
	defer conn.Close()
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

	c.Writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(c.Writer).Encode(users)
}

// PostProfile adds a new user to the database, given the json body of the POST request.
func PostProfile(c *gin.Context) {

	var u = new(us.Profile)
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
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
	c.Writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(c.Writer).Encode(u)
}

func PostDetails(c *gin.Context) {
	profileId, err := strconv.Atoi(c.Params.ByName("profileId"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	var u = new(details.Details)
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	conn := db.Connections()
	defer conn.Close()
	_, err = conn.ExecContext(context.Background(), db.BuildSqlDetails(u, profileId))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	c.Writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(c.Writer).Encode(u)
}

// PostPhoto adds a new photo to the database, given the json body of the POST request.
func PostPhoto(c *gin.Context) {

	var photo = new(ph.Photo)
	if err := json.NewDecoder(c.Request.Body).Decode(&photo); err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}
	conn := db.Connections()
	defer conn.Close()
	_, err := conn.ExecContext(context.Background(), db.BuildSqlPhoto(photo))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	c.Writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(c.Writer).Encode(photo)
}

func GetPhoto(c *gin.Context) {
	profileId := c.Params.ByName("profileId")
	var photos ph.Photo
	conn := db.Connections()
	defer conn.Close()
	// Save the details to the database or perform any necessary operations with the data
	err := conn.QueryRowContext(context.Background(), (fmt.Sprintf("Select * from photo where photo.profileId = %s;", profileId))).Scan(&photos.Id, &photos.File, &photos.ProfileId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No photos: %v\n", err)
		os.Exit(1)
	}

	c.Writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(c.Writer).Encode(photos)
}
