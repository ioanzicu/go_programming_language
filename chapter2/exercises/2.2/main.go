// Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguments or from the standard input if there are no argu-
// ments, and converts each number into units like temperature in Celsius and Fahrenheit,
// length in feet and meters, weight in pounds and kilograms, and the like.
package main

import "fmt"

type Kilogram float64

func (k Kilogram) String() string {
	return fmt.Sprintf("%g Kgs", k)
}

type Pound float64

func (p Pound) String() string {
	return fmt.Sprintf("%g Lbs", p)
}

const (
	OneLbsInKg = 0.45359237
	OneKgInLbs = 2.2046
)

func KgsToLbs(k Kilogram) Pound {
	return Pound(k * OneKgInLbs)
}

func LbsToKgs(p Pound) Kilogram {
	return Kilogram(p * OneLbsInKg)
}

func main() {
	k := KgsToLbs(100)
	fmt.Println("100 Kgs to Lbs = ", k.String())

	l := LbsToKgs(100)
	fmt.Println("100 Lbs to Kgs = ", l.String())
}
