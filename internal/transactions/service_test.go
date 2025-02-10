package transactions

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTransactionByUserID(t *testing.T) {
	mockRepo := NewMockTransactionsRepository()
	service := NewTransactionService(mockRepo)
	ctx := context.Background()

	t.Run("success get from cache", func(t *testing.T) {
		expectedResp := []TransactionResponse{
			{
				TransactionID: "1",
				UserID:        "user1",
				Name:          "Test Transaction",
				Image:         "test.jpg",
				IsBank:        1,
			},
		}

		mockRepo.On("GetTransactionCache", mock.Anything, mock.Anything).
			Return(expectedResp, nil).Once()

		resp, total, err := service.GetTransactionByUserID(ctx, "user1", 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
		assert.Equal(t, 0, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("success get from db", func(t *testing.T) {
		transactions := []Transaction{
			{
				TransactionID: "1",
				UserID:        "user1",
				Name:          "Test Transaction",
				Image:         "test.jpg",
				IsBank:        1,
			},
		}

		expectedResp := []TransactionResponse{
			{
				TransactionID: "1",
				UserID:        "user1",
				Name:          "Test Transaction",
				Image:         "test.jpg",
				IsBank:        1,
			},
		}

		mockRepo.On("GetTransactionCache", mock.Anything, mock.Anything).
			Return([]TransactionResponse{}, redis.Nil).Once()

		mockRepo.On("GetTransactionByUserID", mock.Anything, "user1", 0, 10).
			Return(transactions, nil).Once()

		mockRepo.On("GetTransactionCountByUserID", mock.Anything, "user1").
			Return(1, nil).Once()

		mockRepo.On("SetTransactionCache", mock.Anything, mock.Anything, expectedResp, 5*time.Minute).
			Return(nil).Once()

		resp, total, err := service.GetTransactionByUserID(ctx, "user1", 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
		assert.Equal(t, 1, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error getting from cache", func(t *testing.T) {
		mockRepo.On("GetTransactionCache", mock.Anything, mock.Anything).
			Return([]TransactionResponse{}, errors.New("cache error")).Once()

		resp, total, err := service.GetTransactionByUserID(ctx, "user1", 1, 10)

		assert.Error(t, err)
		assert.Empty(t, resp)
		assert.Equal(t, 0, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error getting from db", func(t *testing.T) {
		mockRepo.On("GetTransactionCache", mock.Anything, mock.Anything).
			Return([]TransactionResponse{}, redis.Nil).Once()

		mockRepo.On("GetTransactionByUserID", mock.Anything, "user1", 0, 10).
			Return([]Transaction{}, sql.ErrNoRows).Once()

		mockRepo.On("GetTransactionCountByUserID", mock.Anything, "user1").
			Return(0, nil).Once()

		mockRepo.On("SetTransactionCache", mock.Anything, mock.Anything, []TransactionResponse{}, 5*time.Minute).
			Return(nil).Once()

		resp, total, err := service.GetTransactionByUserID(ctx, "user1", 1, 10)

		assert.NoError(t, err)
		assert.Empty(t, resp)
		assert.Equal(t, 0, total)
		mockRepo.AssertExpectations(t)
	})
}
