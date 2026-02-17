package account

import (
	"testing"
	"time"
)

func TestAccountCreation(t *testing.T) {
	createdAt, _ := time.Parse(time.RFC3339, "2024-06-01T12:00:00Z")
	account := Account{
		ID:        "acc123",
		Name:      "Test Account",
		CreatedAt: createdAt,
	}
	expectAccount := Account{
		ID:        "acc123",
		Name:      "Test Account",
		CreatedAt: time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC),
	}

	if account.ID != expectAccount.ID {
		t.Errorf("Expected ID %s, got %s", expectAccount.ID, account.ID)
	}
	if account.Name != expectAccount.Name {
		t.Errorf("Expected Name %s, got %s", expectAccount.Name, account.Name)
	}
	if !account.CreatedAt.Equal(expectAccount.CreatedAt) {
		t.Errorf("Expected CreatedAt %s, got %s", expectAccount.CreatedAt, account.CreatedAt)
	}
}
