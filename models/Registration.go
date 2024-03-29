package models

type Registration struct {
	RegLink string `json:"regLink"`
}

type User struct {
	Id              int
	Email, Password string
	Confirmed       bool
}

type CompleteRegistration struct {
	Token string `json:"token"`
}
