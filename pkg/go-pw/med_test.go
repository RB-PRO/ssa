package gopw

import (
	"fmt"
	"testing"
)

func TestMedFilt1(t *testing.T) {
	fmt.Println(medianFilter1([]float64{7, 3, 6, 4, 2}, 3))
	// 3 6 4 4 2
}
