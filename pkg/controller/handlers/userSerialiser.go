package handlers

// RegisterUserRequest represents the request payload for user registration.
type RegisterUserRequest struct {
	Id         int    `json:id`
	First_name string `json:"first_name" binding :"required"`
	Last_name  string `json:"last_name"  binding :"required"`
	User_name  string `json:"user_name"  binding :"required"`
	Password   string `json:"password"   binding :"required"`
	Email      string `json:"email"      binding :"required"`
	IsAdmin    bool   `json:"isAdmin"`
}

// RegisterUserRequest represents the response payload for user registration.
type RegisterUserResponse struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type UserLoginRequest struct {
	User_Name string `json:"user_name" binding :"required"`
	Password  string `json:"password" binding : "required"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	Error   string `json:"error"`
}

type GetUserByIdResponse struct {
	Id         int    `json:id`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	User_name  string `json:"user_name"`
	Email      string `json:"email"`
	Error      string `json:error`
}

type UpdateUserRequest struct {
	Id         int    `json:"id" binding:"required"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	User_name  string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required"`
}

type UpdateUserResponse struct {
	Message string `json:"message"`
	Error   string `json"error"`
}
