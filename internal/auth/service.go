package auth

import (
	"fmt"
	"line-bk-api/internal/user"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
)

type AuthService interface {
	LoginWithPinCode(userID string, pinCode string) (LoginResponse, error)
	LoginWithPassword(userID string, password string) (LoginResponse, error)
}

type authService struct {
	authRepository AuthRepository
	userRepository user.UserRepository
}

func NewAuthService(authRepository AuthRepository, userRepository user.UserRepository) AuthService {
	return &authService{authRepository: authRepository, userRepository: userRepository}
}

func (s *authService) LoginWithPinCode(userID string, pinCode string) (LoginResponse, error) {
	fmt.Println("LoginWithPinCode", userID, pinCode)
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return LoginResponse{}, err
	}

	if user.PinCode != pinCode {
		return LoginResponse{}, utils.AppError{
			Message: "Invalid pin code",
			Code:    401,
		}
	}

	accessToken, err := utils.GenerateAccessToken(user.UserID)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.UserID)
	if err != nil {
		return LoginResponse{}, err
	}

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *authService) LoginWithPassword(userID string, password string) (LoginResponse, error) {

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		logs.Error(err.Error())
		return LoginResponse{}, err
	}

	if user.Password != password {
		return LoginResponse{}, utils.AppError{
			Message: "Invalid password",
			Code:    401,
		}
	}

	accessToken, err := utils.GenerateAccessToken(user.UserID)
	if err != nil {
		logs.Error(err.Error())
		return LoginResponse{}, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.UserID)
	if err != nil {
		logs.Error(err.Error())
		return LoginResponse{}, err
	}

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}
