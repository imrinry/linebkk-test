package account_test

import (
	"context"
	"database/sql"
	"errors"
	"line-bk-api/internal/account"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAccountByUserID(t *testing.T) {
	tests := []struct {
		name          string
		userID        string
		page          int
		limit         int
		cacheKey      string
		cacheHit      bool
		cachedData    []account.AccountResponse
		mockAccounts  []account.Account
		mockTotal     int
		mockErr       error
		mockTotalErr  error
		expectedErr   string
		expectedTotal int
	}{
		{
			name:     "success - get from cache",
			userID:   "user1",
			page:     1,
			limit:    10,
			cacheKey: "account:user1:0:10",
			cacheHit: true,
			cachedData: []account.AccountResponse{
				{AccountID: "1", UserID: "user1", Type: "Test Account"},
			},
			expectedTotal: 0,
		},
		{
			name:     "success - get from database",
			userID:   "user1",
			page:     1,
			limit:    10,
			cacheKey: "account:user1:0:10",
			cacheHit: false,
			mockAccounts: []account.Account{
				{AccountID: "1", UserID: "user1", Type: "Test Account"},
			},
			mockTotal:     1,
			expectedTotal: 1,
		},
		{
			name:        "error - cache error",
			userID:      "user1",
			page:        1,
			limit:       10,
			cacheKey:    "account:user1:0:10",
			cacheHit:    false,
			mockErr:     errors.New("cache error"),
			expectedErr: "cache error",
		},
		{
			name:        "error - database error",
			userID:      "user1",
			page:        1,
			limit:       10,
			cacheKey:    "account:user1:0:10",
			cacheHit:    false,
			mockErr:     sql.ErrConnDone,
			expectedErr: sql.ErrConnDone.Error(),
		},
		{
			name:         "error - count error",
			userID:       "user1",
			page:         1,
			limit:        10,
			cacheKey:     "account:user1:0:10",
			cacheHit:     false,
			mockAccounts: []account.Account{{AccountID: "1", UserID: "user1", Type: "Test Account"}},
			mockTotalErr: errors.New("count error"),
			expectedErr:  "count error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := account.NewMockAccountRepository()

			if !tt.cacheHit {
				if tt.mockErr != nil && !errors.Is(tt.mockErr, redis.Nil) {
					mockRepo.On("GetAccountCache", context.Background(), tt.cacheKey).Return([]account.AccountResponse{}, tt.mockErr)
				} else {
					mockRepo.On("GetAccountCache", context.Background(), tt.cacheKey).Return([]account.AccountResponse{}, redis.Nil)
					mockRepo.On("GetAccountByUserID", context.Background(), tt.userID, 0, tt.limit).Return(tt.mockAccounts, tt.mockErr)

					if tt.mockErr == nil {
						mockRepo.On("GetCountAccounts", context.Background(), tt.userID).Return(tt.mockTotal, tt.mockTotalErr)
						if tt.mockTotalErr == nil {
							mockRepo.On("SetAccountCache", context.Background(), tt.cacheKey, mock.Anything, 5*time.Minute).Return(nil)
						}
					}
				}
			} else {
				mockRepo.On("GetAccountCache", context.Background(), tt.cacheKey).Return(tt.cachedData, nil)
			}

			service := account.NewAccountService(mockRepo)

			// Act
			accounts, total, err := service.GetAccountByUserID(context.Background(), tt.userID, tt.page, tt.limit)

			// Assert
			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedTotal, total)

			if tt.cacheHit {
				assert.Equal(t, tt.cachedData, accounts)
			} else if len(accounts) > 0 {
				expectedResponses := make([]account.AccountResponse, len(tt.mockAccounts))
				for i, a := range tt.mockAccounts {
					expectedResponses[i] = a.ToAccountResponse()
				}
				assert.Equal(t, expectedResponses, accounts)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestNewAccountService(t *testing.T) {
	// Arrange
	mockRepo := account.NewMockAccountRepository()

	// Act
	service := account.NewAccountService(mockRepo)

	// Assert
	assert.NotNil(t, service)
	assert.Implements(t, (*account.AccountService)(nil), service)
}
