package debit_cards

import (
	"context"
	"line-bk-api/pkg/logs"
	"line-bk-api/pkg/utils"
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
	debitCards, err := s.debitCardRepository.GetDebitCardByUserID(ctx, userID, offset, limit)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
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

	return debitCardResponses, total, nil
}
