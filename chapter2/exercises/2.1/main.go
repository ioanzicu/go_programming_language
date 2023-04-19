// Add types, constants, and functions to tempconv for processing temperatures in
// the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magni-
// tude as 1°C.
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = -273.15
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF-FreezingC) // invalid operation: boilingF - FreezingC (mismatched types Fahrenheit and Celsius)

	var c Celsius
	var f Fahrenheit
	var k Kelvin
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) //  invalid operation: c == f (mismatched types Celsius and Fahrenheit)
	fmt.Println(c == Celsius(f))

	c = FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))

	k = CToK(c)
	fmt.Println("Celsius", c, "Kelvins", k)

	k = CToK(0)
	fmt.Println("Celsius", 0, "Kelvins", k)
}
