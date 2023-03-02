package ssa

import (
	"github.com/RB-PRO/ssa/pkg/oss"
	"gonum.org/v1/gonum/mat"
)

// Оценка АКФ сингулярных троек для сегментов pw
func ACF_estimation_of_singular_triples(lagS, win, S int, sET12 mat.Dense) mat.Dense {
	//var Acf_sET12 mat.Dense
	Acf_sET12 := mat.NewDense(lagS, S, nil)
	for j := 0; j < S; j++ {
		Acf_sET12.SetCol(j, AcfMed(lagS, win, *mat.VecDenseCopyOf(sET12.ColView(j))))
	}
	return *Acf_sET12
}
func AcfMed(lagS, win int, sET12_vec mat.VecDense) []float64 {
	// lagS - параметр погружения временного ряда (ВР) TS в траекторное пространство
	// win  - количество отсчетов ВР TS
	// TS   - ВР, содержащий win отсчетов

	acf := make([]float64, lagS)

	Y := BuildTrajectoryMatrix222(sET12_vec, lagS, win)

	cor := oss.ATa(Y) // lagS*lagS матрица корреляц-х произведений
	lon := lagS

	CorPro := oss.Diag_of_Dense(cor, 0) // ВР корреляц-го произведения для лага 0
	acf[0] = oss.Median(CorPro)         // медиана главной диагонали CorPro
	for m := 1; m < lagS; m++ {
		lon--
		diag_cor_minus_1 := oss.Diag_of_Dense(cor, m)
		if m < lagS {
			acf[m] = oss.Median(diag_cor_minus_1) / acf[0]
		}
	}
	acf[0] = 1.0
	return acf
}
