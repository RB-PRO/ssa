package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Arafatk/glot"
	"github.com/xuri/excelize/v2"
	"gonum.org/v1/gonum/mat"
)

const OpSystemFilder string = "/"

// ***
func safeToXlsx(sig []float64, name string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind, val := range sig {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), val)
	}
	if err := file_graph.SaveAs("files" + OpSystemFilder + name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
func safeToXlsxMatrix(X *mat.Dense, xlsxName string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	n, m := X.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			file_graph.SetCellValue("main", getColumnName(i+1)+strconv.Itoa(j+1), X.At(i, j))
		}
	}
	if err := file_graph.SaveAs("files" + OpSystemFilder + xlsxName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

/*
func make_singnal(n int) []float64 { //waveform := make_singnal(1000, 0.01)

		var lambda1 float64 = 0.025
		var lambda2 float64 = 0.01
		var lambda3 float64 = 0.1625
		var lambdaC float64 = 1.0 / 600.0
		var a1 float64 = 1.0
		var a2 float64 = 0.8
		var a3 float64 = 0.63
		waveform := make([]float64, n)
		for index := range waveform {
			waveform[index] += a1 * math.Sin(2*math.Pi*lambda1*float64(index))
			waveform[index] += a2 * math.Sin(2*math.Pi*lambda2*float64(index)+30.0*math.Sin(2.0*float64(math.Pi)*lambdaC*float64(index)))
			waveform[index] += a3 * math.Sin(2*math.Pi*lambda3*float64(index)+52.5*math.Sin(2.0*float64(math.Pi)*lambdaC*float64(index)))
		}
		return waveform
	}
*/
func make_singnal(n int) []float64 {
	waveform := make([]float64, n)
	for index := range waveform {
		waveform[index] = f(float64(index), n)
		//waveform[index] += math.Cos(2 * math.Pi * 0.03 * float64(index))
		//waveform[index] += 0.8 * math.Cos(2*math.Pi*(0.06*float64(index)+0.09/(2.0*float64(n))*float64(index)*float64(index)))
	}
	return waveform
}
func make_singnal_xn(n int) []float64 {
	f, _ := excelize.OpenFile("xn.xlsx")
	cells, _ := f.GetRows("main")
	a := make([]float64, len(cells))

	for i := 0; i < len(cells); i++ {
		b, _ := strconv.ParseFloat(cells[i][0], 64)
		a[i] = b
	}

	return a
}
func f(x float64, n int) float64 {
	var waveX float64
	waveX = math.Cos(2 * math.Pi * 0.03 * x)
	waveX += 0.8 * math.Cos(2*math.Pi*(0.06*x+0.09/(2.0*float64(n))*x*x))
	return waveX
}
func getColumnName(col int) string {
	name := make([]byte, 0, 3) // max 16,384 columns (2022)
	const aLen = 'Z' - 'A' + 1 // alphabet length
	for ; col > 0; col /= aLen + 1 {
		name = append(name, byte('A'+(col-1)%aLen))
	}
	for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
		name[i], name[j] = name[j], name[i]
	}
	return string(name)
}

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
	groupName := strings.Replace(filename, ".png", "", 1)
	groupName = strings.Replace(filename, "png/", "", 1)
	style := "lines"
	x := make([]float64, len(vals))
	for i := 0; i < len(vals); i++ {
		x[i] = float64(i)
	}
	plot.AddFunc2d(groupName, style, x, fct)
	plot.SavePlot(filename)
	return nil
}
func makeGraph2(n int, filename string) error {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	fct := func(x float64) float64 { return (f(x, n)) }
	groupName := strings.Replace(filename, ".png", "", 1)
	style := "lines"
	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = float64(i)
	}
	plot.AddFunc2d(groupName, style, x, fct)
	plot.SavePlot(filename)
	return nil
}

func aTa(matr *mat.Dense) *mat.Dense { // Multipy matrix AT*A
	a := mat.Matrix(matr)
	aT := a.T()
	ad := mat.DenseCopyOf(a)
	aTd := mat.DenseCopyOf(aT)
	n1, _ := aTd.Dims()
	_, m2 := ad.Dims()
	output := mat.NewDense(n1, m2, nil)
	output.Mul(aTd, ad)
	return output
}

func realyPrint(matr *mat.Dense, name string) {
	fmatr := mat.Formatted(matr, mat.Prefix(string(strings.Repeat(" ", 2+len(name)))), mat.Squeeze())
	fmt.Printf(name+" =%.3v\n", fmatr)
}
