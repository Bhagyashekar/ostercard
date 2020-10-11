package command_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ostercard/internal/command"
	"ostercard/internal/station"
	"ostercard/internal/testutil/mocks"
	"ostercard/internal/transport"
)

func TestStartTripCommand(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)
	journey.On("Start", transport.BUS, station.HALBORN).Return(nil)

	cmd := command.NewCommand("start_trip", []string{"Bus", "Holborn"}, journey, card)
	err := cmd.Execute()

	assert.NoError(t, err, "Unexpected start trip error")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestStartTripCommandForStartJourneyError(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)
	journey.On("Start", transport.BUS, station.HALBORN).Return(errors.New("some error"))

	cmd := command.NewCommand("start_trip", []string{"Bus", "Holborn"}, journey, card)
	err := cmd.Execute()

	assert.EqualError(t, err, "some error")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestStartTripCommandForMissingDetails(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)

	cmd := command.NewCommand("start_trip", []string{"Bus"}, journey, card)
	err := cmd.Execute()

	assert.EqualError(t, err, "command usage: start_trip transportType startPoint")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestStartTripCommandForInvalidTransportType(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)

	cmd := command.NewCommand("start_trip", []string{"invalid", "Holborn"}, journey, card)
	err := cmd.Execute()

	assert.EqualError(t, err, "transport type must be Bus and Tube")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestStartTripCommandForInvalidStationZone(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)

	cmd := command.NewCommand("start_trip", []string{"Bus", "invalid"}, journey, card)
	err := cmd.Execute()

	assert.EqualError(t, err, "station name must be Holborn, Earls_court, Wimblendon and Hammersmith")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestEndTripCommand(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)
	journey.On("End", station.EARLSCOURT).Return(nil)
	card.On("Balance").Return(float32(10.0))

	cmd := command.NewCommand("end_trip", []string{"Earls_court"}, journey, card)
	err := cmd.Execute()
	assert.NoError(t, err, "Unexpected end trip error")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestEndTripCommandForEndJourneyError(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)
	journey.On("End", station.EARLSCOURT).Return(errors.New("some error"))

	cmd := command.NewCommand("end_trip", []string{"Earls_court"}, journey, card)
	err := cmd.Execute()
	assert.EqualError(t, err, "some error")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestEndTripCommandForMissingDetails(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)

	cmd := command.NewCommand("end_trip", []string{}, journey, card)
	err := cmd.Execute()

	assert.EqualError(t, err, "command usage: end_trip endPoint")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestEndTripCommandForInvalidStationZone(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)

	cmd := command.NewCommand("end_trip", []string{"invalid"}, journey, card)
	err := cmd.Execute()

	assert.EqualError(t, err, "station name must be Holborn, Earls_court, Wimblendon and Hammersmith")
	mock.AssertExpectationsForObjects(t, card, journey)
}

func TestExecuteForInvalidCommand(t *testing.T) {
	card := new(mocks.Card)
	journey := new(mocks.Journey)


	cmd := command.NewCommand("invalid", []string{}, journey, card)
	err := cmd.Execute()

	assert.EqualError(t, err, "invalid command")
	mock.AssertExpectationsForObjects(t, card, journey)
}
