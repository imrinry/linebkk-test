package account

import (
	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	GetAccountByUserID(userID string, offset int, limit int) ([]Account, error)
	GetCountAccounts(userID string) (int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &repository{db: db}
}

func (r *repository) GetAccountByUserID(userID string, offset int, limit int) ([]Account, error) {

	var accounts []Account
	query := `
		SELECT ac.account_id, ac.user_id, ac.type, ac.currency, ac.account_number, ac.issuer, ac.dummy_col_3,
										ac.status, ac.created_at, ac.updated_at, ac.deleted_at,
										ad.account_id, ad.user_id, ad.color, ad.is_main_account, ad.progress, ad.account_nickname,
										ab.account_id, ab.amount,
										af.account_id, af.flag_id, af.flag_type, af.flag_value, af.created_at, af.updated_at
								FROM accounts ac
								LEFT JOIN account_details ad ON ac.account_id = ad.account_id
								LEFT JOIN account_balances ab ON ac.account_id = ab.account_id
								LEFT JOIN account_flags af ON ac.account_id = af.account_id
								WHERE ac.user_id = ?
								ORDER BY ac.created_at DESC
								LIMIT ? OFFSET ?`

	rows, err := r.db.Queryx(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	accountMap := make(map[string]Account)

	for rows.Next() {
		var account Account
		var accountDetail AccountDetail
		var accountBalance AccountBalances
		var accountFlag AccountFlag
		err = rows.Scan(
			&account.AccountID,
			&account.UserID,
			&account.Type,
			&account.Currency,
			&account.AccountNumber,
			&account.Issuer,
			&account.DummyCol3,
			&account.Status,
			&account.CreatedAt,
			&account.UpdatedAt,
			&account.DeletedAt,
			&accountDetail.AccountID,
			&accountDetail.UserID,
			&accountDetail.Color,
			&accountDetail.IsMainAccount,
			&accountDetail.Progress,
			&accountDetail.AccountNickname,
			&accountBalance.AccountID,
			&accountBalance.Amount,
			&accountFlag.AccountID,
			&accountFlag.FlagID,
			&accountFlag.FlagType,
			&accountFlag.FlagValue,
			&accountFlag.CreatedAt,
			&accountFlag.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		e, ok := accountMap[account.AccountID]
		if ok {
			e.AccountFlags = append(e.AccountFlags, accountFlag)
		} else {
			account.AccountDetail = accountDetail
			account.AccountBalances = accountBalance
			account.AccountFlags = []AccountFlag{accountFlag}
			accountMap[account.AccountID] = account
		}
	}

	for _, account := range accountMap {
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *repository) GetCountAccounts(userID string) (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM accounts WHERE user_id = ?", userID)
	return count, err
}
