package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	db "mpolitakis.LinkApi/Connections"
	ph "mpolitakis.LinkApi/Data/Photo"
	us "mpolitakis.LinkApi/Data/Profile"
)

func main() {
	router := gin.Default()
	router.GET("/profile", GetProfile)
	router.POST("/profile", PostProfile)
	router.POST("/photo", PostPhoto)
	router.Run("localhost:8080")
}

// GetProfile returns all users from the database in json format.
func GetProfile(c *gin.Context) {

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

	c.IndentedJSON(http.StatusCreated, users)
	defer conn.Close()

}

// PostProfile adds a new user to the database, given the json body of the POST request.
func PostProfile(c *gin.Context) {

	var u = new(us.Profile)
	if err := c.BindJSON(&u); err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	conn := db.Connections()

	_, err := conn.ExecContext(context.Background(), db.BuildSql(u))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	c.IndentedJSON(http.StatusCreated, u)
	defer conn.Close()

}

// PostPhoto adds a new photo to the database, given the json body of the POST request.
func PostPhoto(c *gin.Context) {
	var photo = new(ph.Photo)
	if err := c.BindJSON(&photo); err != nil {
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
}
