package card

import (
	"fmt"

	"ostercard/internal/fare"
	"ostercard/internal/journey"
	"ostercard/internal/station"
	"ostercard/internal/transport"
)

type Card struct {
	Balance float64
	Journey journey.Journey
}

func (c *Card) Swipe(transport transport.Transport, stationName station.StationName) error {
	if c.Journey.InTrip {
		err := c.Journey.End(stationName)
		if err != nil {
			return err
		}

		amount, err := fare.NewCalculator(c.Journey.Transport, station.GetZone(c.Journey.StartPoint),
			station.GetZone(c.Journey.EndPoint)).ChargeAfterJourney()
		if err != nil {
			return err
		}
		c.Balance = c.Balance + amount
		return nil
	}

	err := c.Journey.Start(transport, stationName)
	if err != nil {
		return err
	}
	amount := fare.NewCalculator(c.Journey.Transport,
		station.GetZone(c.Journey.StartPoint),
		station.GetZone(c.Journey.EndPoint)).ChargeBeforeJourney(transport)
	hasSufficientBalance := c.hasSufficientBalance(amount)
	if !hasSufficientBalance {
		return fmt.Errorf("insufficient balance")
	}
	c.Balance = c.Balance - amount
	return nil
}

func (c *Card) hasSufficientBalance(fare float64) bool {
	if c.Balance < fare {
		return false
	}
	return true
}
