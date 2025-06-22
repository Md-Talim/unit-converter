package converter

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func TestTemperatureConverter(t *testing.T) {
	converter := &TemperatureConverter{}

	tests := []struct {
		value    float64
		from     string
		to       string
		expected float64
	}{
		{0, "celsius", "fahrenheit", 32},
		{100, "celsius", "fahrenheit", 212},
		{0, "celsius", "kelvin", 273.15},
		{273.15, "kelvin", "celsius", 0},
		{32, "fahrenheit", "celsius", 0},
		{212, "fahrenheit", "celsius", 100},
		{32, "fahrenheit", "kelvin", 273.15},
		{212, "fahrenheit", "kelvin", 373.15},
		{273.15, "kelvin", "fahrenheit", 32},
		{373.15, "kelvin", "fahrenheit", 212},
	}

	for _, test := range tests {
		result, err := converter.Convert(test.value, test.from, test.to)
		if err != nil {
			t.Errorf("unexpected error for %v %s to %s: %v", test.value, test.from, test.to, err)
		}
		if !floatEquals(result, test.expected) {
			t.Errorf("Convert(%v, %s, %s) = %v; want %v", test.value, test.from, test.to, result, test.expected)
		}
	}
}

func TestTemperatureConverterUnsupported(t *testing.T) {
	converter := &TemperatureConverter{}
	_, err := converter.Convert(100, "celsius", "rankine")
	if err == nil {
		t.Error("expected error for unsupported conversion, got nil")
	}
}
