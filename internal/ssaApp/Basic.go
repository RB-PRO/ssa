package ssaApp

import (
	"github.com/RB-PRO/ssa/pkg/graph"
	"github.com/RB-PRO/ssa/pkg/oss"
)

func SsaAnalysis() {
	fmp, _ := oss.Make_singnal_xn("fmp") // Загрузить сигнал из файла fmp.xlsx
	pw, _ := oss.Make_singnal_xn("pw")   // Загрузить сигнал из файла pw.xlsx

	graph.MakeGraphOfArray(fmp, "fmp")
	graph.MakeGraphOfArray(pw, "pw")
	oss.SafeToXlsx(fmp, "fmp") // Сохранить данные сигнала fmp в xlsx
	oss.SafeToXlsx(pw, "pw")   // Сохранить данные сигнала pw в xlsx

	// gui.New()
	// g.VievImage()

	graph.TreeXDXD([]float64{}, []float64{}, []float64{})

	//SSA_spw(pw, fmp)

	//sgl, _ := make_singnal_xn("sgl") // Загрузить сигнал из файла pw.xlsx
	//safeToXlsx(sgl, "sgl")
	//sgl2 := savitzky_goley(sgl, 33, 2)
	//safeToXlsx(sgl2, "sgl2")
}
