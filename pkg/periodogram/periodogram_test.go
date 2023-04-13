package periodogram_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/RB-PRO/ssa/pkg/periodogram"
)

func Test1Periodogram_1(t *testing.T) {
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

	// Если к-во ошибок не равно нулю
	if ErrorCount != 0 {
		t.Errorf("Blackmanharris: Test1: Расхождений между результатами MatLab и полученными результатами - " + strconv.Itoa(ErrorCount) + " из " + strconv.Itoa(N))
	}
}

func Test1Periodogram_2(t *testing.T) {
	fmt.Println("--- Periodogram: Test 2 ---")
	N := 1024                            // к-во отсчётов
	x1024, w1024, rez1024 := Data1024(N) // Загружаем данные для длины окна 1024

	priodogman1024 := periodogram.Periodogram(x1024, w1024, N) // Расчёт периодограммы

	var ErrorCount int

	// Циклом идём по всему массиву и сравниваем данные
	for index := range priodogman1024 {
		if int(math.Abs(100000*(priodogman1024[index]-rez1024[index]))) > 3 {
			ErrorCount++
			//t.Errorf("Blackmanharris: Элемент с индексом %v не соответствует рассчитанному. Должно было быть %v, а получено %v.", index, rez1024[index], priodogman1024[index])
		}
	}

	// Если к-во ошибок не равно нулю
	if ErrorCount != 0 {
		t.Errorf("Blackmanharris: Test1: Расхождений между результатами MatLab и полученными результатами - " + strconv.Itoa(ErrorCount) + " из " + strconv.Itoa(N))
	}
}
