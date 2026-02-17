package transaction

import (
	"testing"
	"time"
)

func TestTransactionCreation(t *testing.T) {
	createdAt, _ := time.Parse(time.RFC3339, "2024-06-01T12:00:00Z")
	transaction := Transaction{
		ID:          "txn123",
		Type:        Credit,
		Description: "Test Transaction",
		Amount:      100.0,
		CreatedAt:   createdAt,
		CreatedBy:   "acc123",
		UpdatedAt:   createdAt,
		UpdatedBy:   "acc123",
	}
	expectTransaction := Transaction{
		ID:          "txn123",
		Type:        Credit,
		Description: "Test Transaction",
		Amount:      100.0,
		CreatedAt:   time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC),
		CreatedBy:   "acc123",
		UpdatedAt:   time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC),
		UpdatedBy:   "acc123",
	}

	if transaction.ID != expectTransaction.ID {
		t.Errorf("Expected ID %s, got %s", expectTransaction.ID, transaction.ID)
	}
	if transaction.Type != expectTransaction.Type {
		t.Errorf("Expected Type %s, got %s", expectTransaction.Type, transaction.Type)
	}
	if transaction.Description != expectTransaction.Description {
		t.Errorf("Expected Description %s, got %s", expectTransaction.Description, transaction.Description)
	}
	if transaction.Amount != expectTransaction.Amount {
		t.Errorf("Expected Amount %f, got %f", expectTransaction.Amount, transaction.Amount)
	}
	if !transaction.CreatedAt.Equal(expectTransaction.CreatedAt) {
		t.Errorf("Expected CreatedAt %s, got %s", expectTransaction.CreatedAt, transaction.CreatedAt)
	}
	if transaction.CreatedBy != expectTransaction.CreatedBy {
		t.Errorf("Expected CreatedBy %s, got %s", expectTransaction.CreatedBy, transaction.CreatedBy)
	}
	if !transaction.UpdatedAt.Equal(expectTransaction.UpdatedAt) {
		t.Errorf("Expected UpdatedAt %s, got %s", expectTransaction.UpdatedAt, transaction.UpdatedAt)
	}
	if transaction.UpdatedBy != expectTransaction.UpdatedBy {
		t.Errorf("Expected UpdatedBy %s, got %s", expectTransaction.UpdatedBy, transaction.UpdatedBy)
	}
}