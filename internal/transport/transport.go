package transport

import "fmt"

type Transport string

const (
	BUS  Transport = "Bus"
	TUBE           = "Tube"
)

func Enum(transport string) Transport {
	switch transport {
	case "Bus":
		return BUS
	case "Tube":
		return TUBE
	default:
		return ""
	}
}

func Validate(transport Transport) error {
	if transport != BUS && transport != TUBE {
		return fmt.Errorf("transport type must be Bus and Tube")
	}
	return nil
}
