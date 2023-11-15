package movmean

import (
	"fmt"
	"testing"
)

func TestMovmean(t *testing.T) {
	// https://www.mathworks.com/help/matlab/ref/movmean.html
	// A = [4 8 6 -1 -2 -3 -1 3 4 5];
	// M = movmean(A,3)
	// M = 1×10
	// 6.0 6.0 4.3 1.0 -2.0 -2.0 -0.3  2.0 4.0 4.5
	A := []float64{4.0, 8.0, 6.0, -1.0, -2.0, -3.0, -1.0, 3.0, 4.0, 5.0}
	k := 3

	result, _ := Movmean(A, k)
	fmt.Println(len(result))
	for _, res := range result {
		fmt.Printf("%.2f ", res)
	}
}
