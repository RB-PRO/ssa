package pchip

import (
	"main/pkg/oss"
	"math"
	"testing"
)

func TestPchip(t *testing.T) {
	win := 1024
	lag := int(math.Floor(float64(win) / 10.0))
	lgl := make([]float64, lag)
	for m := 0; m < len(lgl); m++ {
		lgl[m] = float64(m + 1)
	}

	PhaAcfNrm, _ := oss.Make_singnal_xn("PhaAcfNrm") // Загрузить сигнал из файла PhaAcfNrm.xlsx

	_, pCoef := Pchip(PhaAcfNrm,
		lgl,
		lgl,
		len(PhaAcfNrm), len(lgl))

	oss.SafeToXlsxDualArray(pCoef, "pCoef")
	/*
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	*/
}
