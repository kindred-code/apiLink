package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	routing "mpolitakis.LinkApi/Routing"
)

func main() {
	router := gin.Default()
	routing.Routing(router)
	fmt.Println("Starting server")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
