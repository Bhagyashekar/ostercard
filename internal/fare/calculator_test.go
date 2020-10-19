package fare_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ostercard/internal/fare"
	"ostercard/internal/station"
	"ostercard/internal/transport"
)

func TestChargeBeforeJourneyForBusTransport(t *testing.T) {
	fare := fare.NewCalculator(transport.BUS, station.Zone{1}, station.Zone{1, 2})
	amount := fare.ChargeBeforeJourney(transport.BUS)

	assert.Equal(t, 1.8, amount)
}

func TestChargeBeforeJourneyForTubeTransport(t *testing.T) {
	fare := fare.NewCalculator(transport.TUBE, station.Zone{1}, station.Zone{1, 2})
	amount := fare.ChargeBeforeJourney(transport.TUBE)

	assert.Equal(t, 3.2, amount)
}

func TestChargeAfterJourneyForBusTransport(t *testing.T) {
	fare := fare.NewCalculator(transport.BUS, station.Zone{1}, station.Zone{1, 2})
	amount, err := fare.ChargeAfterJourney()

	assert.Zero(t, amount)
	assert.NoError(t, err)
}

func TestChargeAfterJourneyForTubeTransportForZoneOne(t *testing.T) {
	fare := fare.NewCalculator(transport.TUBE, station.Zone{1}, station.Zone{1})
	amount, err := fare.ChargeAfterJourney()

	assert.Equal(t, 0.7, amount)
	assert.NoError(t, err)
}

func TestChargeAfterJourneyForTubeTransportForAnyZoneOutsideZoneOne(t *testing.T) {
	fare := fare.NewCalculator(transport.TUBE, station.Zone{2}, station.Zone{2})
	amount, err := fare.ChargeAfterJourney()

	assert.Equal(t, 1.2, amount)
	assert.NoError(t, err)
}


func TestChargeAfterJourneyForTubeTransportForTwoZonesIncZoneOne(t *testing.T) {
	fare := fare.NewCalculator(transport.TUBE, station.Zone{1}, station.Zone{1, 2})
	amount, err := fare.ChargeAfterJourney()

	assert.Equal(t, 0.2, amount)
	assert.NoError(t, err)
}

func TestChargeAfterJourneyForTubeTransportForTwoZonesExcZoneOne(t *testing.T) {
	fare := fare.NewCalculator(transport.TUBE, station.Zone{2}, station.Zone{ 2, 3})
	amount, err := fare.ChargeAfterJourney()

	assert.Equal(t, 0.95, amount)
	assert.NoError(t, err)
}