package fare

import (
	"ostercard/internal/card"
	"ostercard/internal/station"
	"ostercard/internal/transport"
)

const (
	zoneOneFare               = 2.50
	anyZoneOutsideZoneOneFare = 2.00
	twoZonesIncZoneOneFare    = 3.00
	twoZonesExcZoneOneFare    = 2.25
	threeZonesFare            = 3.20
	busFare                   = 1.80
	tubeFare                  = 3.20

	countTwo   = 2
	countThree = 3
)

type Fare interface {
	ChargeBeforeJourney(transportType transport.Transport) error
	ChargeAfterJourney(transportType transport.Transport, startPoint, endPoint station.StationName)
}

type fare struct {
	station station.Station
	card    card.Card
}

func New(station station.Station, card card.Card) Fare {
	return &fare{
		station: station,
		card:    card,
	}
}

func (f *fare) ChargeBeforeJourney(transportType transport.Transport) error {
	switch transportType {
	case transport.BUS:
		err := f.card.Debit(busFare)
		if err != nil {
			return err
		}
	case transport.TUBE:
		err := f.card.Debit(tubeFare)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *fare) ChargeAfterJourney(transportType transport.Transport, startPoint, endPoint station.StationName) {
	if transportType == transport.TUBE {
		startZone := f.station.Zone(startPoint)
		endZone := f.station.Zone(endPoint)

		if totalZones(startZone, endZone) == countTwo && !haveZoneOneAsStartAndEnd(startZone, endZone) {
			f.card.Credit(tubeFare - anyZoneOutsideZoneOneFare)
			return
		}
		if totalZones(startZone, endZone) == countTwo && haveZoneOneAsStartAndEnd(startZone, endZone) {
			f.card.Credit(tubeFare - zoneOneFare)
			return
		}
		if totalZones(startZone, endZone) == countThree && haveZoneOneAsStartAndEnd(startZone, endZone) {
			f.card.Credit(tubeFare - twoZonesIncZoneOneFare)
			return
		}
		if totalZones(startZone, endZone) == countThree && !haveZoneOneAsStartAndEnd(startZone, endZone) {
			f.card.Credit(tubeFare - twoZonesExcZoneOneFare)
			return
		}
		if totalZones(startZone, endZone) == countThree {
			f.card.Credit(tubeFare - threeZonesFare)
			return
		}
	}
}

func totalZones(startZone, endZone station.Zone) int {
	return len(startZone) + len(endZone)
}

func haveZoneOneAsStartAndEnd(startZone, endZone station.Zone) bool {
	return contains(startZone, 1) && contains(endZone, 1)
}

func contains(slice []int, value int) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}
