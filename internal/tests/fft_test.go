package tests

import (
	"strconv"
	"testing"

	"github.com/mjibson/go-dsp/fft"
	"github.com/xuri/excelize/v2"
)

func TestFFT(t *testing.T) {
	x := loadXlsx("fft_test.xlsx")
	fft_signal := fft.FFTReal(x)
	safeToXlsx3("fft_test.xlsx", fft_signal)
}

// Получить входной сигнал из файла xlsx
func loadXlsx(filename string) []float64 {
	output := make([]float64, 0)

	file_graph, _ := excelize.OpenFile(filename, excelize.Options{})
	defer file_graph.Close()

	rows, err := file_graph.GetRows("input")
	if err != nil {
		return nil
	}

	for _, row := range rows {
		//fmt.Println(row[0])
		n, err := strconv.ParseFloat(row[0], 64)
		if err == nil {
			output = append(output, n)
		}
	}
	return output
}

// Сохранить в xlsx для дебага
func safeToXlsx3(filename string, y []complex128) {
	file_graph, _ := excelize.OpenFile(filename, excelize.Options{})
	defer file_graph.Close()
	lenFor := len(y)
	for ind := 0; ind < lenFor; ind++ {
		//file_graph.SetCellValue("golang2", "B"+strconv.Itoa(ind+1), coefs.A[ind])
		//file_graph.SetCellValue("golang2", "C"+strconv.Itoa(ind+1), coefs.B[ind])
		file_graph.SetCellValue("output", "A"+strconv.Itoa(ind+1), y[ind])
		file_graph.SetCellValue("output", "B"+strconv.Itoa(ind+1), real(y[ind]))
		file_graph.SetCellValue("output", "C"+strconv.Itoa(ind+1), imag(y[ind]))

		//file_graph.SetCellValue("golang2", "E"+strconv.Itoa(ind+1), coefs.D[ind])
	}
	file_graph.Save()
}
