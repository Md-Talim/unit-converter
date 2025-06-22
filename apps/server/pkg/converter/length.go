package converter

var LengthConverter = newLinearConverter("meters", map[string]float64{
	"millimeters": 0.001,
	"centimeters": 0.01,
	"meters":      1,
	"kilometers":  1000,
	"inches":      0.0254,
	"feet":        0.3048,
	"yards":       0.9144,
	"miles":       1609.34,
})
