package graph_test

import (
	"testing"

	"github.com/Arafatk/glot"
	"github.com/RB-PRO/ssa/pkg/graph"
)

func Test2D_plot(t *testing.T) {
	x := []float64{0.0, 1.0, 2.0, 3.0, 4.0}
	y := []float64{0.0, 4.0, 2.0, 1.0, 3.0}
	err := graph.MakeGraphYX_float64(x, y, "", "2d.png")
	if err != nil {
		t.Error(err)
	}
}
func Test2DD(t *testing.T) {
	// Создаем данные для графика
	x := []float64{0.0, 1.0, 2.0, 3.0, 4.0}
	y := []float64{0.0, 1.0, 4.0, 9.0, 16.0}

	// Создаем новый график
	dimensions := 2
	persist := true
	debug := true
	plot, err := glot.NewPlot(dimensions, persist, debug)
	if err != nil {
		panic(err)
	}

	// Добавляем точки на график
	err = plot.AddPointGroup("SimplePlot", "lines", [][]float64{x, y})
	if err != nil {
		panic(err)
	}

	// Сохраняем график в файл
	err = plot.SavePlot("D:\\Desktop\\Work\\program\\go\\src\\github.com\\RB-PRO\\ssa\\pkg\\graph\\2.png")
	if err != nil {
		panic(err)
	}

}

func Test2ddd(t *testing.T) {
	dimensions := 3
	persist := false
	debug := true
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	plot.AddPointGroup("Sample 1", "lines", []float64{2, 3, 4, 1})
	plot.SetTitle("Test Results")
	plot.SetZrange(-2, 2)
	plot.SavePlot("1.png")
}
