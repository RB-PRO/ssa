// Тестирование функции Pchip
//	PCHIP - Кусочный кубический интерполяционный полином Эрмита
// Результаты работы сохраняются в файл pchip_test.xlsx на лист golang
// Данные результаты необходимо сравнить с данными MatLab на листе MatLab

package pchip_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/RB-PRO/ssa/pkg/pchip"
	"github.com/xuri/excelize/v2"
)

func TestPchip2(t *testing.T) {
	// Согрузка входных данных
	win := 1024
	lag := int(math.Floor(float64(win) / 10.0))
	lgl := make([]float64, lag)
	for m := 0; m < len(lgl); m++ {
		lgl[m] = float64(m + 1)
	}
	PhaAcfNrm := loadXlsx("pchip_test.xlsx")

	// Использование функции
	coefC := pchip.Pchip2(lgl, PhaAcfNrm)
	fmt.Println("Len", coefC)

	safeToXlsx2(coefC)
}

// Сохранить в xlsx для дебага
func safeToXlsx2(coefs []float64) {
	file_graph, _ := excelize.OpenFile("pchip_test.xlsx", excelize.Options{})
	defer file_graph.Close()
	lenFor := len(coefs)
	for ind := 0; ind < lenFor; ind++ {
		//file_graph.SetCellValue("golang2", "B"+strconv.Itoa(ind+1), coefs.A[ind])
		//file_graph.SetCellValue("golang2", "C"+strconv.Itoa(ind+1), coefs.B[ind])
		file_graph.SetCellValue("golang2", "D"+strconv.Itoa(ind+1), coefs[ind])
		//file_graph.SetCellValue("golang2", "E"+strconv.Itoa(ind+1), coefs.D[ind])
	}
	file_graph.Save()
}
