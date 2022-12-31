package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type User struct {
	ID string `json: "id"` 
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

func createUser(c *gin.Context) {
	var newUser User; 

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}

	users = append(users, newUser);
	c.IndentedJSON(http.StatusCreated, newUser); 
}

func main() {
	// Start the server 
	router := gin.Default();
	router.Use(cors.Default()); 
	router.GET("/", getIndex);
	router.GET("/api/v1/users", getUser); 
	router.POST("/api/v1/users", createUser); 
	router.Run(":8000"); 
}