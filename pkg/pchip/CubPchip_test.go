// Тестирование функции Pchip
//	PCHIP - Кусочный кубический интерполяционный полином Эрмита
// Результаты работы сохраняются в файл pchip_test.xlsx на лист golang
// Данные результаты необходимо сравнить с данными MatLab на листе MatLab

package pchip_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/RB-PRO/ssa/pkg/pchip"
	"github.com/xuri/excelize/v2"
)

func TestPchip(t *testing.T) {
	// Согрузка входных данных
	win := 1024
	lag := int(math.Floor(float64(win) / 10.0))
	lgl := make([]float64, lag)
	for m := 0; m < len(lgl); m++ {
		lgl[m] = float64(m + 1)
	}
	PhaAcfNrm := loadXlsx("pchip_test.xlsx")

	// Использование функции
	_, _, coefs := pchip.Pchip(PhaAcfNrm,
		lgl,
		lgl,
		len(PhaAcfNrm), len(lgl))

	safeToXlsx(coefs)
}

// Сохранить в xlsx для дебага
func safeToXlsx(coefs pchip.PchipCoefs) {
	file_graph, _ := excelize.OpenFile("pchip_test.xlsx", excelize.Options{})
	defer file_graph.Close()
	lenFor := len(coefs.A)
	for ind := 0; ind < lenFor; ind++ {
		file_graph.SetCellValue("golang", "B"+strconv.Itoa(ind+1), coefs.A[ind])
		file_graph.SetCellValue("golang", "C"+strconv.Itoa(ind+1), coefs.B[ind])
		file_graph.SetCellValue("golang", "D"+strconv.Itoa(ind+1), coefs.C[ind])
		file_graph.SetCellValue("golang", "E"+strconv.Itoa(ind+1), coefs.D[ind])
	}
	file_graph.Save()
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
