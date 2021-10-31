package auth

import (
	"context"
	"fmt"
	"net/http"

	database "github.com/andey-robins/alexa-skills-analytics/server/db"
	structs "github.com/andey-robins/alexa-skills-analytics/server/structs"
	"golang.org/x/crypto/argon2"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// precondition checks
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong method type to /register\n")
		return
	}

	// unmarshall json post request
	u, err := structs.BodyToUser(r)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}

	// check if user exists and fetch their info if they do
	instance := database.Connect()
	rows, err := instance.Db.Query(context.Background(), "SELECT username, salt, hash FROM Users WHERE username=$1;", u.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Problem querying to database: %v\n", err)
		return
	}
	defer rows.Close()

	var username string
	var salt []byte
	var storedHash []byte
	var someRows = false
	for rows.Next() {
		someRows = true
		err = rows.Scan(&username, &salt, &storedHash)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Problem reading database row: %v\n", err)
			return
		}
		if username == "" || salt == nil || storedHash == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "User doesn't exist yet. Perhaps you meant to register?\n")
			return
		}
	}

	if !someRows {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "User doesn't exist yet. Perhaps you meant to register?\n")
		return
	}

	// salt and hash the password
	passHash := argon2.IDKey([]byte(u.Password), salt, 10, 64*1024, 8, 32)

	// verify salt and hashed password in linear time
	// (hoping to avoid timing sidechannel)
	var equal = true
	for i, b := range passHash {
		if b != storedHash[i] {
			equal = false
		}
	}

	if !equal {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid Password.\n")
		return
	}

	// log user in
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Successfully logged in as %s.\n", u.Username)
}
