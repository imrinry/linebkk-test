package user

import "time"

type User struct {
	UserID       string     `db:"user_id"`
	Name         string     `db:"name"`
	Dummy        string     `db:"dummy_col_1"`
	Email        *string    `db:"email"`
	PhoneNumber  *string    `db:"phone_number"`
	ProfileImage *string    `db:"profile_image"`
	PinCode      string     `db:"pin_code"`
	Password     string     `db:"password"`
	CreatedAt    *time.Time `db:"created_at"`
}

func (u *User) ToUserResponse() UserResponseDTO {
	return UserResponseDTO{
		UserID:       u.UserID,
		Name:         u.Name,
		Dummy:        u.Dummy,
		Email:        u.Email,
		PhoneNumber:  u.PhoneNumber,
		ProfileImage: u.ProfileImage,
	}
}

type UserGreeting struct {
	UserID   string `db:"user_id"`
	Greeting string `db:"greeting"`
	Dummy    string `db:"dummy_col_2"`
}

func (u *UserGreeting) ToUserGreetingResponse() UserGreetingResponseDTO {
	return UserGreetingResponseDTO{
		Greeting: u.Greeting,
	}
}
