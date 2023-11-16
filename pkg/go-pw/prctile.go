package gopw

import (
	"math"
	"sort"
)

// Расчёт процентилей
// Не работает ;)
func prctile(input []float64, percent float64) float64 {
	var percentile float64
	length := len(input)
	if length == 0 {
		return math.NaN()
	}

	if length == 1 {
		return input[0]
	}

	if percent <= 0 || percent > 100 {
		return math.NaN()
	}

	// Start by sorting a copy of the slice
	//c := sortedCopy(input)
	sort.Float64s(input)

	// Multiply percent by length of input
	index := (percent / 100) * float64(len(input))

	// Check if the index is a whole number
	if index == float64(int64(index)) {

		// Convert float to int
		i := int(index)

		// Find the value at the index
		percentile = input[i-1]

	} else if index > 1 {

		// Convert float to int via truncation
		i := int(index)

		// Find the average of the index and following values
		// percentile = Mean([]float64{input[i-1], input[i]}) // Mean(Float64Data{input[i-1], input[i]})
		percentile = (input[i-1] + input[i]) / 2

	} else {
		return math.NaN()
	}

	return percentile

}

// *****************************************************************************
