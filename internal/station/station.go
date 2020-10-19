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

type StationZone map[StationName]Zone

var stn StationZone

func New() {
	stationZone := make(StationZone)
	stationZone[HALBORN] = []int{1}
	stationZone[EARLSCOURT] = []int{1, 2}
	stationZone[WIMBLENDON] = []int{3}
	stationZone[HAMMERSMITH] = []int{2}
	stn = stationZone
}

func GetZone(name StationName) Zone {
	return stn[name]
}
