package main

import (
	"log"
	"net/http"

	"github.com/andey-robins/alexa-skills-analytics/server/auth"
)

func main() {
	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/login", auth.Login)

	log.Println("Starting go server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
