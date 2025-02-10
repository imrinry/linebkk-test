package debit_cards

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type DebitCardRepository interface {
	GetDebitCardByUserID(ctx context.Context, userID string, offset int, limit int) ([]DebitCard, error)
	GetCountDebitCards(ctx context.Context, userID string) (int, error)
	SetDebitCardCache(ctx context.Context, key string, value []DebitCardResponse, expiration time.Duration) error
	GetDebitCardCache(ctx context.Context, key string) ([]DebitCardResponse, error)
}

type debitCardRepository struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewDebitCardRepository(db *sqlx.DB, redis *redis.Client) *debitCardRepository {
	return &debitCardRepository{db: db, redis: redis}
}

func (r *debitCardRepository) GetDebitCardByUserID(ctx context.Context, userID string, offset int, limit int) ([]DebitCard, error) {

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

func (r *debitCardRepository) SetDebitCardCache(ctx context.Context, key string, value []DebitCardResponse, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, key, jsonValue, expiration).Err()
}

func (r *debitCardRepository) GetDebitCardCache(ctx context.Context, key string) ([]DebitCardResponse, error) {
	value, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var debitCards []DebitCardResponse
	err = json.Unmarshal([]byte(value), &debitCards)
	if err != nil {
		return nil, err
	}
	return debitCards, nil
}
