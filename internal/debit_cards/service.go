package debit_cards

import (
	"context"
	"errors"
	"fmt"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
	"time"

	"github.com/go-redis/redis/v8"
)

type DebitCardService interface {
	GetDebitCards(ctx context.Context, userID string, page int, limit int) ([]DebitCardResponse, int, error)
}

type debitCardService struct {
	debitCardRepository DebitCardRepository
}

func NewDebitCardService(debitCardRepository DebitCardRepository) DebitCardService {
	return &debitCardService{debitCardRepository: debitCardRepository}
}

func (s *debitCardService) GetDebitCards(ctx context.Context, userID string, page int, limit int) ([]DebitCardResponse, int, error) {

	offset, limit := utils.GetOffset(page, limit)
	cacheKey := fmt.Sprintf("debit_card:%s:%d:%d", userID, offset, limit)
	cacheData, err := s.debitCardRepository.GetDebitCardCache(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		logs.Error(err)
		return nil, 0, err
	}

	if len(cacheData) > 0 {
		return cacheData, 0, nil
	}
	debitCards, err := s.debitCardRepository.GetDebitCardByUserID(ctx, userID, offset, limit)
	if err != nil {
		logs.Error(err)
	}
	total, err := s.debitCardRepository.GetCountDebitCards(ctx, userID)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}

	debitCardResponses := make([]DebitCardResponse, len(debitCards))
	for i, debitCard := range debitCards {
		debitCardResponses[i] = debitCard.ToDebitCardResponse()
	}

	err = s.debitCardRepository.SetDebitCardCache(ctx, cacheKey, debitCardResponses, 5*time.Minute)
	if err != nil {
		logs.Error(err)
	}

	return debitCardResponses, total, nil
}
