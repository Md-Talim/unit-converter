package converter

import "testing"

func TestLengthConverter(t *testing.T) {
	tests := []struct {
		value    float64
		from     string
		to       string
		expected float64
	}{
		{1000, "millimeters", "meters", 1},
		{1, "meters", "millimeters", 1000},
		{100, "centimeters", "meters", 1},
		{1, "meters", "centimeters", 100},
		{1, "kilometers", "meters", 1000},
		{1000, "meters", "kilometers", 1},
		{1, "inches", "meters", 0.0254},
		{1, "meters", "inches", 39.3700787},
		{1, "feet", "meters", 0.3048},
		{1, "meters", "feet", 3.2808399},
		{1, "yards", "meters", 0.9144},
		{1, "meters", "yards", 1.0936133},
		{1, "miles", "meters", 1609.34},
		{1609.34, "meters", "miles", 1},
	}

	for _, test := range tests {
		result, err := LengthConverter.Convert(test.value, test.from, test.to)
		if err != nil {
			t.Errorf("unexpected error for %v %s to %s: %v", test.value, test.from, test.to, err)
		}
		if !floatEquals(result, test.expected) {
			t.Errorf("Convert(%v, %s, %s) = %v; want %v", test.value, test.from, test.to, result, test.expected)
		}
	}
}

func TestLengthConverterUnsupported(t *testing.T) {
	_, err := LengthConverter.Convert(1, "meters", "parsecs")
	if err == nil {
		t.Error("expected error for unsupported conversion, got nil")
	}
}
