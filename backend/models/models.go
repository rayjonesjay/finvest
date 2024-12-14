package models

type Auth struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}