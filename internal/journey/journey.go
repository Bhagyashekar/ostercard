package journey

import (
	"fmt"

	"ostercard/internal/station"
	"ostercard/internal/transport"
)

type Journey struct {
	Transport  transport.Transport
	StartPoint station.StationName
	EndPoint   station.StationName
	InTrip     bool
}

func (j *Journey) Start(transport transport.Transport, startPoint station.StationName) error {
	if j.InTrip == true {
		return fmt.Errorf("already in a trip complete it before starting the other trip")
	}
	j.Transport = transport
	j.StartPoint = startPoint
	j.InTrip = true
	return nil
}

func (j *Journey) End(endPoint station.StationName) error {
	if j.InTrip != true {
		return fmt.Errorf("first start the trip, before ending it")
	}
	j.EndPoint = endPoint
	j.InTrip = false
	return nil
}
