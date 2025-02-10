package debit_cards

import "time"

type DebitCard struct {
	CardID           string           `db:"card_id"`
	UserID           string           `db:"user_id"`
	Name             string           `db:"name"`
	CardType         string           `db:"card_type"`
	IssueAt          time.Time        `db:"issue_at"`
	ExpiredAt        *time.Time       `db:"expired_at"`
	DebitCardDetails DebitCardDetails `db:"debit_card_details"`
	DebitCardStatus  DebitCardStatus  `db:"debit_card_status"`
	DebitCardDesign  DebitCardDesign  `db:"debit_card_design"`
}

type DebitCardDetails struct {
	CardID string `db:"card_id"`
	UserID string `db:"user_id"`
	Issuer string `db:"issuer"`
	Number string `db:"number"`
}

type DebitCardStatus struct {
	CardID        string  `db:"card_id"`
	UserID        string  `db:"user_id"`
	Status        string  `db:"status"`
	BlockedReason *string `db:"blocked_reason"`
}

type DebitCardDesign struct {
	CardID      string  `db:"card_id"`
	UserID      string  `db:"user_id"`
	Color       *string `db:"color"`
	BorderColor *string `db:"border_color"`
}

func (d *DebitCard) ToDebitCardResponse() DebitCardResponse {

	return DebitCardResponse{
		CardID:           d.CardID,
		UserID:           d.UserID,
		Name:             d.Name,
		CardType:         d.CardType,
		IssueAt:          d.IssueAt,
		ExpiredAt:        d.ExpiredAt,
		DebitCardDetails: d.DebitCardDetails.ToDebitCardDetailsResponse(),
		DebitCardStatus:  d.DebitCardStatus.ToDebitCardStatusResponse(),
		DebitCardDesign:  d.DebitCardDesign.ToDebitCardDesignResponse(),
	}
}

func (d *DebitCardDetails) ToDebitCardDetailsResponse() DebitCardDetailsResponse {
	return DebitCardDetailsResponse{
		CardID: d.CardID,
		UserID: d.UserID,
		Issuer: d.Issuer,
		Number: d.Number,
	}
}

func (d *DebitCardStatus) ToDebitCardStatusResponse() DebitCardStatusResponse {
	return DebitCardStatusResponse{
		CardID:        d.CardID,
		UserID:        d.UserID,
		Status:        d.Status,
		BlockedReason: d.BlockedReason,
	}
}

func (d *DebitCardDesign) ToDebitCardDesignResponse() DebitCardDesignResponse {
	return DebitCardDesignResponse{
		CardID:      d.CardID,
		UserID:      d.UserID,
		Color:       d.Color,
		BorderColor: d.BorderColor,
	}
}
