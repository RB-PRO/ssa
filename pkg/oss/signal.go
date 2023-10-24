package oss

import (
	"math"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Make_singnal(n int) []float64 {
	waveform := make([]float64, n)
	for index := range waveform {
		waveform[index] = F(float64(index), n)
		//waveform[index] += math.Cos(2 * math.Pi * 0.03 * float64(index))
		//waveform[index] += 0.8 * math.Cos(2*math.Pi*(0.06*float64(index)+0.09/(2.0*float64(n))*float64(index)*float64(index)))
	}
	return waveform
}
func Make_singnal_xn(filename string, Path string) ([]float64, int) {
	f, _ := excelize.OpenFile(Path + filename + ".xlsx")
	cells, _ := f.GetRows("main")
	a := make([]float64, len(cells))

	for i := 0; i < len(cells); i++ {
		b, _ := strconv.ParseFloat(cells[i][0], 64)
		a[i] = b
	}

	return a, len(cells)
}
func F(x float64, n int) float64 {
	var waveX float64
	waveX = math.Cos(2 * math.Pi * 0.03 * x)
	waveX += 0.8 * math.Cos(2*math.Pi*(0.06*x+0.09/(2.0*float64(n))*x*x))
	return waveX
}
