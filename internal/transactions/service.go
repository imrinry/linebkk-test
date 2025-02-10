package transactions

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

type TransactionService interface {
	GetTransactionByUserID(ctx context.Context, userID string, page int, limit int) ([]TransactionResponse, int, error)
}

type transactionService struct {
	transactionRepository TransactionRepository
}

func NewTransactionService(transactionRepository TransactionRepository) TransactionService {
	return &transactionService{transactionRepository: transactionRepository}
}

func (s *transactionService) GetTransactionByUserID(ctx context.Context, userID string, page int, limit int) ([]TransactionResponse, int, error) {

	offset, limit := utils.GetOffset(page, limit)
	cacheKey := fmt.Sprintf("transaction:%s:%d:%d", userID, offset, limit)
	cacheData, err := s.transactionRepository.GetTransactionCache(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		logs.Error(err)
		return []TransactionResponse{}, 0, err
	}

	if len(cacheData) > 0 {
		return cacheData, 0, nil
	}

	transactions, err := s.transactionRepository.GetTransactionByUserID(ctx, userID, offset, limit)
	if err != nil && err != sql.ErrNoRows {
		logs.Error(err)
		return nil, 0, err
	}
	total, err := s.transactionRepository.GetTransactionCountByUserID(ctx, userID)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	transactionResponses := make([]TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		transactionResponses[i] = *transaction.ToTransactionResponse()
	}
	err = s.transactionRepository.SetTransactionCache(ctx, cacheKey, transactionResponses, 5*time.Minute)
	if err != nil {
		logs.Error(err)
	}
	return transactionResponses, total, nil
}
