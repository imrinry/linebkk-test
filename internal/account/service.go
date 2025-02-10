package account

import (
	"database/sql"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
)

type AccountService interface {
	GetAccountByUserID(userID string, page int, limit int) ([]AccountResponse, int, error)
}

type service struct {
	accountRepository AccountRepository
}

func NewAccountService(accountRepository AccountRepository) AccountService {
	return &service{accountRepository: accountRepository}
}

func (s *service) GetAccountByUserID(userID string, page int, limit int) ([]AccountResponse, int, error) {

	offset, limit := utils.GetOffset(page, limit)
	accounts, err := s.accountRepository.GetAccountByUserID(userID, offset, limit)

	if err != nil && err != sql.ErrNoRows {
		logs.Error(err)
		return nil, 0, err
	}

	accountResponses := make([]AccountResponse, len(accounts))
	for i, account := range accounts {
		accountResponses[i] = account.ToAccountResponse()
		
	}

	total, err := s.accountRepository.GetCountAccounts(userID)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}

	return accountResponses, total, nil
}
