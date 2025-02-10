package transactions

import (
	"context"
	"database/sql"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
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
	return transactionResponses, total, nil
}
