package debit_cards

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DebitCardRepository interface {
	GetDebitCardByUserID(ctx context.Context, userID string, offset int, limit int) ([]DebitCard, error)
	GetCountDebitCards(ctx context.Context, userID string) (int, error)
}

type debitCardRepository struct {
	db *sqlx.DB
}

func NewDebitCardRepository(db *sqlx.DB) *debitCardRepository {
	return &debitCardRepository{db: db}
}

func (r *debitCardRepository) GetDebitCardByUserID(ctx context.Context, userID string, offset int, limit int) ([]DebitCard, error) {
	fmt.Println("userID", userID)
	fmt.Println("offset", offset)
	fmt.Println("limit", limit)

	var debitCards []DebitCard
	query := `
		SELECT 
			 	dc.card_id, dc.user_id, dc.card_type, dc.issue_at, dc.expired_at,
				dcd.card_id, dcd.user_id, dcd.issuer, dcd.number,
				dcs.card_id, dcs.user_id, dcs.status, dcs.blocked_reason,
				dcdes.card_id, dcdes.user_id, dcdes.color, dcdes.border_color
		FROM 
			debit_cards dc
		LEFT JOIN debit_card_details dcd ON dc.card_id = dcd.card_id
		LEFT JOIN debit_card_status dcs ON dc.card_id = dcs.card_id
		LEFT JOIN debit_card_design dcdes ON dc.card_id = dcdes.card_id
		WHERE 
			dc.user_id = ?
		ORDER BY dc.issue_at DESC
		LIMIT ? OFFSET ?
	`
	rows, err := r.db.Queryx(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var debitCard DebitCard
		var debitCardDetails DebitCardDetails
		var debitCardStatus DebitCardStatus
		var debitCardDesign DebitCardDesign

		err = rows.Scan(&debitCard.CardID, &debitCard.UserID, &debitCard.CardType, &debitCard.IssueAt, &debitCard.ExpiredAt,
			&debitCardDetails.CardID, &debitCardDetails.UserID, &debitCardDetails.Issuer, &debitCardDetails.Number,
			&debitCardStatus.CardID, &debitCardStatus.UserID, &debitCardStatus.Status, &debitCardStatus.BlockedReason,
			&debitCardDesign.CardID, &debitCardDesign.UserID, &debitCardDesign.Color, &debitCardDesign.BorderColor,
		)
		if err != nil {
			return nil, err
		}
		debitCard.DebitCardDetails = debitCardDetails
		debitCard.DebitCardStatus = debitCardStatus
		debitCard.DebitCardDesign = debitCardDesign
		debitCards = append(debitCards, debitCard)
	}

	fmt.Println("debitCards repository", debitCards)

	return debitCards, nil
}

func (r *debitCardRepository) GetCountDebitCards(ctx context.Context, userID string) (int, error) {
	query := `
		SELECT COUNT(*) FROM debit_cards WHERE user_id = ?
	`
	var count int
	err := r.db.GetContext(ctx, &count, query, userID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
