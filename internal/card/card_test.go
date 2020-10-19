package card_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"

	"ostercard/internal/card"
	"ostercard/internal/station"
	"ostercard/internal/transport"
)

func TestSwipeToStartTrip(t *testing.T) {
	card := card.Card{Balance: 30.0}

	err := card.Swipe(transport.TUBE, station.HALBORN)

	assert.Equal(t, float64(26.8), card.Balance, "balance is not 22.9")
	assert.NoError(t, err)
}

func TestSwipeToEndTrip(t *testing.T) {
	card := card.Card{Balance: 30.0}
	err := card.Swipe(transport.TUBE, station.HALBORN)
	require.NoError(t, err)

	err = card.Swipe(transport.TUBE, station.EARLSCOURT)

	assert.NoError(t, err)
	assert.Equal(t, float64(26.8), card.Balance, "balance is not 22.9")
}
