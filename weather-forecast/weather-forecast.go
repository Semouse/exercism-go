// Package weather provides tools to report
// current weather condition for current location.
package weather

// CurrentCondition represents current weather condition.
var CurrentCondition string

// CurrentLocation represents current city.
var CurrentLocation string

// Forecast returns a string value contains current weather condition for current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
