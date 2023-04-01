package tempconv

import "fmt"

// package tempconv

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (c Fahrenheit) String() string {
	return fmt.Sprintf("%g°C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)
