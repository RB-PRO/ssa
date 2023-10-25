package ssa

import (
	"fmt"
	"log"
	"math"

	"github.com/RB-PRO/ssa/pkg/graph"
	"github.com/RB-PRO/ssa/pkg/oss"
	"github.com/RB-PRO/ssa/pkg/pchip"
	"github.com/pconstantinou/savitzkygolay"
	"gonum.org/v1/gonum/mat"
)

// # 8
// Мгновенная частота нормированной АКФ сингулярных троек sET12 для сегментов pw
func (s *SPW) MomentFrequency() *SPW {
	insFrc_AcfNrm, insFrc_AcfNrmErr := Instantaneous_frequency_of_normalized_ACF_sET12(s.AcfNrm_sET12, s.S, s.lag, s.Dt, s.lgl)
	if insFrc_AcfNrmErr != nil {
		log.Println(insFrc_AcfNrmErr)
	}

	// filter savitzky-goley
	filter, savitzky_goley_Error := savitzkygolay.NewFilterWindow(53)
	if savitzky_goley_Error != nil {
		log.Fatalln(savitzky_goley_Error)
	}

	// Савицкий-Голей
	smo_insFrc_AcfNrm, _ := filter.Process(insFrc_AcfNrm, s.lgl)

	if s.Graph {
		Folder8 := fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 8)
		oss.СreateFolderIfNotExists(Folder8)

		oss.Matlab_arr_float(s.Ns, Folder8, "ns"+".xlsx")
		oss.Matlab_arr_float(insFrc_AcfNrm, Folder8, "insFrc_AcfNrm"+".xlsx")
		oss.Matlab_arr_float(smo_insFrc_AcfNrm, Folder8, "smo_insFrc_AcfNrm"+".xlsx")

		FolderPNG := fmt.Sprintf("%s/png/", s.Dir.zeropath)
		oss.СreateFolderIfNotExists(FolderPNG)
		err_insFrc_AcfNrm := graph.MakeGraphYX_float64(
			insFrc_AcfNrm,
			s.Ns,
			FolderPNG, "insFrc_AcfNrm"+".png")
		if err_insFrc_AcfNrm != nil {
			fmt.Println(err_insFrc_AcfNrm)
		}
		err_insFrc_AcfNrm = graph.MakeGraphYX_float64(
			smo_insFrc_AcfNrm,
			s.Ns,
			FolderPNG, "smo_insFrc_AcfNrm"+".png")
		if err_insFrc_AcfNrm != nil {
			fmt.Println(err_insFrc_AcfNrm)
		}
		err_insFrc_AcfNrm = graph.MakeGraphYX_VecDense(
			*mat.NewVecDense(len(s.Ns), s.Ns),
			*mat.NewVecDense(len(insFrc_AcfNrm), insFrc_AcfNrm),
			*mat.NewVecDense(len(smo_insFrc_AcfNrm), smo_insFrc_AcfNrm),
			"origin", FolderPNG, "insFrc_AcfNrm")
		if err_insFrc_AcfNrm != nil {
			fmt.Println(err_insFrc_AcfNrm)
		}
	}
	return s
}

// Расчёт мгновенной частота нормированной АКФ сингулярных троек sET12 для сегментов pw
func Instantaneous_frequency_of_normalized_ACF_sET12(AcfNrm_sET12 *mat.Dense, S, lag int, dt float64, lgl []float64) ([]float64, error) {
	insFrc_AcfNrm := make([]float64, S)
	for j := 0; j < S; j++ {
		PhaAcfNrm := MakePhaAcfNrm(AcfNrm_sET12.ColView(j))

		//_, _, coef := pchip.Pchip(oss.VecDense_in_float64(*PhaAcfNrm), lgl,			lgl,			PhaAcfNrm.Len(), len(lgl))

		//fmt.Println(len(lgl), PhaAcfNrm.Len())

		slopes := pchip.Pchip2(lgl, oss.VecDense_in_float64(*PhaAcfNrm))

		//spline := pchip.NewCubic(lgl, oss.VecDense_in_float64(PhaAcfNrm))

		//oss.SafeToXlsx(pCoef, "pCoef")
		//oss.SafeToXlsxDualArray(pCoef, "pCoef")
		//fmt.Println(pAcf[0], len(pCoef))

		FrcAcfNrm := make([]float64, lag)
		for m := 1; m < lag; m++ {

			//FrcAcfNrm[m] = math.Abs(coef.B[m-1]) / (2.0 * math.Pi) / dt // pchip
			FrcAcfNrm[m] = math.Abs(slopes[m-1]) / (2.0 * math.Pi) / dt // pchip

			//FrcAcfNrm[m] = math.Abs(spline.Weights[m-1]) / (2.0 * math.Pi * dt) // spline
		}
		FrcAcfNrm[0] = FrcAcfNrm[1]
		insFrc_AcfNrm[j] = oss.Median_floatArr(FrcAcfNrm) // средняя(медианная) мгновенная частотта j-го сегмента pw
	}
	return insFrc_AcfNrm, nil
}

// Расчёты вектора PhaAcfNrm, модуль от Акосинуса.
func MakePhaAcfNrm(vect mat.Vector) *mat.VecDense {
	output := mat.VecDenseCopyOf(vect)
	for i := 0; i < output.Len(); i++ {
		output.SetVec(i, math.Abs(math.Acos(output.AtVec(i))))
	}
	return output
}
