package oss

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
	"gonum.org/v1/gonum/mat"
)

// ***
func SafeToXlsx(sig []float64, Path, FileName string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind, val := range sig {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), val)
	}
	// "files" + OpSystemFilder + name + ".xlsx"
	if err := file_graph.SaveAs(Path + FileName); err != nil {
		fmt.Println(err)
	}
}
func SafeToXlsxMatrix(X *mat.Dense, Path, FileName string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	n, m := X.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			file_graph.SetCellValue("main", GetColumnName(j+1)+strconv.Itoa(i+1), X.At(i, j))
		}
	}
	// Path + "files" + OpSystemFilder + xlsxName + ".xlsx"
	if err := file_graph.SaveAs(Path + FileName); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
}
func SafeToXlsxDualArray(X [][]float64, Path, FileName string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind1, val1 := range X {
		for ind2, val2 := range val1 {
			file_graph.SetCellValue("main", GetColumnName(ind2+1)+strconv.Itoa(ind1), val2)
		}
	}
	// Path + "files" + OpSystemFilder + xlsxName + ".xlsx"
	if err := file_graph.SaveAs(Path + FileName); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
}
func SafeToXlsxM(X mat.Dense, FilePathName string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	n, m := X.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			file_graph.SetCellValue("main", GetColumnName(j+1)+strconv.Itoa(i+1), X.At(i, j))
		}
	}
	// Path + "files" + OpSystemFilder + xlsxName + ".xlsx"
	if err := file_graph.SaveAs(FilePathName); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
}
func GetColumnName(col int) string { /*
		name := make([]byte, 0, 3) // max 16,384 columns (2022)
		const aLen = 'Z' - 'A' + 1 // alphabet length
		for ; col > 0; col /= aLen + 1 {
			name = append(name, byte('A'+(col-1)%aLen))
		}
		for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
			name[i], name[j] = name[j], name[i]
		}
		return string(name)
	*/
	asd, _ := excelize.ColumnNumberToName(col)
	return asd
}

// Создать вложенные подпапки, если их не существует
func СreateFolderIfNotExists(folderPath string) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
		// fmt.Printf("Папка %s создана\n", folderPath)
	} else if err != nil {
		return err
	}
	// else { fmt.Printf("Папка %s уже существует\n", folderPath)	}
	return nil
}
