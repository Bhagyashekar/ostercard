package journey_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"ostercard/internal/journey"
	"ostercard/internal/station"
	"ostercard/internal/testutil/mocks"
	"ostercard/internal/transport"
)

func TestStart(t *testing.T) {
	fare := new(mocks.Fare)
	fare.On("ChargeBeforeJourney", transport.BUS).Return(nil)

	journey := journey.New(fare)
	err := journey.Start(transport.BUS, station.HALBORN)
	assert.NoError(t, err, "expected no error while starting the journey")
	mock.AssertExpectationsForObjects(t, fare)
}

func TestStartWhwnError(t *testing.T) {
	fare := new(mocks.Fare)
	fare.On("ChargeBeforeJourney", transport.BUS).Return(errors.New("some error"))

	journey := journey.New(fare)
	err := journey.Start(transport.BUS, station.HALBORN)
	assert.EqualError(t, err, "some error")
	mock.AssertExpectationsForObjects(t, fare)
}

func TestStartWhenTripAlreadyInProgress(t *testing.T) {
	fare := new(mocks.Fare)
	fare.On("ChargeBeforeJourney", transport.BUS).Return(nil)

	journey := journey.New(fare)
	err := journey.Start(transport.BUS, station.HALBORN)
	require.NoError(t, err, "expected no error while starting the journey")

	err = journey.Start(transport.BUS, station.HALBORN)
	assert.EqualError(t, err, "already in a trip complete it before starting the other trip")
	mock.AssertExpectationsForObjects(t, fare)
}

func TestEnd(t *testing.T) {
	fare := new(mocks.Fare)
	fare.On("ChargeBeforeJourney", transport.BUS).Return(nil)
	fare.On("ChargeAfterJourney", transport.BUS, station.HALBORN, station.EARLSCOURT).Return(nil)

	journey := journey.New(fare)
	err := journey.Start(transport.BUS, station.HALBORN)
	require.NoError(t, err, "expected no error while starting the journey")
	err = journey.End(station.EARLSCOURT)
	assert.NoError(t, err, "expected no error while ending the journey")
	mock.AssertExpectationsForObjects(t, fare)
}

func TestEndWhenTripAlreadyNotInProgress(t *testing.T) {
	fare := new(mocks.Fare)

	journey := journey.New(fare)
	err := journey.End(station.EARLSCOURT)
	assert.EqualError(t, err, "first start the trip, before ending it")
	mock.AssertExpectationsForObjects(t, fare)
}
