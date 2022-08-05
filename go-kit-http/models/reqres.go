package models

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
}

type GetUserRequest struct {
	Id string `json:"id"`
}

type GetUserResponse struct {
	Email string `json:"email"`
}
