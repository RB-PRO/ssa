package main

const OpSystemFilder string = "\\"

const rcond = 1e-15

func main() {
	fmp, _ := make_singnal_xn("fmp")
	pw, _ := make_singnal_xn("pw")
	//fmt.Println("fmp:", fmpN, "//", "pw", pwN)
	safeToXlsx(fmp, "fmp") // Сохранить данные в xlsx
	safeToXlsx(pw, "pw")   // Сохранить данные в xlsx

	ssa_spw(pw, fmp)

	/*
		matr := mat.NewDense(5, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
		realyPrint(matr, "matr")
		fmt.Println("diag(matr,0)")
		b0 := diag_of_Dense(*matr, 0)
		fmt.Println(b0)
		fmt.Println("diag(matr,1)")
		b1 := diag_of_Dense(*matr, 1)
		fmt.Println(b1)
		fmt.Println("diag(matr,-1)")
		bm1 := diag_of_Dense(*matr, -1)
		fmt.Println(bm1)
	*/

	// ***

	//C, LBD, RC := SSA(N, L, sig, 2)
	//safeToXlsxM(C, "C")
	//safeToXlsx(LBD, "LBD")
	//safeToXlsxM(RC, "RC")

	//makeGraph2(N, "png"+OpSystemFilder+"sig.png")
}
