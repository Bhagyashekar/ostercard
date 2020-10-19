package fare

import (
	"sort"

	"ostercard/internal/station"
	"ostercard/internal/transport"
)

type Calculator struct {
	Rules []Rule
}

func NewCalculator(transportType transport.Transport, startZone, endZone station.Zone) *Calculator {
	rules := NewRules([]interface{}{transportType, startZone, endZone})
	return &Calculator{Rules: rules}
}

func (c *Calculator) ChargeAfterJourney() (float64, error) {
	var applicableAmount []float64
	for _, rule := range c.Rules {
		amount, err := rule.ApplicableAmount()
		if err != nil {
			return -1, err
		}
		if amount != -1 {
			applicableAmount = append(applicableAmount, amount)
		}
	}

	sort.Float64s(applicableAmount)
	if len(applicableAmount) == 0 {
		return 0, nil
	}
	return applicableAmount[len(applicableAmount)-1], nil
}

func (c *Calculator) ChargeBeforeJourney(transportType transport.Transport) float64 {
	switch transportType {
	case transport.BUS:
		return busFare
	case transport.TUBE:
		return tubeFare
	}
	return -1
}
