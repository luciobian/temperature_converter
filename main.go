package main

import (
	"fmt"
	"strings"
	"temperature_converter/converter"
)

func main() {
	repeat := "y"

	for strings.ToLower(repeat) == "y" {
		fmt.Println("Welcome to the Temperature Converter")

		fmt.Println("Please enter the temperature you want to convert:")
		var value float32
		fmt.Scanln(&value)

		fmt.Println("Select the input temperature type:")
		converter.PrintTemperatureTypes()

		var system string
		fmt.Scanln(&system)

		fmt.Println("Select the type of temperature you want to convert to:")
		converter.PrintTemperatureTypes()

		var convertTo string
		fmt.Scanln(&convertTo)

		temperature := converter.Temperature{Value: value, System: system, ConvertTo: convertTo}

		newValue := converter.CalculateConversion(temperature)
		fmt.Println("----------------------------------------------------------------------------------------------------------")
		fmt.Printf("For the entered temperature of %.2f °%s, the result of the conversion to the %s system is %.2f °%s. \n",
			temperature.Value, strings.ToUpper(temperature.System), converter.MapTemperatureName[temperature.ConvertTo], newValue, strings.ToUpper(temperature.ConvertTo))
		fmt.Println("----------------------------------------------------------------------------------------------------------")

		fmt.Print("Do you want to make another conversion? (y/n)")
		fmt.Scanln(&repeat)
	}

	fmt.Println("Thank you for using the Temperature Converter! Goodbye!")
}
