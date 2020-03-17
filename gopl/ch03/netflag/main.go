package main

import "fmt"

// Flags ...
type Flags uint

const (
	flagUp           Flags = 1 << iota // is up
	flagBroadcast                      // supports broadcast access capability
	flagLoopback                       // is a loopback interface
	flagPointToPoint                   // belongs to a point-to-point link
	flagMulticast                      // supports multicast access capability
)

func isUp(v Flags) bool {
	return v&flagUp == flagUp
}

func turnDown(v *Flags) {
	*v &^= flagUp
}

func setBroadcast(v *Flags) {
	*v |= flagBroadcast
}

func isCast(v Flags) bool {
	return v&(flagBroadcast|flagMulticast) != 0
}

func main() {
	var v = flagMulticast | flagUp
	fmt.Printf("%b %t\n", v, isUp(v))
	turnDown(&v)
	fmt.Printf("%b %t\n", v, isUp(v))
	setBroadcast(&v)
	fmt.Printf("%b %t\n", v, isUp(v))
	fmt.Printf("%b %t\n", v, isCast(v))
	fmt.Println(flagLoopback)
	fmt.Println(flagPointToPoint)
	fmt.Println(flagMulticast)
}
