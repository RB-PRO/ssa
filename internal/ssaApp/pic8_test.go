package ssaApp_test

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/RB-PRO/ssa/internal/ssaApp"
	"github.com/xuri/excelize/v2"
	"gonum.org/v1/gonum/mat"
)

// Тестирование функции расчёта мгновенной частота нормированной АКФ сингулярных троек sET12 для сегментов pw
func TestInstantaneous_frequency_of_normalized_ACF_sET12(t *testing.T) {
	//varArr, _ := xlsx2floatArr("pic8_test.xlsx", "Лист1")
	//SaveVarStr("pic8_test.txt", varArr)

	// Поулчить значение переменной из файла
	pic8_test_txt, pic8_test_txt_Error := OpenVarStr("pic8_test.txt")
	if pic8_test_txt_Error != nil {
		t.Error(pic8_test_txt_Error)
	}
	// Сделать из строки массив float64
	float64Dense, float64Dense_Error := Str2FloatArr(pic8_test_txt)
	if float64Dense_Error != nil {
		t.Error(float64Dense_Error)
	}

	AcfNrm_sET12 := mat.NewDense(102, 200, float64Dense)

	S := 200
	win := 1024
	lag := int(math.Floor(float64(win) / 10.0)) // % наибольший лаг АКФ <= win/10
	cad := 30                                   // 30 кадров/сек
	dt := 1 / float64(cad)                      // интервал дискретизации времени, сек
	lgl := make([]float64, lag)
	for m := 0; m < len(lgl); m++ {
		lgl[m] = float64(m + 1)
	}

	insFrc_AcfNrm, insFrc_AcfNrm_Error := ssaApp.Instantaneous_frequency_of_normalized_ACF_sET12(*AcfNrm_sET12, S, lag, dt, lgl)
	if insFrc_AcfNrm_Error != nil {
		t.Error(insFrc_AcfNrm_Error)
	}

	SaveVarStr("insFrc_AcfNrm.txt", fmt.Sprintf("%v", insFrc_AcfNrm))

}

// читает файл и выдаёт строку для использования в тестах
func xlsx2floatArr(filename string, sSheet string) (output string, errorFile error) {
	xlsxFile, errorFile := excelize.OpenFile(filename, excelize.Options{})
	if errorFile != nil {
		return "", errorFile
	}
	defer xlsxFile.Close()

	// Получить все строки в sSheet
	rows, errorFile := xlsxFile.GetRows(sSheet)
	if errorFile != nil {
		return "", errorFile
	}
	for _, row := range rows {
		for _, colCell := range row {
			output += colCell + ","
		}
	}
	return output[:len(output)-1], nil
}

// Сохранить переменную в файл
func SaveVarStr(filename, data string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

// Сохранить значение файла в переменную
func OpenVarStr(filename string) (string, error) {
	b, err := os.ReadFile(filename) // just pass the file name
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Преобразовать строку в массив float
func Str2FloatArr(data string) ([]float64, error) {
	strs := strings.Split(data, ",")
	outputArray := make([]float64, len(strs))
	for indexStrs, valueStr := range strs {

		if n, err := strconv.ParseFloat(valueStr, 64); err == nil {
			outputArray[indexStrs] = n
		} else {
			fmt.Println(">>>", valueStr)

			//	return nil, err
		}
	}
	return outputArray, nil
}
