package fare_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ostercard/internal/fare"
	stationtype "ostercard/internal/station"
	"ostercard/internal/testutil/mocks"
	"ostercard/internal/transport"
)

func TestChargeBeforeJourneyForBusTransport(t *testing.T) {
	card := new(mocks.Card)
	station := new(mocks.Station)
	card.On("Debit", float32(1.80)).Return(nil)

	fare := fare.New(station, card)
	err := fare.ChargeBeforeJourney(transport.BUS)
	assert.NoError(t, err, "Unexpected error while charging before journey for bus")
	mock.AssertExpectationsForObjects(t, card, station)
}

func TestChargeBeforeJourneyForTubeTransport(t *testing.T) {
	card := new(mocks.Card)
	station := new(mocks.Station)
	card.On("Debit", float32(3.2)).Return(nil)

	fare := fare.New(station, card)
	err := fare.ChargeBeforeJourney(transport.TUBE)
	assert.NoError(t, err, "Unexpected error while charging before journey for tube")
	mock.AssertExpectationsForObjects(t, card, station)
}

func TestChargeBeforeJourneyForDebitError(t *testing.T) {
	card := new(mocks.Card)
	station := new(mocks.Station)
	card.On("Debit", float32(3.2)).Return(errors.New("insufficient balance"))

	fare := fare.New(station, card)
	err := fare.ChargeBeforeJourney(transport.TUBE)
	assert.EqualError(t, err, "insufficient balance")
	mock.AssertExpectationsForObjects(t, card, station)
}

func TestChargeAfterJourneyForBusTransport(t *testing.T) {
	card := new(mocks.Card)
	station := new(mocks.Station)

	fare := fare.New(station, card)
	fare.ChargeAfterJourney(transport.BUS, stationtype.HALBORN, stationtype.EARLSCOURT)
	mock.AssertExpectationsForObjects(t, card, station)
}

func TestChargeAfterJourneyForTubeTransportForZoneOne(t *testing.T) {
	card := new(mocks.Card)
	station := new(mocks.Station)
	station.On("Zone", stationtype.HALBORN).Return(stationtype.Zone{1})
	station.On("Zone", stationtype.EARLSCOURT).Return(stationtype.Zone{1})
	card.On("Credit", float32(0.7)).Return()

	fare := fare.New(station, card)
	fare.ChargeAfterJourney(transport.TUBE, stationtype.HALBORN, stationtype.EARLSCOURT)
	mock.AssertExpectationsForObjects(t, card, station)
}

func TestChargeAfterJourneyForTubeTransportForAnyZoneOutsideZoneOne(t *testing.T) {
	card := new(mocks.Card)
	station := new(mocks.Station)
	station.On("Zone", stationtype.HALBORN).Return(stationtype.Zone{2})
	station.On("Zone", stationtype.EARLSCOURT).Return(stationtype.Zone{2})
	card.On("Credit", float32(1.2)).Return()

	fare := fare.New(station, card)
	fare.ChargeAfterJourney(transport.TUBE, stationtype.HALBORN, stationtype.EARLSCOURT)
	mock.AssertExpectationsForObjects(t, card, station)
}

func TestChargeAfterJourneyForTubeTransportForTwoZonesIncZoneOne(t *testing.T) {
	card := new(mocks.Card)
	station := new(mocks.Station)
	station.On("Zone", stationtype.HALBORN).Return(stationtype.Zone{1})
	station.On("Zone", stationtype.EARLSCOURT).Return(stationtype.Zone{1, 2})
	card.On("Credit", float32(0.2)).Return()

	fare := fare.New(station, card)
	fare.ChargeAfterJourney(transport.TUBE, stationtype.HALBORN, stationtype.EARLSCOURT)
	mock.AssertExpectationsForObjects(t, card, station)
}

func TestChargeAfterJourneyForTubeTransportForTwoZonesExcZoneOne(t *testing.T) {
	card := new(mocks.Card)
	station := new(mocks.Station)
	station.On("Zone", stationtype.HALBORN).Return(stationtype.Zone{2})
	station.On("Zone", stationtype.EARLSCOURT).Return(stationtype.Zone{3, 2})
	card.On("Credit", float32(0.95)).Return()

	fare := fare.New(station, card)
	fare.ChargeAfterJourney(transport.TUBE, stationtype.HALBORN, stationtype.EARLSCOURT)
	mock.AssertExpectationsForObjects(t, card, station)
}
