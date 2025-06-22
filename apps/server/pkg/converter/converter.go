package converter

import "fmt"

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
	case "weight":
		return &WeightConverter
	default:
		return nil
	}
}

type LinearConverter struct {
	baseUnit string
	factors  map[string]float64
}

func newLinearConverter(base string, factors map[string]float64) LinearConverter {
	return LinearConverter{baseUnit: base, factors: factors}
}

func (l *LinearConverter) Convert(value float64, from string, to string) (float64, error) {
	fromFactor, fromOk := l.factors[from]
	toFactor, toOk := l.factors[to]
	if !fromOk {
		return 0, fmt.Errorf("unrecognized 'from' unit: %s. Supported units: %v", from, keys(l.factors))
	}
	if !toOk {
		return 0, fmt.Errorf("unrecognized 'to' unit: %s. Supported units: %v", to, keys(l.factors))
	}

	// Normalize to base unit, then to target
	baseValue := value * fromFactor
	targetValue := baseValue / toFactor

	return targetValue, nil
}
