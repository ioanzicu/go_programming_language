package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopbacl interfacce
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	fmt.Printf("FlagUp\t\t\t= %b\n", FlagUp)                 // "1
	fmt.Printf("FlagBroadcast\t\t= %b\n", FlagBroadcast)     // "10
	fmt.Printf("FlagLoopback\t\t= %b\n", FlagLoopback)       // "100
	fmt.Printf("FlagPointToPoint\t= %b\n", FlagPointToPoint) // "1000
	fmt.Printf("FlagMulticast\t\t= %b\n", FlagMulticast)     // "10000

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 false"

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
