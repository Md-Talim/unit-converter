package converter

import (
	"testing"
)

func TestWeightConverter(t *testing.T) {
	tests := []struct {
		value    float64
		from     string
		to       string
		expected float64
	}{
		{1000, "grams", "kilograms", 1},
		{1, "kilograms", "grams", 1000},
		{1000, "milligrams", "grams", 1},
		{1, "grams", "milligrams", 1000},
		{1, "kilograms", "pounds", 2.20462262},
		{2.20462262, "pounds", "kilograms", 1},
		{1, "grams", "ounces", 0.03527396},
		{1, "ounces", "grams", 28.3495231},
		{1, "pounds", "ounces", 16},
		{16, "ounces", "pounds", 1},
	}

	for _, test := range tests {
		result, err := WeightConverter.Convert(test.value, test.from, test.to)
		if err != nil {
			t.Errorf("unexpected error for %v %s to %s: %v", test.value, test.from, test.to, err)
		}
		if !floatEquals(result, test.expected) {
			t.Errorf("Convert(%v, %s, %s) = %v; want %v", test.value, test.from, test.to, result, test.expected)
		}
	}
}

func TestWeightConverterUnsupported(t *testing.T) {
	_, err := WeightConverter.Convert(100, "grams", "stones")
	if err == nil {
		t.Error("expected error for unsupported conversion, got nil")
	}
}
