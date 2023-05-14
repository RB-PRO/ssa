package graph_test

import (
	"testing"

	"github.com/RB-PRO/ssa/pkg/graph"
)

func Test2D_plot(t *testing.T) {

	x := []float64{0.0, 1.0, 2.0, 3.0, 4.0}
	y := []float64{0.0, 4.0, 2.0, 1.0, 3.0}

	err := graph.MakeGraphYX_float64(x, y, "2d")
	if err != nil {
		t.Error(err)
	}

}
