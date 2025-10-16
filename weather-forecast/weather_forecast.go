//  Package weather can forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition the current condition.
var CurrentCondition string
// CurrentLocation the current location.
var CurrentLocation string

// Forecast returns an string value equal to the city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
