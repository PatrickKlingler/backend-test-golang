package main

import (
	"context"
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type fact struct {
	Text   string
	Number int
	Found  bool
	Type   string
}

// albums slice to seed record album data.
var facts = []fact{
	{Text: "83 is the atomic number of bismuth.", Number: 83, Found: true, Type: "trivia"},
	{Text: "102 is the atomic number of nobelium, an actinide.", Number: 102, Found: true, Type: "trivia"},
}

// getFacts responds with the list of facts as JSON.
func getFacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, facts)
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	router := gin.Default()
	router.GET("/facts", getFacts)

	router.Run("localhost:8080")
}
