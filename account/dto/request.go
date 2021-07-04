package dto

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
)
