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

## Running Application

Oster card can be run in 2 different modes

#### Interactive mode

Allows the user to perform operations interactively 

```
./ostercard.sh
```

Example
```                     
Welcome to the oster card
start_trip Tube Holborn
Successfully started trip to Holborn in Tube 
end_trip Earls_court
Successfully ended trip to Earls_court 
Card Balance after journey is 27.00
start_trip Bus Earls_court
Successfully started trip to Earls_court in Bus 
end_trip Wimblendon
Successfully ended trip to Wimblendon 
Card Balance after journey is 25.20
start_trip Tube Earls_court
Successfully started trip to Earls_court in Tube 
end_trip Hammersmith
Successfully ended trip to Hammersmith 
Card Balance after journey is 22.95
```

#### Input from File mode

User can store all the operations in a file and can specify the file name as an argument to the CLI
```
./ostercard.sh input/file_inputs.txt
```
 
 Example
 ```
start_trip Tube Holborn
end_trip Earls_court
start_trip Bus Earls_court
end_trip Wimblendon
start_trip Tube Earls_court
end_trip Hammersmith
Successfully started trip to Holborn in Tube 
Successfully ended trip to Earls_court 
Card Balance after journey is 27.00
Successfully started trip to Earls_court in Bus 
Successfully ended trip to Wimblendon 
Card Balance after journey is 25.20
Successfully started trip to Earls_court in Tube 
Successfully ended trip to Hammersmith 
Card Balance after journey is 22.95
```
 

