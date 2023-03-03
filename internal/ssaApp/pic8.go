package ssaApp

import (
	"math"

	"github.com/RB-PRO/ssa/pkg/oss"
	"github.com/RB-PRO/ssa/pkg/pchip"
	"gonum.org/v1/gonum/mat"
)

// #Расчёт мгновенной частота нормированной АКФ сингулярных троек sET12 для сегментов pw
func Instantaneous_frequency_of_normalized_ACF_sET12(AcfNrm_sET12 mat.Dense, S, lag int, dt float64, lgl []float64) ([]float64, error) {
	insFrc_AcfNrm := make([]float64, S)
	for j := 0; j < S; j++ {
		PhaAcfNrm := MakePhaAcfNrm(AcfNrm_sET12.ColView(j))

		_, _, coef := pchip.Pchip(oss.VecDense_in_float64(*PhaAcfNrm),
			lgl,
			lgl,
			PhaAcfNrm.Len(), len(lgl))

		//spline := pchip.NewCubic(lgl, oss.VecDense_in_float64(PhaAcfNrm))

		//oss.SafeToXlsx(pCoef, "pCoef")
		//oss.SafeToXlsxDualArray(pCoef, "pCoef")
		//fmt.Println(pAcf[0], len(pCoef))

		FrcAcfNrm := make([]float64, lag)
		for m := 1; m < lag; m++ {

			FrcAcfNrm[m] = math.Abs(coef.B[m-1]) / (2.0 * math.Pi) / dt // pchip
			//FrcAcfNrm[m] = math.Abs(spline.Weights[m-1]) / (2.0 * math.Pi * dt) // spline
		}
		FrcAcfNrm[0] = FrcAcfNrm[1]
		insFrc_AcfNrm[j] = oss.Median_floatArr(FrcAcfNrm) // средняя(медианная) мгновенная частотта j-го сегмента pw
	}
	return insFrc_AcfNrm, nil
}
