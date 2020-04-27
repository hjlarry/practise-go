package main

import "fmt"

// 来自net包的例子，用于给一个无符号整数的最低5bit的每个bit指定一个名字
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
