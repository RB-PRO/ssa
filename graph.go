package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Arafatk/glot"
	"github.com/xuri/excelize/v2"
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
	fmt.Println("makeGraphYX_float64_2")
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	plot.AddPointGroup(filename, "lines", [][]float64{y, x})
	plot.SavePlot("png" + OpSystemFilder + filename + ".png")
	return nil
}

// Построить график по координатам X и Y. Источник - mat.VecDense
func makeGraphYX_VecDense(x, y1, y2 mat.VecDense, f1, f2 string) error {
	x_arr := vecDense_in_float64(x)
	y1_arr := vecDense_in_float64(y1)
	y2_arr := vecDense_in_float64(y2)
	if len(x_arr) != len(y1_arr) {
		return errors.New("Length y1 of different for " + f1)
	}
	if len(x_arr) != len(y2_arr) {
		return errors.New("Length y2 of different for " + f2)
	}

	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	plot.AddPointGroup(f1, "lines", [][]float64{x_arr, y1_arr})
	plot.AddPointGroup(f2, "lines", [][]float64{x_arr, y2_arr})

	plot.SavePlot("png" + OpSystemFilder + f2 + ".png")
	return nil
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
	r1 := color{r: 0, g: 255, b: 255}
	r2 := color{r: 255, g: 255, b: 0}

	min_val := minDense(C)
	max_val := maxDense(C)
	delta := max_val - min_val

	// create xlsx
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")

	//err := file_graph.SetSheetProps("main", opts*SheetPropsOptions)

	r, c := C.Dims()
	file_graph.SetColWidth("main", getColumnName(1), getColumnName(c), 5)
	for i := 0; i < r; i++ {
		file_graph.SetRowHeight("main", i, 35)
		for j := 0; j < c; j++ {
			k := uint8(math.Round((C.At(i, j) - min_val) / delta * 100.0))
			colorOutput := colorSet(r1, r2, k)

			r := fmt.Sprintf("%x", colorOutput.r)
			g := fmt.Sprintf("%x", colorOutput.g)
			b := fmt.Sprintf("%x", colorOutput.b)

			style, err := file_graph.NewStyle(`{"fill":{"type":"pattern","color":["#` + r + g + b + `"],"pattern":1}}`)
			if err != nil {
				fmt.Println(err)
			}
			cell := getColumnName(j+1) + strconv.Itoa(i+1)
			file_graph.SetCellStyle("main", cell, cell, style)
		}
	}
	if err := file_graph.SetRowVisible("main", r, true); err != nil {
		fmt.Println(err)
	}
	if err := file_graph.SetColVisible("main", "A:"+getColumnName(c), true); err != nil {
		fmt.Println(err)
	}
	if err := file_graph.SaveAs("files" + OpSystemFilder + filename + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
}

// Узнать цвет градиента. k=[0;100]
func colorSet(r1, r2 color, k uint8) color {
	return color{r: r1.r*(1-k) + r2.r*k, g: r1.g*(1-k) + r2.g*k, b: r1.b*(1-k) + r2.b*k}
}

// минимальное значение элемента матрицы
func minDense(matr mat.Dense) float64 {
	var min float64 = matr.At(0, 0)
	r, c := matr.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if min > matr.At(i, j) {
				min = matr.At(i, j)
			}
		}
	}
	return min
}

// максимальное значение элемента матрицы
func maxDense(matr mat.Dense) float64 {
	var max float64 = matr.At(0, 0)
	r, c := matr.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if max < matr.At(i, j) {
				max = matr.At(i, j)
			}
		}
	}
	return max
}
