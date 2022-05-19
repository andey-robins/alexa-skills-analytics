package types

type User struct {
	Uid         string `json:"uid"`
	AlexaId     string `json:"alexaId"`
	AlexaDevice string `json:"alexaDevice"`
	Email       string `json:"email"`
	LastUpdated string `json:"lastUpdated"`
}

type Answer struct {
	Uid         string `json:"uid"`
	QuestionId  string `json:"questionId"`
	Answer      string `json:"answer"`
	LastUpdated string `json:"time"`
}
