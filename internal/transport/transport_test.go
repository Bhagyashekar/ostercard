package transport_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ostercard/internal/transport"
)

var testTransportEnum = []struct {
	in  string
	out transport.Transport
}{
	{"Bus", transport.BUS},
	{"Tube", transport.TUBE},
	{"Unknown", ""},
}

func TestTransportEnum(t *testing.T) {
	for _, tt := range testTransportEnum {
		assert.Equal(t, tt.out, transport.Enum(tt.in), "Incorrect string representation")
	}
}

var testValidate = []struct {
	in  transport.Transport
	out error
}{
	{transport.BUS, nil},
	{transport.TUBE, nil},
	{"", errors.New("transport type must be Bus and Tube")},
}

func TestValidate(t *testing.T) {
	for _, tt := range testValidate {
		assert.Equal(t, tt.out, transport.Validate(tt.in), "transport validation failed")
	}
}
