// Package weather provides functionality related to weather, including calls to a weather API
// and various weather-related manipulation functions.
package weather

func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
