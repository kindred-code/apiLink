package main

import (

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"context"
	"fmt"
	"net/http"
	"os"
	_ "github.com/lib/pq"



	us "gitlab.com/mpolitakis/linkapi/User"
	db "gitlab.com/mpolitakis/linkapi/Connections"

)

func main() {
	router := gin.Default()
	router.GET("/user",GetUser)
	router.POST("/user", PostUser)

	router.Run("localhost:8080")
}



func GetUser(c *gin.Context) {

	conn := db.Connections()

	// albums slice to seed record album data.
	var users = []us.User{}
	rows ,err := conn.Query("Select * from users;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next(){
		var user us.User
		err = rows.Scan(&user.ID,&user.Email, &user.Username, &user.Password);
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}
		users = append(users, user)
	}

	c.IndentedJSON(http.StatusCreated, users)
	defer conn.Close()
	return


}

func PostUser(c *gin.Context){

	var u = new(us.User)
	if err := c.BindJSON(&u); err != nil {
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}

	conn := db.Connections()

	_, err := conn.ExecContext(context.Background(),db.BuildSql(u))
	if err != nil{
		fmt.Fprintf(os.Stderr, "Wrong data format: %v\n", err)
		os.Exit(1)
	}


	c.IndentedJSON(http.StatusCreated, u)
	defer conn.Close()
	return

}

