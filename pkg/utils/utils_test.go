package utils_test

import (
	"errors"
	"fmt"
	"line-bk-api/config"
	"line-bk-api/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestGetOffset1(t *testing.T) {
	tests := []struct {
		name       string
		page       int
		limit      int
		wantOffset int
		wantLimit  int
	}{
		{
			name:       "valid page and limit",
			page:       2,
			limit:      10,
			wantOffset: 10,
			wantLimit:  10,
		},
		{
			name:       "page less than 1",
			page:       0,
			limit:      10,
			wantOffset: 0,
			wantLimit:  10,
		},
		{
			name:       "limit less than 1",
			page:       1,
			limit:      0,
			wantOffset: 0,
			wantLimit:  config.DefaultLimit,
		},
		{
			name:       "limit greater than max limit",
			page:       1,
			limit:      1000,
			wantOffset: 0,
			wantLimit:  config.MaxLimit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOffset, gotLimit := utils.GetOffset(tt.page, tt.limit)
			assert.Equal(t, tt.wantOffset, gotOffset)
			assert.Equal(t, tt.wantLimit, gotLimit)
		})
	}
}

func TestGetLimit2(t *testing.T) {
	tests := []struct {
		name      string
		limit     int
		wantLimit int
	}{
		{
			name:      "valid limit",
			limit:     10,
			wantLimit: 10,
		},
		{
			name:      "limit less than 1",
			limit:     0,
			wantLimit: config.DefaultLimit,
		},
		{
			name:      "limit greater than max limit",
			limit:     1000,
			wantLimit: config.MaxLimit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLimit := utils.GetLimit(tt.limit)
			assert.Equal(t, tt.wantLimit, gotLimit)
		})
	}
}

func TestGetTotalPages2(t *testing.T) {
	tests := []struct {
		name      string
		total     int
		limit     int
		wantPages int
	}{
		{
			name:      "exact division",
			total:     100,
			limit:     10,
			wantPages: 10,
		},
		{
			name:      "with remainder",
			total:     105,
			limit:     10,
			wantPages: 11,
		},
		{
			name:      "zero total",
			total:     0,
			limit:     10,
			wantPages: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPages := utils.GetTotalPages(tt.total, tt.limit)
			assert.Equal(t, tt.wantPages, gotPages)
		})
	}
}

func TestGetNextPage2(t *testing.T) {
	tests := []struct {
		name         string
		page         int
		totalPages   int
		wantNextPage int
	}{
		{
			name:         "has next page",
			page:         1,
			totalPages:   2,
			wantNextPage: 2,
		},
		{
			name:         "no next page",
			page:         2,
			totalPages:   2,
			wantNextPage: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNextPage := utils.GetNextPage(tt.page, tt.totalPages)
			assert.Equal(t, tt.wantNextPage, gotNextPage)
		})
	}
}

func TestGetPreviousPage2(t *testing.T) {
	tests := []struct {
		name             string
		page             int
		wantPreviousPage int
	}{
		{
			name:             "has previous page",
			page:             2,
			wantPreviousPage: 1,
		},
		{
			name:             "no previous page",
			page:             1,
			wantPreviousPage: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPreviousPage := utils.GetPreviousPage(tt.page)
			assert.Equal(t, tt.wantPreviousPage, gotPreviousPage)
		})
	}
}

func TestNotFoundError2(t *testing.T) {
	tests := []struct {
		name        string
		message     string
		wantCode    int
		wantMessage string
	}{
		{
			name:        "not found error",
			message:     "record not found",
			wantCode:    404,
			wantMessage: "record not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.NewNotFoundError(tt.message)
			appErr, ok := err.(utils.AppError)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, appErr.Code)
			assert.Equal(t, tt.wantMessage, appErr.Message)
		})
	}
}

func TestUnexpectedError2(t *testing.T) {
	tests := []struct {
		name        string
		wantCode    int
		wantMessage string
	}{
		{
			name:        "unexpected error",
			wantCode:    500,
			wantMessage: "unexpected error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.NewUnexpectedError()
			appErr, ok := err.(utils.AppError)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, appErr.Code)
			assert.Equal(t, tt.wantMessage, appErr.Message)
		})
	}
}

func TestValidationError2(t *testing.T) {
	tests := []struct {
		name        string
		message     string
		wantCode    int
		wantMessage string
	}{
		{
			name:        "validation error",
			message:     "invalid input",
			wantCode:    422,
			wantMessage: "invalid input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.NewValidationError(tt.message)
			appErr, ok := err.(utils.AppError)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, appErr.Code)
			assert.Equal(t, tt.wantMessage, appErr.Message)
		})
	}
}

func TestBadRequestError2(t *testing.T) {
	tests := []struct {
		name        string
		message     string
		wantCode    int
		wantMessage string
	}{
		{
			name:        "bad request error",
			message:     "bad request",
			wantCode:    400,
			wantMessage: "bad request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.NewBadRequestError(tt.message)
			appErr, ok := err.(utils.AppError)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, appErr.Code)
			assert.Equal(t, tt.wantMessage, appErr.Message)
		})
	}
}

func TestAppError_Error_Message2(t *testing.T) {
	tests := []struct {
		name        string
		appError    utils.AppError
		wantMessage string
	}{
		{
			name: "error message",
			appError: utils.AppError{
				Code:    400,
				Message: "test error",
			},
			wantMessage: "test error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantMessage, tt.appError.Error())
		})
	}
}

func TestHandleResponse2(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		wantStatus  int
		wantMessage string
	}{
		{
			name: "app response",
			input: utils.AppResponse{
				Message: "success",
				Code:    200,
				Data:    "test data",
			},
			wantStatus:  200,
			wantMessage: "success",
		},
		{
			name: "app pagination response",
			input: utils.AppPaginationResponse{
				Message:    "success",
				Code:       200,
				Data:       "test data",
				Total:      100,
				Page:       1,
				Limit:      10,
				TotalPages: 10,
				NextPage:   2,
				PrevPage:   0,
			},
			wantStatus:  200,
			wantMessage: "success",
		},
		{
			name:        "invalid response type",
			input:       "invalid",
			wantStatus:  500,
			wantMessage: "Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/test", func(c *fiber.Ctx) error {
				return utils.HandleResponse(c, tt.input)
			})

			req := httptest.NewRequest("GET", "/test", nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}

func TestGenerateAccessToken(t *testing.T) {
	// Mock config values
	config.AccessTokenExpired = "1h"
	config.JWTIssuer = "test_issuer"
	config.JWTSecret = "test_secret"
	userID := "12345"
	tokenStr, err := utils.GenerateAccessToken(userID)
	assert.NoError(t, err, "Expected no error when generating access token")
	assert.NotEmpty(t, tokenStr, "Generated token should not be empty")

	// Parse and validate token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.JWTSecret), nil
	})

	assert.NoError(t, err, "Expected no error when parsing the token")
	assert.NotNil(t, token, "Parsed token should not be nil")
	assert.True(t, token.Valid, "Token should be valid")

	// Verify claims
	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok, "Expected token claims to be of type jwt.MapClaims")
	assert.Equal(t, userID, claims["sub"], "User ID in token should match")
	assert.Equal(t, config.JWTIssuer, claims["iss"], "Issuer should match")

	expTime := int64(claims["exp"].(float64))
	assert.Greater(t, expTime, time.Now().Unix(), "Expiration time should be in the future")

	// Error case: Invalid expiration duration
	config.AccessTokenExpired = "invalid_duration"
	tokenStr, err = utils.GenerateAccessToken(userID)
	assert.Error(t, err, "Expected an error due to invalid expiration duration")
	assert.Empty(t, tokenStr, "Token should be empty when an error occurs")
}

func TestGenerateRefreshToken(t *testing.T) {
	// Mock config values
	config.RefreshTokenExpired = "1h"
	config.JWTIssuer = "test_issuer"
	config.JWTSecret = "test_secret"
	userID := "12345"
	tokenStr, err := utils.GenerateRefreshToken(userID)
	assert.NoError(t, err, "Expected no error when generating refresh token")
	assert.NotEmpty(t, tokenStr, "Generated token should not be empty")

	// Parse and validate token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.JWTSecret), nil
	})

	assert.NoError(t, err, "Expected no error when parsing the token")
	assert.NotNil(t, token, "Parsed token should not be nil")
	assert.True(t, token.Valid, "Token should be valid")

	// Verify claims
	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok, "Expected token claims to be of type jwt.MapClaims")
	assert.Equal(t, userID, claims["sub"], "User ID in token should match")
	assert.Equal(t, config.JWTIssuer, claims["iss"], "Issuer should match")

	expTime := int64(claims["exp"].(float64))
	assert.Greater(t, expTime, time.Now().Unix(), "Expiration time should be in the future")

	// Error case: Invalid expiration duration
	config.RefreshTokenExpired = "invalid_duration"
	tokenStr, err = utils.GenerateRefreshToken(userID)
	assert.Error(t, err, "Expected an error due to invalid expiration duration")
	assert.Empty(t, tokenStr, "Token should be empty when an error occurs")
}

func TestValidateAccessToken(t *testing.T) {
	// Mock config values
	config.JWTSecret = "test_secret"
	config.AccessTokenExpired = "1h"
	config.RefreshTokenExpired = "1h"
	config.JWTIssuer = "test_issuer"

	userID := "12345"
	tokenStr, err := utils.GenerateAccessToken(userID)
	assert.NoError(t, err, "Expected no error when generating a valid token")
	assert.NotEmpty(t, tokenStr, "Generated token should not be empty")
	fmt.Println("Generated Token:", tokenStr)

	// Valid token
	validatedUserID, err := utils.ValidateAccessToken(tokenStr)
	assert.NoError(t, err, "Expected no error when validating a valid token")
	assert.Equal(t, userID, validatedUserID, "Validated user ID should match the original user ID")

	// Invalid token
	invalidToken := "invalid.token.string"
	validatedUserID, err = utils.ValidateAccessToken(invalidToken)
	assert.Error(t, err, "Expected an error when validating an invalid token")
	assert.Empty(t, validatedUserID, "Validated user ID should be empty when token is invalid")
}

func TestHandleError(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	// Test AppError case
	appErr := utils.AppError{
		Message: "Test error",
		Code:    400,
	}
	err := utils.HandleError(ctx, appErr)
	assert.NoError(t, err)
	assert.Equal(t, 400, ctx.Response().StatusCode())

	// Test generic error case
	genericErr := errors.New("generic error")
	err = utils.HandleError(ctx, genericErr)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Response().StatusCode())

	// Test default case with nil error
	err = utils.HandleError(ctx, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Response().StatusCode())
}


