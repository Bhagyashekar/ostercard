package journey

import (
	"fmt"

	"ostercard/internal/fare"
	"ostercard/internal/station"
	"ostercard/internal/transport"
)

type Journey interface {
	Start(transport transport.Transport, startPoint station.StationName) error
	End(endPoint station.StationName) error
}

type journey struct {
	transport  transport.Transport
	startPoint station.StationName
	endPoint   station.StationName
	inTrip     bool
	fare       fare.Fare
}

func New(fare fare.Fare) Journey {
	return &journey{
		fare: fare,
	}
}

func (j *journey) Start(transport transport.Transport, startPoint station.StationName) error {
	if j.inTrip == true {
		return fmt.Errorf("already in a trip complete it before starting the other trip")
	}
	j.transport = transport
	j.startPoint = startPoint
	j.inTrip = true
	if err := j.fare.ChargeBeforeJourney(j.transport); err != nil {
		return err
	}
	return nil
}

func (j *journey) End(endPoint station.StationName) error {
	if j.inTrip != true {
		return fmt.Errorf("first start the trip, before ending it")
	}
	j.endPoint = endPoint
	j.fare.ChargeAfterJourney(j.transport, j.startPoint, j.endPoint)
	j.inTrip = false
	return nil
}
