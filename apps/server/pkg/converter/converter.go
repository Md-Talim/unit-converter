package converter

type UnitConverter interface {
	Convert(value float64, from string, to string) (float64, error)
}

type ConverterContext struct {
	strategy ConversionStrategy
}

func (c *ConverterContext) SetStrategy(s ConversionStrategy) {
	c.strategy = s
}

func (c *ConverterContext) Convert(value float64) float64 {
	return c.strategy.Convert(value)
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
