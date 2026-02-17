package transaction

import "time"

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type Transaction struct {
	ID        string          `json:"id"`
	Type      TransactionType `json:"type"` // "credit" or "debit"
	Description string          `json:"description"`
	Amount    float64         `json:"amount"`
	CreatedAt time.Time       `json:"created_at"`
	CreatedBy string          `json:"created_by"` // Account ID that created the transaction
	UpdatedAt time.Time       `json:"updated_at"`
	UpdatedBy string          `json:"updated_by"` // Account ID that last updated the transaction
}