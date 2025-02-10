package account

import "github.com/shopspring/decimal"

type AccountResponse struct {
	AccountID      string                 `json:"account_id"`
	UserID         string                 `json:"user_id"`
	Type           string                 `json:"type"`
	Currency       string                 `json:"currency"`
	AccountNumber  string                 `json:"account_number"`
	Issuer         string                 `json:"issuer"`
	Status         string                 `json:"status"`
	AccountDetail  AccountDetailResponse  `json:"account_detail"`
	AccountBalance AccountBalanceResponse `json:"account_balance"`
	AccountFlag    []AccountFlagResponse  `json:"account_flag"`
}

type AccountDetailResponse struct {
	AccountID       string  `json:"account_id"`
	UserID          string  `json:"user_id"`
	Color           string  `json:"color"`
	IsMainAccount   bool    `json:"is_main_account"`
	Progress        int     `json:"progress"`
	AccountNickname *string `json:"account_nickname"`
}

type AccountBalanceResponse struct {
	AccountID string          `json:"account_id"`
	Amount    decimal.Decimal `json:"amount"`
}

type AccountFlagResponse struct {
	FlagID    string `json:"flag_id"`
	AccountID string `json:"account_id"`
	FlagType  string `json:"flag_type"`
	FlagValue string `json:"flag_value"`
}
