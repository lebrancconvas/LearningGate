package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: "username"`
	Displayname string `json: "displayname"`
	Password string `json: "password"`
	Email string `json: "email"` 
}

var users []User; 

func getIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users); 
}

func main() {
	// Start the server 
	router := gin.Default();
	router.GET("/", getIndex);
	router.GET("/api/v1/users", getUser); 
	router.Run(":8000"); 
}