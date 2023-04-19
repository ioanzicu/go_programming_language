package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// *celsiusFlag satisfies the flag.Value interface
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

/*
go run main.go
20°C

go run main.go -temp -15C
-15°C

go run main.go -temp 212°F
100°C

go run main.go -temp 212°K
invalid value "212°K" for flag -temp: invalid temperature "212°K"
Usage of /tmp/go-build3945960859/b001/exe/main:
  -temp value
        the temperature (default 20°C)
exit status 2

go run main.go -temp 273.15K
invalid value "273.15K" for flag -temp: invalid temperature "273.15K"
Usage of /tmp/go-build778600315/b001/exe/main:
  -temp value
        the temperature (default 20°C)
exit status 2
*/

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
