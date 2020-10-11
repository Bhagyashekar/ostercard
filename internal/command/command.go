package command

import (
	"fmt"

	"ostercard/internal/card"
	"ostercard/internal/journey"
	"ostercard/internal/station"
	"ostercard/internal/transport"
)

const (
	startTrip = "start_trip"
	endTrip   = "end_trip"
)

type Command struct {
	name    string
	value   []string
	journey journey.Journey
	card    card.Card
}

func NewCommand(name string, value []string, journey journey.Journey, card card.Card) *Command {
	return &Command{
		name:    name,
		value:   value,
		journey: journey,
		card:    card,
	}
}

func (c *Command) Execute() error {
	switch c.name {
	case startTrip:
		if len(c.value) != 2 {
			return fmt.Errorf("command usage: start_trip transportType startPoint")
		}
		transportType := transport.Enum(c.value[0])
		if err := transport.Validate(transportType); err != nil {
			return err
		}
		stationName := station.Enum(c.value[1])
		if err := station.Validate(stationName); err != nil {
			return err
		}
		if err := c.journey.Start(transportType, stationName); err != nil {
			return err
		}
		fmt.Printf("Successfully started trip to %s in %s \n ", c.value[1], c.value[0])

	case endTrip:
		if len(c.value) != 1 {
			return fmt.Errorf("command usage: end_trip endPoint")
		}
		stationName := station.Enum(c.value[0])
		if err := station.Validate(stationName); err != nil {
			return err
		}
		if err := c.journey.End(stationName); err != nil {
			return err
		}
		balance := c.card.Balance()
		fmt.Printf("Successfully ended trip to %s \n", c.value[0])
		fmt.Printf("Card Balance after journey is %0.2f\n ", balance)

	default:
		return fmt.Errorf("invalid command")
	}
	return nil
}
