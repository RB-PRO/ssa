package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Arafatk/glot"
	"gonum.org/v1/gonum/mat"
)

type color struct {
	r, g, b uint8
}

// "png"+OpSystemFilder+

func makeGraph(n int, dt float64, pointsX []float64, filename string) error {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	fct := func(x float64) float64 { return (f(x, n)) }
	groupName := "Исходная функция"
	style := "lines"
	for i := range pointsX {
		pointsX[i] = dt
		dt += 0.1
	}
	plot.AddFunc2d(groupName, style, pointsX, fct)
	plot.SavePlot(filename)
	return nil
}
func makeGraphOfArray(vals []float64, filename string) error {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	fct := func(x float64) float64 { return (vals[int(x)]) }
	groupName := strings.Replace("png"+OpSystemFilder+filename+".png", ".png", "", 1)
	groupName = strings.Replace("png"+OpSystemFilder+filename+".png", "png/", "", 1)
	style := "lines"
	x := make([]float64, len(vals))
	for i := 0; i < len(vals); i++ {
		x[i] = float64(i)
	}
	plot.AddFunc2d(groupName, style, x, fct)
	plot.SavePlot("png" + OpSystemFilder + filename + ".png")
	return nil
}

// Построить график по координатам X и Y. Источник - float64[]
func makeGraphYX_float64(x, y []float64, filename string) error {
	if len(x) != len(y) {
		return errors.New("Length different for " + filename)
	}
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	plot.AddPointGroup(filename, "lines", [][]float64{y, x})
	plot.SavePlot("png" + OpSystemFilder + filename + ".png")
	return nil
}

// Построить график по координатам X и Y. Источник - mat.VecDense
func makeGraphYX_VecDense(x, y mat.VecDense, filename string) error {
	x_arr := vecDense_in_float64(x)
	y_arr := vecDense_in_float64(y)
	if len(x_arr) != len(y_arr) {
		return errors.New("Length of different for " + filename)
	}
	return makeGraphYX_float64(x_arr, y_arr, filename)
}

func vecDense_in_float64(vec mat.VecDense) []float64 {
	leng, _ := vec.Dims()
	output := make([]float64, leng)
	for ind := range output {
		output[ind] = vec.AtVec(ind)
	}
	return output
}

func imagesc(C mat.Dense, filename string) {
	r1 := color{r: 0, g: 0, b: 255}
	r2 := color{r: 255, g: 255, b: 0}

	colorOutput := colorSet(r1, r2, 2)

	fmt.Println(colorOutput)

}

// Узнать цвет градиента
func colorSet(r1, r2 color, k uint8) color {
	return color{r: r1.r*(1-k) + r2.r*k, g: r1.g*(1-k) + r2.g*k, b: r1.b*(1-k) + r2.b*k}
}
