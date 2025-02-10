package account

import (
	"time"

	"github.com/shopspring/decimal"
)

type Account struct {
	AccountID       string          `db:"account_id"`
	UserID          string          `db:"user_id"`
	Type            string          `db:"type"`
	Currency        string          `db:"currency"`
	AccountNumber   string          `db:"account_number"`
	Issuer          string          `db:"issuer"`
	DummyCol3       string          `db:"dummy_col_3"`
	Status          string          `db:"status"`
	CreatedAt       time.Time       `db:"created_at"`
	UpdatedAt       time.Time       `db:"updated_at"`
	DeletedAt       *time.Time      `db:"deleted_at"`
	AccountDetail   AccountDetail   `db:"account_detail"`
	AccountBalances AccountBalances `db:"account_balances"`
	AccountFlags    []AccountFlag   `db:"account_flags"`
}

type AccountDetail struct {
	AccountID       string  `db:"account_id"`
	UserID          string  `db:"user_id"`
	Color           string  `db:"color"`
	IsMainAccount   bool    `db:"is_main_account"`
	Progress        int     `db:"progress"`
	AccountNickname *string `db:"account_nickname"`
}

type AccountBalances struct {
	AccountID string          `db:"account_id"`
	Amount    decimal.Decimal `db:"amount"`
}

type AccountFlag struct {
	FlagID    string    `db:"flag_id"`
	AccountID string    `db:"account_id"`
	FlagType  string    `db:"flag_type"`
	FlagValue string    `db:"flag_value"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (a *Account) ToAccountResponse() AccountResponse {
	res := AccountResponse{
		AccountID:      a.AccountID,
		UserID:         a.UserID,
		Type:           a.Type,
		Currency:       a.Currency,
		AccountNumber:  a.AccountNumber,
		Issuer:         a.Issuer,
		Status:         a.Status,
		AccountDetail:  a.AccountDetail.ToAccountDetailResponse(),
		AccountBalance: a.AccountBalances.ToAccountBalanceResponse(),
	}

	res.AccountFlag = make([]AccountFlagResponse, len(a.AccountFlags))
	for i, flag := range a.AccountFlags {
		res.AccountFlag[i] = flag.ToAccountFlagResponse()
	}

	return res
}

func (a *AccountDetail) ToAccountDetailResponse() AccountDetailResponse {
	return AccountDetailResponse{
		AccountID:       a.AccountID,
		UserID:          a.UserID,
		Color:           a.Color,
		IsMainAccount:   a.IsMainAccount,
		Progress:        a.Progress,
		AccountNickname: a.AccountNickname,
	}
}

func (a *AccountBalances) ToAccountBalanceResponse() AccountBalanceResponse {
	return AccountBalanceResponse{
		AccountID: a.AccountID,
		Amount:    a.Amount,
	}
}

func (a *AccountFlag) ToAccountFlagResponse() AccountFlagResponse {
	return AccountFlagResponse{
		FlagID:    a.FlagID,
		AccountID: a.AccountID,
		FlagType:  a.FlagType,
		FlagValue: a.FlagValue,
	}
}
