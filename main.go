package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
)

var quotes = []string{
	"The journey of a thousand miles begins with a single step.",
	"Not all those who wander are lost.",
	"This too shall pass.",
	"Act as if what you do makes a difference. It does.",
	"It always seems impossible until it’s done.",
}

func main() {
	// Create a default Gin router
	router := gin.Default()

	// GET /quotes - return all Monty Python quotes
	router.GET("/quotes", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"quotes": quotes})
	})

	// GET /quote - return a random Monty Python quote
	router.GET("/quote", func(c *gin.Context) {
		index := rand.Intn(len(quotes))
		c.JSON(http.StatusOK, gin.H{"quote": quotes[index]})
	})

	// POST /quote - add a new Monty Python quote to the list
	router.POST("/quote", func(c *gin.Context) {
		var newQuote struct {
			Quote string `json:"quote" binding:"required"`
		}
		if err := c.ShouldBindJSON(&newQuote); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		quotes = append(quotes, newQuote.Quote)
		c.JSON(http.StatusOK, gin.H{"message": "Quote added successfully!"})
	})

	// Start the server on port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	hostPort := fmt.Sprintf("0.0.0.0:%s", port)
	router.Run(hostPort)
}
