// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF = 32.0
	const boilingF = 212.0

	fmt.Printf("%g°F = %g°C\n", freezingF, fToc(freezingF)) // 32°F = 0°C
	fmt.Printf("%g°F = %g°C\n", boilingF, fToc(boilingF))   // 212°F = 100°C
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
