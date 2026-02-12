// Package weather is a package for predicting weather.
package weather

var (
	// CurrentCondition is a variable for storing the current condition.
	CurrentCondition string
	// CurrentLocation is a variable for storing the current location.
	CurrentLocation  string
)

// Forecast returns the current location and the current conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
