package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
	"time"

	"github.com/go-redis/redis/v8"
)

type AccountService interface {
	GetAccountByUserID(ctx context.Context, userID string, page int, limit int) ([]AccountResponse, int, error)
}

type service struct {
	accountRepository AccountRepository
}

func NewAccountService(accountRepository AccountRepository) AccountService {
	return &service{accountRepository: accountRepository}
}

func (s *service) GetAccountByUserID(ctx context.Context, userID string, page int, limit int) ([]AccountResponse, int, error) {

	offset, limit := utils.GetOffset(page, limit)
	cacheKey := fmt.Sprintf("account:%s:%d:%d", userID, offset, limit)
	cacheData, err := s.accountRepository.GetAccountCache(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		logs.Error(err)
		return nil, 0, err
	}

	if len(cacheData) > 0 {
		return cacheData, 0, nil
	}

	accounts, err := s.accountRepository.GetAccountByUserID(ctx, userID, offset, limit)

	if err != nil && err != sql.ErrNoRows {
		logs.Error(err)
		return nil, 0, err
	}

	accountResponses := make([]AccountResponse, len(accounts))
	for i, account := range accounts {
		accountResponses[i] = account.ToAccountResponse()

	}

	total, err := s.accountRepository.GetCountAccounts(ctx, userID)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}

	err = s.accountRepository.SetAccountCache(ctx, cacheKey, accountResponses, 5*time.Minute)
	if err != nil {
		logs.Error(err)
	}
	return accountResponses, total, nil
}
