package models

type User struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	User_name  string `json:"user_name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IsAdmin    bool   `json:"isAdmin"`
}
