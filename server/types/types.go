package types

import "time"

type User struct {
	Uid         string    `json:"uid"`
	AlexaId     string    `json:"alexaId"`
	AlexaDevice string    `json:"alexaDevice"`
	Email       string    `json:"email"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type Answer struct {
	Uid        string    `json:"uid"`
	QuestionId string    `json:"questionId"`
	Answer     string    `json:"answer"`
	Time       time.Time `json:"time"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
