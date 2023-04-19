// Explain why the help message contains °C when the default value of 20.0 does not.

/*
ANSWER:

Beacuse the help message prints the value using method func (c Celsius) String() string

%g	Exponent as needed, only necessary digits
so 20.0 is printed as 20°C
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f°K", k)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = -273.15
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273)
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
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
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

go run main.go -temp 303.15K
30°C

go run main.go -temp 273.15°K
0°C
*/

var stdout io.Writer = os.Stdout

func main() {
	flag.Parse()
	fmt.Fprintln(stdout, *temp)
}
