package converter

import "fmt"

type TemperatureConverter struct{}

func (t *TemperatureConverter) Convert(value float64, from string, to string) (float64, error) {
	key := from + ":" + to
	strategy, ok := temperatureStrategies[key]
	if !ok {
		return 0, fmt.Errorf("conversion not supported")
	}

	return strategy.Convert(value), nil
}
