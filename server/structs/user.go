package structs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string
	Password string
}

// from a json body, read and unmarshall a user object
// returns a filled user object if the request completes successfully
// and an error otherwise
func BodyToUser(r *http.Request) (User, error) {
	var u User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return u, fmt.Errorf("error reading request body: %v", err)
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return u, fmt.Errorf("error unmarshalling json request: %v", err)
	}

	return u, nil
}
