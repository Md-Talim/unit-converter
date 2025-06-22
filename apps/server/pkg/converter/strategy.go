package converter

type ConversionStrategy interface {
	Convert(value float64) float64
}

type CelsiusToFahrenheit struct{}

func (c *CelsiusToFahrenheit) Convert(value float64) float64 {
	return value*9/5 + 32
}

type CelsiusToKelvin struct{}

func (c *CelsiusToKelvin) Convert(value float64) float64 {
	return value + 273.15
}

type FahrenheitToCelsius struct{}

func (f *FahrenheitToCelsius) Convert(value float64) float64 {
	return (value - 32) * 5 / 9
}

type KelvinToCelsius struct{}

func (k *KelvinToCelsius) Convert(value float64) float64 {
	return value - 273.15
}

type FahrenheitToKelvin struct{}

func (f *FahrenheitToKelvin) Convert(value float64) float64 {
	return (value-32)*5/9 + 273.15
}

type KelvinToFahrenheit struct{}

func (k *KelvinToFahrenheit) Convert(value float64) float64 {
	return (value-273.15)*9/5 + 32
}

var temperatureStrategies = map[string]ConversionStrategy{
	"celsius:fahrenheit": &CelsiusToFahrenheit{},
	"celsius:kelvin":     &CelsiusToKelvin{},
	"fahrenheit:celsius": &FahrenheitToCelsius{},
	"fahrenheit:kelvin":  &FahrenheitToKelvin{},
	"kelvin:celsius":     &KelvinToCelsius{},
	"kelvin:fahrenheit":  &KelvinToFahrenheit{},
}
