package user

import (
	"database/sql"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
)

type UserService interface {
	GetUsers(page, limit int) ([]UserResponseDTO, int, error)
	GetUserByID(id string) (UserResponseDTO, error)
	GetUserGreeting(id string, page, limit int) ([]UserGreetingResponseDTO, int, error)
}

type service struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &service{userRepo: userRepo}
}

func (s *service) GetUsers(page, limit int) ([]UserResponseDTO, int, error) {
	offset, limit := utils.GetOffset(page, limit)
	users, err := s.userRepo.GetAllUsers(offset, limit)
	if err != nil {
		logs.Error(err)
		return nil, 0, utils.NewUnexpectedError()
	}
	total, err := s.userRepo.GetCountUsers()
	if err != nil {
		logs.Error(err)
		return nil, 0, utils.NewUnexpectedError()
	}

	userResponses := make([]UserResponseDTO, len(users))
	for i, user := range users {
		userResponses[i] = user.ToUserResponse()
	}

	return userResponses, total, nil
}

func (s *service) GetUserByID(id string) (UserResponseDTO, error) {
	if err := validateUserID(id); err != nil {
		return UserResponseDTO{}, err
	}

	user, err := s.userRepo.GetUserByID(id)

	if err == sql.ErrNoRows {
		return UserResponseDTO{}, utils.NewNotFoundError("User not found")
	}

	if err != nil && err != sql.ErrNoRows {
		logs.Error(err)
		return UserResponseDTO{}, utils.NewUnexpectedError()
	}

	return user.ToUserResponse(), nil
}

func validateUserID(id string) error {
	if id == "" {
		return utils.NewBadRequestError("user id is required")
	}
	return nil
}

func (s *service) GetUserGreeting(id string, page, limit int) ([]UserGreetingResponseDTO, int, error) {
	offset, limit := utils.GetOffset(page, limit)
	greetings, err := s.userRepo.GetUserGreeting(id, offset, limit)
	if err != nil {
		logs.Error(err)
		return nil, 0, utils.NewUnexpectedError()
	}
	total, err := s.userRepo.GetUserGreetingCount(id)
	if err != nil {
		logs.Error(err)
		return nil, 0, utils.NewUnexpectedError()
	}

	greetingResponses := make([]UserGreetingResponseDTO, len(greetings))
	for i, greeting := range greetings {
		greetingResponses[i] = greeting.ToUserGreetingResponse()
	}

	return greetingResponses, total, nil
}
