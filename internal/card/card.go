package card

import "fmt"

type card struct {
	balance float32
}

type Card interface {
	Balance() float32
	HasSufficientBalance(fare float32) bool
	Credit(fare float32)
	Debit(fare float32) error
}

func New(balance float32) Card {
	return &card{balance: balance}
}

func (c *card) Balance() float32 {
	return c.balance
}

func (c *card) HasSufficientBalance(fare float32) bool {
	if c.balance < fare {
		return false
	}
	return true
}

func (c *card) Credit(fare float32) {
	c.balance = c.balance + fare
}

func (c *card) Debit(fare float32) error {
	hasSufficientBalance := c.HasSufficientBalance(fare)
	if !hasSufficientBalance {
		return fmt.Errorf("insufficient balance")
	}
	c.balance = c.balance - fare
	return nil
}
