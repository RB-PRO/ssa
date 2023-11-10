package pmtm_test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/RB-PRO/ssa/pkg/oss"
	"github.com/RB-PRO/ssa/pkg/pmtm"
	"github.com/xuri/excelize/v2"
)

func TestPmtm(t *testing.T) {
	x, _ := oss.Make_singnal_xn("pmtm")
	y := pmtm.Pmtm(x, 1024)
	saveTXT2("test222.txt", y)
	safeToXlsx(x, y)
}

// Сохранить в xlsx для дебага
func safeToXlsx(x, y []float64) {

	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.NewSheet("golang")
	file_graph.DeleteSheet("Sheet1")
	lenFor := len(x)
	for ind := 0; ind < lenFor; ind++ {
		file_graph.SetCellValue("golang", "A"+strconv.Itoa(ind+1), x[ind])
	}
	lenFor = len(y)
	for ind := 0; ind < lenFor; ind++ {
		file_graph.SetCellValue("golang", "B"+strconv.Itoa(ind+1), y[ind])
	}
	if err := file_graph.SaveAs("save_pmtm" + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

func saveTXT2(FileName string, data []float64) {
	filePW, ErrOpenFile := os.Create(FileName)
	if ErrOpenFile != nil {
		panic(ErrOpenFile)
	}
	defer filePW.Close()
	for i := range data {
		if _, err := filePW.WriteString(fmt.Sprintf("%.8f\n", data[i])); err != nil {
			log.Println(err)
		}
	}
}
