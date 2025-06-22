package converter

import "math"

// keys returns the keys of a map as a slice of strings.
func keys(m map[string]float64) []string {
	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// floatEquals compares two float64 values for approximate equality.
func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}
