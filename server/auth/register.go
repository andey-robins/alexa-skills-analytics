package auth

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"golang.org/x/crypto/argon2"

	database "github.com/andey-robins/alexa-skills-analytics/server/db"
	structs "github.com/andey-robins/alexa-skills-analytics/server/structs"
)

func Register(w http.ResponseWriter, r *http.Request) {

	// precondition checks
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong method type to /register\n")
		return
	}

	// unmarshall json post request
	u, err := structs.BodyToUser(r)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}

	// verify user doesn't already exist
	instance := database.Connect()
	rows, err := instance.Db.Query(context.Background(), "SELECT COUNT(*) FROM Users WHERE username=$1;", u.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Problem querying to database: %v\n", err)
		return
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Problem reading database row: %v\n", err)
			return
		}
		if count != 0 {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprint(w, "User already exists.\n")
			return
		}
	}

	// salt and hash the password before storing it in the DB
	salt := make([]byte, 32)
	_, err = rand.Read(salt)
	if err != nil {
		fmt.Fprintf(w, "Unable to read enough random bytes: %v", err)
		return
	}

	passHash := argon2.IDKey([]byte(u.Password), salt, 10, 64*1024, 8, 32)

	// Store into DB
	_, err = instance.Db.Exec(context.Background(), "INSERT INTO Users (username, salt, hash) VALUES ($1, $2, $3)", u.Username, salt, passHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Problem creating user: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Successfully created user.\n")
}
