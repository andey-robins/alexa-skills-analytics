package main

import (
	"github.com/andey-robins/alexa-skills-analytics/server/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}())

	router.GET("/users", api.GetUsers)
	router.POST("/newUser", api.PostUser)
	router.GET("/answers", api.GetAnswers)
	router.POST("/newAnswer", api.PostAnswer)

	router.Run("localhost:80")
}
