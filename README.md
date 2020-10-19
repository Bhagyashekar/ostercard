# The Oyster Card Problem

Oster card system is a limited version of London’s Oyster card system. 
The card is loaded with £30.0 

## Operation

* **start_trip** ->starts the trip using given transport type and start point

* **end_trip** -> ends the trip at the given end point

## Stations and zones:

**Station Zone(s)**
Holborn 1
Earl’s Court 1, 2
Wimbledon 3
Hammersmith 2

## Fares:

**Journey Fare**
Anywhere in Zone 1 £2.50
Any one zone outside zone 1 £2.00
Any two zones including zone 1 £3.00
Any two zones excluding zone 1 £2.25
Any three zones £3.20
Any bus journey £1.80
The maximum possible fare is therefore £3.20.

## Running tests
```
make test
```

##Running oster card application

```
go run cmd/ostercard/main.go
```