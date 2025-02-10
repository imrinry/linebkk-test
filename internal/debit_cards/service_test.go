package debit_cards_test



import (
	"context"
	"errors"
	"line-bk-api/internal/debit_cards"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDebitCards(t *testing.T) {
	mockRepo := &debit_cards.MockDebitCardRepository{}
	service := debit_cards.NewDebitCardService(mockRepo)
	ctx := context.Background()

	t.Run("success - get from cache", func(t *testing.T) {
		expectedResponse := []debit_cards.DebitCardResponse{
			{
				CardID: "card1",
				UserID: "user1",
			},
		}

		mockRepo.On("GetDebitCardCache", mock.Anything, "debit_card:user1:0:10").
			Return(expectedResponse, nil).Once()

		responses, total, err := service.GetDebitCards(ctx, "user1", 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, responses)
		assert.Equal(t, 0, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("success - get from db", func(t *testing.T) {
		mockRepo.On("GetDebitCardCache", mock.Anything, "debit_card:user1:0:10").
			Return([]debit_cards.DebitCardResponse{}, redis.Nil).Once()

		debitCards := []debit_cards.DebitCard{
			{
				CardID: "card1",
				UserID: "user1",
			},
		}
		mockRepo.On("GetDebitCardByUserID", mock.Anything, "user1", 0, 10).
			Return(debitCards, nil).Once()

		mockRepo.On("GetCountDebitCards", mock.Anything, "user1").
			Return(1, nil).Once()

		expectedResponse := []debit_cards.DebitCardResponse{
			{
				CardID: "card1",
				UserID: "user1",
			},
		}
		mockRepo.On("SetDebitCardCache", mock.Anything, "debit_card:user1:0:10", expectedResponse, 5*time.Minute).
			Return(nil).Once()

		responses, total, err := service.GetDebitCards(ctx, "user1", 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, responses)
		assert.Equal(t, 1, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error - cache error", func(t *testing.T) {
		mockRepo.On("GetDebitCardCache", mock.Anything, "debit_card:user1:0:10").
			Return([]debit_cards.DebitCardResponse{}, errors.New("cache error")).Once()

		responses, total, err := service.GetDebitCards(ctx, "user1", 1, 10)

		assert.Error(t, err)
		assert.Nil(t, responses)
		assert.Equal(t, 0, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error - db error", func(t *testing.T) {
		mockRepo.On("GetDebitCardCache", mock.Anything, "debit_card:user1:0:10").
			Return([]debit_cards.DebitCardResponse{}, redis.Nil).Once()

		mockRepo.On("GetDebitCardByUserID", mock.Anything, "user1", 0, 10).
			Return([]debit_cards.DebitCard{}, errors.New("db error")).Once()

		mockRepo.On("GetCountDebitCards", mock.Anything, "user1").
			Return(0, errors.New("count error")).Once()

		responses, total, err := service.GetDebitCards(ctx, "user1", 1, 10)

		assert.Error(t, err)
		assert.Nil(t, responses)
		assert.Equal(t, 0, total)
		mockRepo.AssertExpectations(t)
	})
}
