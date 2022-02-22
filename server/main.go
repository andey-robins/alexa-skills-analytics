package main

import (
	"log"
	"net/http"
	"time"

	"github.com/andey-robins/alexa-skills-analytics/server/auth"
	"github.com/andey-robins/alexa-skills-analytics/server/types"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.GET("/sample-answers", getSampleAnswers)
	server.GET("/sample-answers/:id", getSampleAnswerById)
	server.POST("/answer", postAnswer)

	server.POST("/login", func(ctx *gin.Context) {
		var loginService auth.LoginService = auth.StaticLoginService()
		var jwtService auth.JWTService = auth.JWTAuthService()
		token := auth.LoginHandler(loginService, jwtService).Login(ctx)
		log.Println(token)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
		}
	})

	server.Run(":8080")
}

var sampleAnswers = []types.Answer{
	{Uid: "1", QuestionId: "1", Answer: "1", Time: time.Now()},
	{Uid: "2", QuestionId: "2", Answer: "2", Time: time.Now()},
}

func getSampleAnswers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sampleAnswers)
}

func postAnswer(c *gin.Context) {
	var answer types.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sampleAnswers = append(sampleAnswers, answer)
	c.IndentedJSON(http.StatusOK, answer)
}

func getSampleAnswerById(c *gin.Context) {
	id := c.Param("id")
	for _, answer := range sampleAnswers {
		if answer.Uid == id {
			c.IndentedJSON(http.StatusOK, answer)
			return
		}
	}
}
