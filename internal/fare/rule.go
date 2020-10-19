package fare

import (
	"fmt"

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

type Rule interface {
	ApplicableAmount() (float64, error)
}

type rule struct {
	values []interface{}
}

func NewRules(values []interface{}) []Rule {
	rule := &rule{values: values}
	return []Rule{
		&oneZoneRule{rule},
		&anyOneZoneOutsideZoneOneRule{rule},
		&twoZonesIncZoneOneRule{rule},
		&twoZonesExcZoneOneRule{rule},
		&threeZonesRule{rule},
	}
}

type oneZoneRule struct {
	*rule
}

func (r *oneZoneRule) ApplicableAmount() (float64, error) {
	transportType, startZone, endZone, err := getTransportAndStationZones(r.values)
	if err != nil {
		return -1, err
	}
	if transportType == transport.TUBE && (totalZones(startZone, endZone) == countTwo && haveZoneOneAsStartAndEnd(startZone, endZone)) {
		return tubeFare - zoneOneFare, nil
	}
	return -1, nil
}

type anyOneZoneOutsideZoneOneRule struct {
	*rule
}

func (r *anyOneZoneOutsideZoneOneRule) ApplicableAmount() (float64, error) {
	transportType, startZone, endZone, err := getTransportAndStationZones(r.values)
	if err != nil {
		return -1, err
	}
	if transportType == transport.TUBE && (totalZones(startZone, endZone) == countTwo && !haveZoneOneAsStartAndEnd(startZone, endZone)) {
		return tubeFare - anyZoneOutsideZoneOneFare, nil
	}
	return -1, nil
}

type twoZonesIncZoneOneRule struct {
	*rule
}

func (r *twoZonesIncZoneOneRule) ApplicableAmount() (float64, error) {
	transportType, startZone, endZone, err := getTransportAndStationZones(r.values)
	if err != nil {
		return -1, err
	}
	if transportType == transport.TUBE && (totalZones(startZone, endZone) == countThree && haveZoneOneAsStartAndEnd(startZone, endZone)) {
		return tubeFare - twoZonesIncZoneOneFare, nil
	}
	return -1, nil
}

type twoZonesExcZoneOneRule struct {
	*rule
}

func (r *twoZonesExcZoneOneRule) ApplicableAmount() (float64, error) {
	transportType, startZone, endZone, err := getTransportAndStationZones(r.values)
	if err != nil {
		return -1, err
	}
	if transportType == transport.TUBE && (totalZones(startZone, endZone) == countThree && !haveZoneOneAsStartAndEnd(startZone, endZone)) {
		return tubeFare - twoZonesExcZoneOneFare, nil
	}
	return -1, nil
}

type threeZonesRule struct {
	*rule
}

func (r *threeZonesRule) ApplicableAmount() (float64, error) {
	transportType, startZone, endZone, err := getTransportAndStationZones(r.values)
	if err != nil {
		return -1, err
	}
	if transportType == transport.TUBE && totalZones(startZone, endZone) == countThree {
		return tubeFare - threeZonesFare, nil
	}
	return -1, nil
}

func getTransportAndStationZones(values []interface{}) (transport.Transport, station.Zone, station.Zone, error) {
	if len(values) != 3 {
		return "", nil, nil, fmt.Errorf("there are no sufficient values supplied to rule")
	}

	return values[0].(transport.Transport), values[1].(station.Zone), values[2].(station.Zone), nil
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
