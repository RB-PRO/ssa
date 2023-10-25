package graph

import (
	"fmt"
	"os"

	"gonum.org/v1/gonum/mat"
)

// Сохранить двумерный массив в формате .dat
func SaveDat_2(Data mat.Dense, Path, FileName string) error {

	// Создаём файл
	f, ErrorCreateFile := os.Create(Path + FileName)
	if ErrorCreateFile != nil {
		return ErrorCreateFile
	}
	defer f.Close()

	// Сформировать данные для записи в файл
	var OutputString string
	RowsLen, _ := Data.Dims()
	for i := 0; i < RowsLen; i++ {
		row := Data.RawRowView(i)
		for j := range row {
			OutputString += fmt.Sprint(row[j]) + "\t"
		}
		OutputString += "\n"
	}

	// Записать данные в файл
	_, ErrorWriteFile := f.Write([]byte(OutputString))
	if ErrorWriteFile != nil {
		return ErrorWriteFile
	}

	return nil
}

// Сохранить массив в формат .dat
func SaveDat(Data []float64, Path, FileName string) error {
	// Создаём файл
	f, ErrorCreateFile := os.Create(Path + FileName)
	if ErrorCreateFile != nil {
		return ErrorCreateFile
	}
	defer f.Close()

	// Сформировать данные для записи в файл
	var OutputString string
	for _, val := range Data {
		OutputString += fmt.Sprint(val) + "\n"
	}

	// Записать данные в файл
	_, ErrorWriteFile := f.Write([]byte(OutputString))
	if ErrorWriteFile != nil {
		return ErrorWriteFile
	}

	return nil
}
