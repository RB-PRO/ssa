package main

import (
	"fmt"
)

const OpSystemFilder string = "/"

const rcond = 1e-15

func main() {
	fmp, fmpN := make_singnal_xn("fmp")
	pw, pwN := make_singnal_xn("pw")
	fmt.Println("fmp:", fmpN, "//", "pw", pwN)
	safeToXlsx(fmp, "fmp") // Сохранить данные в xlsx
	safeToXlsx(pw, "pw")   // Сохранить данные в xlsx

	ssa_spw(pw, fmp)

	// ***

	//C, LBD, RC := SSA(N, L, sig, 2)
	//safeToXlsxM(C, "C")
	//safeToXlsx(LBD, "LBD")
	//safeToXlsxM(RC, "RC")

	//makeGraph2(N, "png"+OpSystemFilder+"sig.png")
}
