// Package weather provides tools to
// get weather conditions.
package weather

var (
	// CurrentCondition represents a weather condition.
	CurrentCondition string
	// CurrentLocation represents a city location.
	CurrentLocation string
)

// Forecast returns a string representing the weather condition in a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
