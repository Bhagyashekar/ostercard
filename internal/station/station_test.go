package station_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ostercard/internal/station"
)

var testStationNameEnum = []struct {
	in  string
	out station.StationName
}{
	{"Holborn", station.HALBORN},
	{"Earls_court", station.EARLSCOURT},
	{"Wimblendon", station.WIMBLENDON},
	{"Hammersmith", station.HAMMERSMITH},
	{"Unknown", -1},
}

func TestStationNameEnum(t *testing.T) {
	for _, tt := range testStationNameEnum {
		assert.Equal(t, tt.out, station.Enum(tt.in), "Incorrect string representation")
	}
}

var testValidate = []struct {
	in  station.StationName
	out error
}{
	{station.HALBORN, nil},
	{station.EARLSCOURT, nil},
	{station.WIMBLENDON, nil},
	{station.HAMMERSMITH, nil},
	{-1, errors.New("station name must be Holborn, Earls_court, Wimblendon and Hammersmith")},
}

func TestValidate(t *testing.T) {
	for _, tt := range testValidate {
		assert.Equal(t, tt.out, station.Validate(tt.in), "station name validation failed")
	}
}

