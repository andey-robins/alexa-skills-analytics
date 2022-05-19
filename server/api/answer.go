package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/andey-robins/alexa-skills-analytics/server/db"
	"github.com/andey-robins/alexa-skills-analytics/server/types"
	"github.com/gin-gonic/gin"
)

func GetAnswers(c *gin.Context) {
	db := db.Connect()
	rows, err := db.Db.Query(`SELECT * FROM answers;`)
	if err != nil {
		log.Printf("[Err]: Error is selecting answers - %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	answers := make([]types.Answer, 0)
	for rows.Next() {
		var uid int
		var questionId, answer, createTime, updateTime string
		if err = rows.Scan(&uid, &questionId, &answer, &createTime, &updateTime); err != nil {
			log.Printf("[Err]: Error in reading row from answers - %v\n", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		answers = append(answers, types.Answer{
			Uid:         strconv.Itoa(uid),
			QuestionId:  questionId,
			Answer:      answer,
			LastUpdated: updateTime,
		})
	}

	c.JSON(http.StatusOK, answers)
}

func PostAnswer(c *gin.Context) {
	// this decodes the post info
	var newAnswer types.Answer
	if err := c.BindJSON(&newAnswer); err != nil {
		log.Printf("[Err]: Error in decoding new answer - %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// this prepares the query
	db := db.Connect()
	queryStatement := `INSERT INTO answers (qid, answer) VALUES (@QuestionId, @Answer);`
	query, err := db.Db.Prepare(queryStatement)
	if err != nil {
		log.Printf("[Err]: Error in preparing sql query for new answer - %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	defer query.Close()

	// this actually performs the query
	newRecord := query.QueryRowContext(context.Background(),
		sql.Named("QuestionId", newAnswer.QuestionId),
		sql.Named("Answer", newAnswer.Answer),
	)

	// check that it was successfully created by seeing if a row is returned
	// see query.QueryRowContext(...) for more info
	// TODO: Come back to refactor this and make it more clear/check it better
	var newID int
	newRecord.Scan(&newID)

	// reaching here means everything was created successfully
	c.JSON(http.StatusCreated, nil)
}
