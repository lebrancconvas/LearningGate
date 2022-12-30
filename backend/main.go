package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func getIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func main() {
	// Start the server 
	router := gin.Default();
	router.GET("/", getIndex);
	router.Run(":8000"); 
}