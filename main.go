package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type image struct {
	RequestedTopic string   `json:"requestedTopic"`
	Images         []string `json:"Images"`
	Number         int      `json:"Number"`
}

var images = []image{
	{
		RequestedTopic: "the+room+memes",
		Images: []string{
			"https://preview.redd.it/0ndypduzlf641.jpg?width=640&crop=smart&auto=webp&s=5cf49a2deb07cfa5bf628e1dfb074b0452bb8041",
			"https://preview.redd.it/q6ez8ub45i931.jpg?width=640&crop=smart&auto=webp&s=4b4859c1cc752ac3310496852b8fd19a77aa7fc7",
		},
		Number: 2,
	},
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/topics", getImageTopics)
	return router
}

func main() {
	router := setupRouter()
	port := os.Getenv("PORT")

	if port != "" {
		router.Run("0.0.0.0:" + port)
	} else {
		router.Run("0.0.0.0:5000")
	}
}

func getImageTopics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, images)
}
