package converter

type UnitConverter interface {
	Convert(value float64, from string, to string) (float64, error)
}

type ConversionStrategy interface {
	Convert(value float64) float64
}

func GetConverter(category string) UnitConverter {
	switch category {
	case "temperature":
		return &TemperatureConverter{}
	// TODO: Implement weight and length
	default:
		return nil
	}
}
