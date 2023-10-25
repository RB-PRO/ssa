package ssaApp

import (
	"github.com/RB-PRO/ssa/pkg/oss"
)

func SsaAnalysis() {
	// Path := "files/"
	fmp, _ := oss.Make_singnal_xn("fmp") // Загрузить сигнал из файла fmp.xlsx
	pw, _ := oss.Make_singnal_xn("pw")   // Загрузить сигнал из файла pw.xlsx

	// graph.MakeGraphOfArray(fmp, Path+"files/"+"fmp.xlsx")
	// graph.MakeGraphOfArray(pw, Path+"files/"+"pw.xlsx")
	// oss.SafeToXlsx(fmp, Path+"files/"+"fmp") // Сохранить данные сигнала fmp в xlsx
	// oss.SafeToXlsx(pw, Path+"files/"+"pw")   // Сохранить данные сигнала pw в xlsx

	// gui.New()
	// g.VievImage()

	// graph.TreeXDXD([]float64{}, []float64{}, []float64{})

	//SSA_spw(pw, fmp)
	SSS_spw2(pw, fmp)

	//sgl, _ := make_singnal_xn("sgl") // Загрузить сигнал из файла pw.xlsx
	//safeToXlsx(sgl, "sgl")
	//sgl2 := savitzky_goley(sgl, 33, 2)
	//safeToXlsx(sgl2, "sgl2")
}
