package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response structure for JSON responses
type Response struct {
	Message string `json:"message,omitempty"`
	Number  int    `json:"number,omitempty"`
	Input   string `json:"input,omitempty"`
}

func main() {
	r := gin.Default()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Welcome endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello, World!",
		})
	})

	// Random number endpoint
	r.GET("/random", func(c *gin.Context) {
		randomNum := rand.Intn(100) // Random number between 0 and 99
		response := Response{Number: randomNum}
		c.JSON(http.StatusOK, response)
	})

	// Start the server on port 8080
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
