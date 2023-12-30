package converter

var MapTemperatureName = map[string]string{
	"c": "Celsius",
	"f": "Fahrenheit",
	"k": "Kelvin",
}

var CelsiusMap = map[string]func(temp Temperature) float32{
	"c": func(temp Temperature) float32 {
		return temp.Value
	},
	"f": func(temp Temperature) float32 {
		return (temp.Value * 9 / 5) + 32
	},
	"k": func(temp Temperature) float32 {
		return temp.Value + 273.15
	},
}

var FahrenheitMap = map[string]func(temp Temperature) float32{
	"c": func(temp Temperature) float32 {
		return 5 / 9 * (temp.Value - 32)
	},
	"f": func(temp Temperature) float32 {
		return temp.Value
	},
	"k": func(temp Temperature) float32 {
		return (temp.Value + 459.67) / 1.8
	},
}

var KelvinMap = map[string]func(temp Temperature) float32{
	"c": func(temp Temperature) float32 {
		return temp.Value - 273.15
	},
	"f": func(temp Temperature) float32 {
		return 9/5*(temp.Value-273.15) + 32
	},
	"k": func(temp Temperature) float32 {
		return temp.Value
	},
}
