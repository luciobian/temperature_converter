package converter

import (
	"fmt"
	"strings"
)

type Temperature struct {
	Value     float32
	System    string
	ConvertTo string
}

func CalculateConversion(temperature Temperature) float32 {
	switch strings.ToLower(temperature.System) {
	case "c":
		return CalculateFromCelsius(temperature)
	case "f":
		return CalculateFromFahrenheit(temperature)
	case "k":
		return CalculateFromKelvin(temperature)
	default:
		return 0
	}
}

func CalculateFromCelsius(temperature Temperature) float32 {
	if conversion, exists := CelsiusMap[temperature.ConvertTo]; exists {
		return conversion(temperature)
	}
	return 0
}

func CalculateFromFahrenheit(temperature Temperature) float32 {
	if conversion, exists := FahrenheitMap[temperature.ConvertTo]; exists {
		return conversion(temperature)
	}
	return 0
}

func CalculateFromKelvin(temperature Temperature) float32 {
	if conversion, exists := KelvinMap[temperature.ConvertTo]; exists {
		return conversion(temperature)
	}
	return 0
}

func PrintTemperatureTypes() {
	fmt.Println("C. Celsius (°C)")
	fmt.Println("F. Fahrenheit (°F)")
	fmt.Println("K. Kelvin (K)")
}
