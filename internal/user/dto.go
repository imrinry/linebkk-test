package user

type UserResponseDTO struct {
	UserID       string  `json:"user_id"`
	Name         string  `json:"name"`
	Dummy        string  `json:"dummy_col_1"`
	Email        *string `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	ProfileImage *string `json:"profile_image"`
}

type CreateUserRequestDTO struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

type UserGreetingResponseDTO struct {
	Greeting string `json:"greeting"`
}
