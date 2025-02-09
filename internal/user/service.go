package user

import (
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
)

type UserService interface {
	GetUsers(page, limit int) ([]User, error)
}

type service struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &service{userRepo: userRepo}
}

func (s *service) GetUsers(page, limit int) ([]User, error) {

	offset := utils.GetOffset(page, limit)
	users, err := s.userRepo.GetAllUsers(offset, limit)
	if err != nil {
		logs.Error(err)
		return nil, utils.NewUnexpectedError()
	}

	return users, nil
}
