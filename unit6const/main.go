package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func main() {
	// fmt.Printf("Friday %v\n", Friday)
	fmt.Printf("Flags %v\n", FlagUp)
	fmt.Printf("FlagBroadcast2 %v\n", FlagBroadcast)
	fmt.Printf("FlagLoopback %v\n", FlagLoopback)
	fmt.Printf("FlagPointToPoint %v\n", FlagPointToPoint)
	fmt.Printf("FlagMulticast %v\n", FlagMulticast)

}
