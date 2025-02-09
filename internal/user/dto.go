package user

type UserResponseDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
