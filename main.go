package main

const OpSystemFilder string = "/" // "\\" for Windows, "/" for Linux

const rcond = 1e-15

func main() {
	fmp, _ := make_singnal_xn("fmp") // Загрузить сигнал из файла fmp.xlsx
	pw, _ := make_singnal_xn("pw")   // Загрузить сигнал из файла pw.xlsx

	safeToXlsx(fmp, "fmp") // Сохранить данные сигнала fmp в xlsx
	safeToXlsx(pw, "pw")   // Сохранить данные сигнала pw в xlsx

	ssa_spw(pw, fmp)

	//sgl, _ := make_singnal_xn("sgl") // Загрузить сигнал из файла pw.xlsx
	//safeToXlsx(sgl, "sgl")
	//sgl2 := savitzky_goley(sgl, 33, 2)
	//safeToXlsx(sgl2, "sgl2")
}
