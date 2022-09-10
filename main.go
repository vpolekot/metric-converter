// Package metricConverter is a converter to-from metric system.
//
// Lightweight implementation.
package metricConverter

import "fmt"

const CentimetersPerInch = 2.54
const MetersPerFoot = 0.3048
const LitersPerPint = 0.4732

// Shows converted inches to centimeter without rounding.
func GetInchToCentimeter(inches float64) {
	fmt.Printf("%v in. = %v cm \n", inches, inches*CentimetersPerInch)
}

// Shows converted liters to pints without rounding.
func GetLiterToPint(liters float64) {
	fmt.Printf("%v liters = %v pints \n", liters, liters/LitersPerPint)
}

// Shows converted feets to meter without rounding.
func GetFootToMeter(feets float64) {
	fmt.Printf("%v ft. = %v m \n", feets, feets*MetersPerFoot)
}
