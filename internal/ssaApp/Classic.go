package ssaApp

import (
	"main/pkg/oss"
)

func SsaAnalysis() {
	fmp, _ := oss.Make_singnal_xn("fmp") // Загрузить сигнал из файла fmp.xlsx
	pw, _ := oss.Make_singnal_xn("pw")   // Загрузить сигнал из файла pw.xlsx

	oss.SafeToXlsx(fmp, "fmp") // Сохранить данные сигнала fmp в xlsx
	oss.SafeToXlsx(pw, "pw")   // Сохранить данные сигнала pw в xlsx

	SSA_spw(pw, fmp)

	//sgl, _ := make_singnal_xn("sgl") // Загрузить сигнал из файла pw.xlsx
	//safeToXlsx(sgl, "sgl")
	//sgl2 := savitzky_goley(sgl, 33, 2)
	//safeToXlsx(sgl2, "sgl2")
}
