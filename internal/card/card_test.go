package card_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ostercard/internal/card"
)

func TestBalance(t *testing.T) {
	card := card.New(20.0)

	balance := card.Balance()

	assert.Equal(t, float32(20.0), balance, "balance is not 20.0")
}

func TestHasSufficientBalance(t *testing.T) {
	card := card.New(20.0)

	hasSufficientBalance := card.HasSufficientBalance(20.0)

	assert.True(t, hasSufficientBalance, "has sufficient balance is not true")
}

func TestHasSufficientBalanceWhenNoSufficientBalance(t *testing.T) {
	card := card.New(20.0)

	hasSufficientBalance := card.HasSufficientBalance(20.1)

	assert.False(t, hasSufficientBalance, "has sufficient balance is not false")
}

func TestCredit(t *testing.T) {
	card := card.New(20.0)

	card.Credit(2.9)
	balance := card.Balance()

	assert.Equal(t, float32(22.9), balance, "balance is not 22.9")
}

func TestDebit(t *testing.T) {
	card := card.New(20.0)

	err := card.Debit(2.9)

	assert.NoError(t, err, "expected no error while debit")
	balance := card.Balance()
	assert.Equal(t, float32(17.1), balance, "balance is not 17.1")
}

func TestDebitWhenInsufficientBalance(t *testing.T) {
	card := card.New(20.0)

	err := card.Debit(22.9)

	assert.EqualError(t, err, "insufficient balance")
}
