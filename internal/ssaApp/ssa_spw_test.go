package ssaApp

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/RB-PRO/ssa/pkg/oss"
	"github.com/xuri/excelize/v2"
	"gonum.org/v1/gonum/mat"
)

func TestMakePhaAcfNrm(t *testing.T) {
	vect := mat.NewVecDense(102, []float64{1.0000, 0.9054, 0.6393, 0.2519, -0.1853, -0.5847, -0.8752, -1.0000, -0.9355, -0.6940, -0.3206, 0.1179, 0.5310, 0.8437, 1.0000, 0.9675, 0.7523, 0.3935, -0.0439, -0.4790, -0.8125, -1.0000, -0.9978, -0.8073, -0.4651, -0.0353, 0.4084, 0.7542, 0.9684, 1.0000, 0.8441, 0.5302, 0.1178, -0.3280, -0.6917, -0.9316, -1.0000, -0.8813, -0.5982, -0.2035, 0.2426, 0.6238, 0.8925, 1.0000, 0.9211, 0.6706, 0.2948, -0.1512, -0.5519, -0.8509, -1.0000, -0.9638, -0.7488, -0.3935, 0.0520, 0.4703, 0.8032, 0.9945, 1.0000, 0.8197, 0.4894, 0.0551, -0.3688, -0.7215, -0.9485, -1.0000, -0.8673, -0.5755, -0.1633, 0.2595, 0.6341, 0.9000, 1.0000, 0.9158, 0.6630, 0.2720, -0.1504, -0.5473, -0.8519, -1.0000, -0.9636, -0.7493, -0.3822, 0.0385, 0.4548, 0.7949, 0.9889, 1.0000, 0.8252, 0.4847, 0.0725, -0.3541, -0.7173, -0.9463, -1.0000, -0.8691, -0.5668, -0.1754, 0.2519, 0.6359, 0.9018, 1.0000})

	PhaAcfNrm := MakePhaAcfNrm(vect)

	file_graph := excelize.NewFile()
	file_graph.NewSheet("main")
	file_graph.DeleteSheet("Sheet1")
	sig := oss.Vec_in_ArrFloat(PhaAcfNrm)
	for ind, val := range sig {
		file_graph.SetCellValue("main", "A"+strconv.Itoa(ind+1), val)
	}
	if err := file_graph.SaveAs("PhaAcfNrm_test" + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
