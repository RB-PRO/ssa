package movmean

import (
	"fmt"
	"math"
)

func Movmean(data []float64, window int) ([]float64, error) {
	if window <= 1 {
		return nil, fmt.Errorf("length of slice must be greater than 1")
	}
	if window > len(data) {
		return nil, fmt.Errorf("m cannot be greater than length of slice")
	}
	if window%2 == 0 {
		window++
	}

	n := len(data)
	M := make([]float64, n)

	for i := 0; i < n; i++ {
		sum := 0.0
		count := 0

		for j := i - window/2; j < i+window/2+1; j++ {
			if j >= 0 && j < len(data) {
				// fmt.Println(data[j])
				sum += data[j]
				count++
			}
		}

		// fmt.Println(i, "<>", i-window/2, ":", i+window/2, "//", sum, float64(count))
		// fmt.Println()

		// for j := i - window + 1; j <= i; j++ {
		// 	if j >= 0 && j < n {
		// 		sum += data[j]
		// 		count++
		// 	}
		// }

		M[i] = sum / float64(count)
	}

	return M, nil
}
func MovMeanStd(ts []float64, m int) ([]float64, []float64, error) {
	if m <= 1 {
		return nil, nil, fmt.Errorf("length of slice must be greater than 1")
	}

	if m > len(ts) {
		return nil, nil, fmt.Errorf("m cannot be greater than length of slice")
	}

	var i int

	c := make([]float64, len(ts)+1)
	csqr := make([]float64, len(ts)+1)
	for i = 0; i < len(ts)+1; i++ {
		if i == 0 {
			c[i] = 0
			csqr[i] = 0
		} else {
			c[i] = ts[i-1] + c[i-1]
			csqr[i] = ts[i-1]*ts[i-1] + csqr[i-1]
		}
	}

	mean := make([]float64, len(ts)-m+1)
	std := make([]float64, len(ts)-m+1)
	for i = 0; i < len(ts)-m+1; i++ {
		mean[i] = (c[i+m] - c[i]) / float64(m)
		std[i] = math.Sqrt((csqr[i+m]-csqr[i])/float64(m) - mean[i]*mean[i])
	}

	return mean, std, nil
}
