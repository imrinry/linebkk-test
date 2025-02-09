package auth

type LoginWithPinCodeRequest struct {
	UserID  string `json:"user_id" validate:"required" example:"000018b0e1a211ef95a30242ac180002"`
	PinCode string `json:"pin_code" validate:"required" example:"123456"`
}

type LoginWithPasswordRequest struct {
	UserID   string `json:"user_id" validate:"required" example:"000018b0e1a211ef95a30242ac180002"`
	Password string `json:"password" validate:"required" example:"123456"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}
