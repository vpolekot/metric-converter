package metricConverter

import "fmt"

const CentimetersPerInch = 2.54
const MetersPerFoot = 0.3048
const LitersPerPint = 0.4732

func GetInchToCentimeter(inches float64) {
	fmt.Printf("%v in. = %v cm", inches, inches*CentimetersPerInch)
}

func GetLiterToPint(liters float64) {
	fmt.Printf("%v liters = %v pints", liters, liters/LitersPerPint)
}

func GetFootToMeter(feets float64) {
	fmt.Printf("%v ft. = %v m", feets, feets*MetersPerFoot)
}
