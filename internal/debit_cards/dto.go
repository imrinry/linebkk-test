package debit_cards

import "time"

type DebitCardResponse struct {
	CardID           string                   `json:"card_id"`
	UserID           string                   `json:"user_id"`
	Name             string                   `json:"name"`
	CardType         string                   `json:"card_type"`
	IssueAt          time.Time                `json:"issue_at"`
	ExpiredAt        *time.Time               `json:"expired_at"`
	DebitCardDetails DebitCardDetailsResponse `json:"debit_card_details"`
	DebitCardStatus  DebitCardStatusResponse  `json:"debit_card_status"`
	DebitCardDesign  DebitCardDesignResponse  `json:"debit_card_design"`
}

type DebitCardDetailsResponse struct {
	CardID string `json:"card_id"`
	UserID string `json:"user_id"`
	Issuer string `json:"issuer"`
	Number string `json:"number"`
}

type DebitCardStatusResponse struct {
	CardID        string  `json:"card_id"`
	UserID        string  `json:"user_id"`
	Status        string  `json:"status"`
	BlockedReason *string `json:"blocked_reason"`
}

type DebitCardDesignResponse struct {
	CardID      string  `json:"card_id"`
	UserID      string  `json:"user_id"`
	Color       *string `json:"color"`
	BorderColor *string `json:"border_color"`
}
