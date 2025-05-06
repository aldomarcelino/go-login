package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}


type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}



