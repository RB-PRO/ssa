package periodogram_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/RB-PRO/ssa/pkg/periodogram"
	"github.com/xuri/excelize/v2"
)

func Test1Periodogram_32(t *testing.T) {
	fmt.Println("--- Periodogram: Test 1 ---")
	N := 32                      // к-во отсчётов
	x32, w32, rez32 := Data32(N) // Загружаем данные для длины окна 32

	priodogman32 := periodogram.Periodogram(x32, w32, N) // Расчёт периодограммы

	var ErrorCount int

	// Циклом идём по всему массиву и сравниваем данные
	for index := range priodogman32 {
		if int(math.Abs(100000*(priodogman32[index]-rez32[index]))) > 3 {
			ErrorCount++
			// t.Errorf("Blackmanharris: Элемент с индексом %v не соответствует рассчитанному. Должно было быть %v, а получено %v.",index, rez32[index], priodogman32[index])
		}
	}
	safeToXlsx(priodogman32)
	// Если к-во ошибок не равно нулю
	if ErrorCount != 0 {
		t.Errorf("Blackmanharris: Test1: Расхождений между результатами MatLab и полученными результатами - " + strconv.Itoa(ErrorCount) + " из " + strconv.Itoa(N))
	}
}

func Test1Periodogram_1024(t *testing.T) {
	fmt.Println("--- Periodogram: Test 2 ---")
	N := 1024                            // к-во отсчётов
	x1024, w1024, rez1024 := Data1024(N) // Загружаем данные для длины окна 1024

	priodogman1024 := periodogram.Periodogram(x1024, w1024, N) // Расчёт периодограммы

	var ErrorCount int

	// Циклом идём по всему массиву и сравниваем данные
	for index := range priodogman1024 {
		if int(math.Abs(10*(priodogman1024[index]-rez1024[index]))) > 3 {
			ErrorCount++
			//t.Errorf("Blackmanharris: Элемент с индексом %v не соответствует рассчитанному. Должно было быть %v, а получено %v.", index, rez1024[index], priodogman1024[index])
		}
	}
	safeToXlsx(priodogman1024)
	// Если к-во ошибок не равно нулю
	if ErrorCount != 0 {
		t.Errorf("Blackmanharris: Test1: Расхождений между результатами MatLab и полученными результатами - " + strconv.Itoa(ErrorCount) + " из " + strconv.Itoa(N))
	}
}

// Сохранить в xlsx для дебага
func safeToXlsx(x []float64) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	lenFor := len(x)
	for ind := 0; ind < lenFor; ind++ {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), x[ind])
	}
	if err := file_graph.SaveAs("save_periodogram" + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
