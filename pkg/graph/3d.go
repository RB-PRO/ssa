package graph

import "github.com/Arafatk/glot"

func TreeXDXD(x, y, z []float64) {
	// plot, _ := glot.NewPlot(dimensions, persist, debug)
	dimensions := 3
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	fct := func(x, y float64) float64 { return x - y }
	groupName := "Stright Line"
	style := "lines"
	pointsY := []float64{1, 2, 3, 4, 5}
	pointsX := []float64{1, 2, 3, 4, 5}
	plot.AddFunc3d(groupName, style, pointsX, pointsY, fct)
	plot.SetXrange(0, 5)
	plot.SetYrange(0, 5)
	plot.SetZrange(0, 5)
	plot.SavePlot("1.png")
}
