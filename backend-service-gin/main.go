package main

import (
    "context"
    "fmt"
    "log"

	"net/http"

    "github.com/gin-gonic/gin"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type fact struct {
    Text string
    Number  int
    Found bool
	Type string
}

// albums slice to seed record album data.
var facts = []fact{
    {Text: "83 is the atomic number of bismuth.", Number: 83, Found: true, type: "trivia"},
    {Text: "102 is the atomic number of nobelium, an actinide.", Number: 102, Found: true, type: "trivia"},
}

// getFacts responds with the list of facts as JSON.
func getFacts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, facts)
}

// getFactByNumber locates the fact whose Number value matches the number
// parameter sent by the client, then returns that fact as a response.
func getAlbumByID(c *gin.Context) {
	number := c.Param("number")

    // Loop through the list of facts, looking for
    // an album whose number value matches the parameter.
    for _, f := range facts {
        if f.Number == number {
            c.IndentedJSON(http.StatusOK, f)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "fact not found"})
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
    router.GET("/facts/:number", getFactByNumber)

    router.Run("localhost:8080")
}