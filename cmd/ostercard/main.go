package main

import (
	"log"

	"ostercard/internal/card"
	"ostercard/internal/station"
	"ostercard/internal/transport"
)

func main() {
	station.New()
	card := card.Card{Balance: 30.0}

	err := card.Swipe(transport.TUBE, station.HALBORN)
	if err != nil {
		log.Fatal(err)
	}
	err = card.Swipe(transport.TUBE, station.EARLSCOURT)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Balance after first journey is", card.Balance)

	err = card.Swipe(transport.BUS, station.EARLSCOURT)
	if err != nil {
		log.Fatal(err)
	}
	err = card.Swipe(transport.BUS, station.WIMBLENDON)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Balance after second journey is", card.Balance)

	err = card.Swipe(transport.TUBE, station.EARLSCOURT)
	if err != nil {
		log.Fatal(err)
	}
	err = card.Swipe(transport.TUBE, station.HAMMERSMITH)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Balance after third journey is", card.Balance)
}
