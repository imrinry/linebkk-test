package transactions

type Transaction struct {
	TransactionID string `db:"transaction_id"`
	UserID        string `db:"user_id"`
	Name          string `db:"name"`
	Image         string `db:"image"`
	IsBank        int    `db:"isBank"`
}

func (t *Transaction) ToTransactionResponse() *TransactionResponse {
	return &TransactionResponse{
		TransactionID: t.TransactionID,
		UserID:        t.UserID,
		Name:          t.Name,
		Image:         t.Image,
		IsBank:        t.IsBank,
	}
}
