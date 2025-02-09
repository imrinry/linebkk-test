package user

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (u *User) ToUserResponse() UserResponseDTO {
	return UserResponseDTO{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}
