package main

import (
	"math"
	"strconv"

	"github.com/xuri/excelize/v2"
)

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
func make_singnal_xn(filename string) ([]float64, int) {
	f, _ := excelize.OpenFile(filename + ".xlsx")
	cells, _ := f.GetRows("main")
	a := make([]float64, len(cells))

	for i := 0; i < len(cells); i++ {
		b, _ := strconv.ParseFloat(cells[i][0], 64)
		a[i] = b
	}

	return a, len(cells)
}
func f(x float64, n int) float64 {
	var waveX float64
	waveX = math.Cos(2 * math.Pi * 0.03 * x)
	waveX += 0.8 * math.Cos(2*math.Pi*(0.06*x+0.09/(2.0*float64(n))*x*x))
	return waveX
}
