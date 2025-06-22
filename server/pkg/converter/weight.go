package converter

var WeightConverter = newLinearConverter("grams", map[string]float64{
	"milligrams": 0.001,
	"grams":      1,
	"kilograms":  1000,
	"pounds":     453.59237,
	"ounces":     28.3495231,
})
