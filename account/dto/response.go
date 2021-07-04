package dto

type (
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserResponse struct {
		Email string `json:"email"`
	}
)
