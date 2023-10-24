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
func Matlab_arr_float(arr []float64, number int, fileName string, Path string) error {
	err := Matlab_mkDir(number, Path)
	if err != nil {
		fmt.Println(err)
	}
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind, val := range arr {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), val)
	}
	if err := file_graph.SaveAs("File_For_MatLab" + OpSystemFilder + strconv.Itoa(number) + OpSystemFilder + fileName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
	return nil
}

// Сохранить данные массива float64
func Matlab_mat_Vector(vect mat.Vector, number int, fileName string, Path string) error {
	err := Matlab_mkDir(number, Path)
	if err != nil {
		fmt.Println(err)
	}
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind := 0; ind < vect.Len(); ind++ {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), vect.AtVec(ind))
	}
	if err := file_graph.SaveAs("File_For_MatLab" + OpSystemFilder + strconv.Itoa(number) + OpSystemFilder + fileName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
	return nil
}

// Сохранить данные переменной int
func Matlab_variable(data int, number int, fileName string, Path string) error {
	err := Matlab_mkDir(number, Path)
	if err != nil {
		fmt.Println(err)
	}
	// create file
	f, err := os.Create(Path + "File_For_MatLab" + OpSystemFilder + strconv.Itoa(number) + OpSystemFilder + fileName + ".txt")
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
func Matlab_mat_Dense(X *mat.Dense, number int, fileName string, Path string) error {
	err := Matlab_mkDir(number, Path)
	if err != nil {
		fmt.Println(err)
	}

	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	n, m := X.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			file_graph.SetCellValue("main", GetColumnName(j+1)+strconv.Itoa(i+1), X.At(i, j))
		}
	}
	if err := file_graph.SaveAs(Path + "File_For_MatLab" + OpSystemFilder + strconv.Itoa(number) + OpSystemFilder + fileName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
	return nil
}

// В случае несуществования создать папку umber в matlab папке
func Matlab_mkDir(number int, Path string) error {
	if !exists(Path+"File_For_MatLab"+OpSystemFilder+strconv.Itoa(number), Path) { // Если файл не создан
		err := os.Mkdir(Path+"File_For_MatLab"+OpSystemFilder+strconv.Itoa(number), 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

// Проверка на существование папки
func exists(PathChild, Path string) bool {
	_, err := os.Stat(Path + PathChild)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
