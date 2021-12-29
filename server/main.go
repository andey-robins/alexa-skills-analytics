package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andey-robins/alexa-skills-analytics/server/auth"
)

func main() {
	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/login", testLogin)

	log.Println("Starting go server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func testLogin(w http.ResponseWriter, r *http.Request) {
	response, err := getTestJson()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprint(w, string(response))
}

func getTestJson() ([]byte, error) {
	j := make(map[string]string)
	j["token"] = "token123"
	return json.MarshalIndent(j, "", "  ")
}
