package oss

// Этот модуль необходим для сохранения данных в папку File_For_MatLab. А файл plotting.m рисует графики на основании полученных данных.

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/xuri/excelize/v2"
	"gonum.org/v1/gonum/mat"
)

// Сохранить данные массива float64
//
//	xlsx
func Matlab_arr_float(arr []float64, Path, FileName string) error {
	// err := Matlab_mkDir(number, Path)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind, val := range arr {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), val)
	}
	// "File_For_MatLab" + OpSystemFilder + strconv.Itoa(number) + OpSystemFilder + fileName + ".xlsx"
	if err := file_graph.SaveAs(Path + FileName); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
	return nil
}

// Сохранить данные массива float64
//
//	xlsx
func Matlab_mat_Vector(vect mat.Vector, Path, FileName string) error {
	// err := Matlab_mkDir(number, Path)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind := 0; ind < vect.Len(); ind++ {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), vect.AtVec(ind))
	}
	// "File_For_MatLab" + OpSystemFilder + strconv.Itoa(number) + OpSystemFilder + fileName + ".xlsx"
	if err := file_graph.SaveAs(Path + FileName); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
	return nil
}

// Сохранить данные переменной int
//
//	txt
func Matlab_variable(data int, Path, FileName string) error {
	// err := Matlab_mkDir(number, Path)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// create file
	// Path + "File_For_MatLab" + OpSystemFilder + strconv.Itoa(number) + OpSystemFilder + fileName + ".txt"
	f, err := os.Create(Path + FileName)
	if err != nil {
		log.Println(err)
	}
	// remember to close the file
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%v", data))

	return err
}

func Mean(arr []float64) float64 {
	var meanVar float64
	for _, val := range arr {
		meanVar += val
	}
	return meanVar / float64(len(arr))
}

// *****************************************************************************

func Prctile(input []float64, percent float64) float64 {
	var percentile float64
	length := len(input)
	if length == 0 {
		return math.NaN()
	}

	if length == 1 {
		return input[0]
	}

	if percent <= 0 || percent > 100 {
		return math.NaN()
	}

	// Start by sorting a copy of the slice
	//c := sortedCopy(input)
	sort.Float64s(input)

	// Multiply percent by length of input
	index := (percent / 100) * float64(len(input))

	// Check if the index is a whole number
	if index == float64(int64(index)) {

		// Convert float to int
		i := int(index)

		// Find the value at the index
		percentile = input[i-1]

	} else if index > 1 {

		// Convert float to int via truncation
		i := int(index)

		// Find the average of the index and following values
		percentile = Mean([]float64{input[i-1], input[i]}) // Mean(Float64Data{input[i-1], input[i]})

	} else {
		return math.NaN()
	}

	return percentile

}

// *****************************************************************************

// Сохранить данные матрицы mat.Dense
func Matlab_mat_Dense(X *mat.Dense, Path, FilePathName string) error {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	n, m := X.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			file_graph.SetCellValue("main", GetColumnName(j+1)+strconv.Itoa(i+1), X.At(i, j))
		}
	}
	// Path + "File_For_MatLab" + OpSystemFilder + strconv.Itoa(number) + OpSystemFilder + fileName + ".xlsx"
	if err := file_graph.SaveAs(Path + FilePathName); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
	return nil
}
