package station

import "fmt"

type Zone []int

type StationName int

const (
	HALBORN StationName = iota
	EARLSCOURT
	WIMBLENDON
	HAMMERSMITH
)

func Enum(stationName string) StationName {
	switch stationName {
	case "Holborn":
		return HALBORN
	case "Earls_court":
		return EARLSCOURT
	case "Wimblendon":
		return WIMBLENDON
	case "Hammersmith":
		return HAMMERSMITH
	default:
		return -1
	}
}

func Validate(name StationName) error {
	if name != HALBORN && name != EARLSCOURT && name != WIMBLENDON && name != HAMMERSMITH {
		return fmt.Errorf("station name must be Holborn, Earls_court, Wimblendon and Hammersmith")
	}
	return nil
}

type Station interface {
	Zone(name StationName) Zone
}

type StationZone map[StationName]Zone

type station struct {
	StationZone
}

func New() Station {
	stationZone := make(StationZone)
	stationZone[HALBORN] = []int{1}
	stationZone[EARLSCOURT] = []int{1, 2}
	stationZone[WIMBLENDON] = []int{3}
	stationZone[HAMMERSMITH] = []int{2}
	return &station{stationZone}
}

func (s *station) Zone(name StationName) Zone {
	return s.StationZone[name]
}
