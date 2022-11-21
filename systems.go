package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
	"gonum.org/v1/gonum/mat"
)

// ***
func safeToXlsx(sig []float64, name string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind, val := range sig {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), val)
	}
	if err := file_graph.SaveAs("files" + OpSystemFilder + name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
func safeToXlsxMatrix(X *mat.Dense, xlsxName string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	n, m := X.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			file_graph.SetCellValue("main", getColumnName(j+1)+strconv.Itoa(i+1), X.At(i, j))
		}
	}
	if err := file_graph.SaveAs("files" + OpSystemFilder + xlsxName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
}
func safeToXlsxDualArray(X [][]float64, xlsxName string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	for ind1, val1 := range X {
		for ind2, val2 := range val1 {
			file_graph.SetCellValue("main", getColumnName(ind2+1)+strconv.Itoa(ind1), val2)
		}
	}
	if err := file_graph.SaveAs("files" + OpSystemFilder + xlsxName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
}
func safeToXlsxM(X mat.Dense, xlsxName string) {
	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	n, m := X.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			file_graph.SetCellValue("main", getColumnName(j+1)+strconv.Itoa(i+1), X.At(i, j))
		}
	}
	if err := file_graph.SaveAs("files" + OpSystemFilder + xlsxName + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	file_graph.Close()
}
func getColumnName(col int) string { /*
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
