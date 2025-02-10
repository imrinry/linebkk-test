package transactions

type TransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	IsBank        int    `json:"is_bank"`
}
