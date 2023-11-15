package tests_test

import (
	"testing"

	gomathtests "github.com/RB-PRO/ssa/pkg/go-MathTests"
)

func TestLoadSave(t *testing.T) {
	Err := gomathtests.Save("save.txt", []float64{1, 2, 3, 4, 5})
	if Err != nil {
		t.Error(Err)
	}
}
func TestPlot(t *testing.T) {
	Err := gomathtests.Plot("save.png", []float64{1, 2, 2, 7, 5})
	if Err != nil {
		t.Error(Err)
	}
}
