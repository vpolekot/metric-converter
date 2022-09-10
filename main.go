package metricConverter

import "fmt"

const CentimetersPerInch = 2.54
const MetersPerFoot = 0.3048
const LitersPerPint = 0.4732

func getInchToCentimeter(inches float64) {
	fmt.Printf("%v in. = %v cm", inches, inches*CentimetersPerInch)
}

func getLiterToPint(liters float64) {
	fmt.Printf("%v liters = %v pints", liters, liters/LitersPerPint)
}

func getFootToMeter(feets float64) {
	fmt.Printf("%v ft. = %v m", feets, feets*MetersPerFoot)
}
