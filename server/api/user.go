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

func GetUsers(c *gin.Context) {
	db := db.Connect()
	rows, err := db.Db.Query(`SELECT * FROM users;`)
	if err != nil {
		log.Printf("[Err]: Error is selecting users - %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	users := make([]types.User, 0)
	for rows.Next() {
		var uid int
		var alexaId, alexaDevice, email, createTime, updateTime string
		if err = rows.Scan(&uid, &alexaId, &alexaDevice, &email, &createTime, &updateTime); err != nil {
			log.Printf("[Err]: Error in reading row from users - %v\n", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		users = append(users, types.User{
			Uid:         strconv.Itoa(uid),
			AlexaId:     alexaId,
			AlexaDevice: alexaDevice,
			Email:       email,
			LastUpdated: updateTime,
		})
	}

	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	// this decodes the post info
	var newUser types.User
	if err := c.BindJSON(&newUser); err != nil {
		log.Printf("[Err]: Error in decoding new user - %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// this prepares the query
	db := db.Connect()
	queryStatement := `INSERT INTO users (alexaId, alexaDevice, email ) VALUES (@AlexaId, @AlexaDevice, @Email);`
	query, err := db.Db.Prepare(queryStatement)
	if err != nil {
		log.Printf("[Err]: Error in preparing sql query for new user - %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	defer query.Close()

	// this actually performs the query
	newRecord := query.QueryRowContext(context.Background(),
		sql.Named("AlexaId", newUser.AlexaId),
		sql.Named("AlexaDevice", newUser.AlexaDevice),
		sql.Named("Email", newUser.Email),
	)

	// check that it was successfully created by seeing if a row is returned
	// see query.QueryRowContext(...) for more info
	// TODO: Come back to refactor this and make it more clear/check it better
	var newID int
	newRecord.Scan(&newID)

	// reaching here means everything was created successfully
	c.JSON(http.StatusCreated, nil)
}
