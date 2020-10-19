package journey_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ostercard/internal/journey"
	"ostercard/internal/station"
	"ostercard/internal/transport"
)

func TestStart(t *testing.T) {
	journey := &journey.Journey{}
	err := journey.Start(transport.BUS, station.HALBORN)

	assert.NoError(t, err, "expected no error while starting the journey")
}

func TestStartWhenTripAlreadyInProgress(t *testing.T) {
	journey := &journey.Journey{}
	err := journey.Start(transport.BUS, station.HALBORN)
	require.NoError(t, err, "expected no error while starting the journey")

	err = journey.Start(transport.BUS, station.HALBORN)
	assert.EqualError(t, err, "already in a trip complete it before starting the other trip")
}

func TestEnd(t *testing.T) {
	journey := &journey.Journey{}
	err := journey.Start(transport.BUS, station.HALBORN)
	require.NoError(t, err, "expected no error while starting the journey")

	err = journey.End(station.EARLSCOURT)
	assert.NoError(t, err, "expected no error while ending the journey")
}

func TestEndWhenTripAlreadyNotInProgress(t *testing.T) {
	journey := &journey.Journey{}
	err := journey.End(station.EARLSCOURT)

	assert.EqualError(t, err, "first start the trip, before ending it")
}
