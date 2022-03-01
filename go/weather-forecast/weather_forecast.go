// Package weather provides tools to announce forecasts.
package weather

// CurrentCondition stores the weather condition in the location where the forecast is performed.
var CurrentCondition string

// CurrentLocation stores the name of the city where the forecast is performed.
var CurrentLocation string

// Forecast returns a human readable sentence announcing the weather condition in the specified location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
