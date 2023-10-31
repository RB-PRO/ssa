package ssa

import (
	"fmt"
	"math"

	"github.com/RB-PRO/ssa/pkg/oss"
)

// # 5
// Оценка АКФ сингулярных троек для сегментов pw
// Визуализация АКФ сингулярных троек для сегментов pw
func (s *SPW) AKF_Form() *SPW {
	s.lag = int(math.Floor(float64(s.Win) / 10.0)) // % наибольший лаг АКФ <= win/10
	lagS := 2 * s.lag
	s.Acf_sET12 = ACF_estimation_of_singular_triples(lagS, s.Win, s.S, *s.SET12)

	lgl := make([]float64, s.lag)
	for m := 0; m < len(lgl); m++ {
		lgl[m] = float64(m + 1)
	}
	s.lgl = lgl

	time := make([]float64, s.lag)
	for m := 1; m < len(time); m++ {
		time[m] = time[m-1] + s.Dt
	}
	s.time = time // Сохраняем результат

	Folder5 := fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 5)
	oss.СreateFolderIfNotExists(Folder5)
	if s.Xlsx {
		oss.Matlab_arr_float(s.Ns, Folder5, "ns"+".xlsx")
		oss.Matlab_arr_float(time, Folder5, "time"+".xlsx")
		oss.SafeToXlsxMatrix(s.Acf_sET12, Folder5, "Acf_sET12"+".xlsx")
		oss.Matlab_mat_Dense(s.Acf_sET12, Folder5, "Acf_sET12_2"+".xlsx")
	}
	// if s.Graph {
	// 	oss.Matlab_arr_float(s.Ns, Folder5, "ns"+".xlsx")
	// 	oss.Matlab_arr_float(time, Folder5, "time"+".xlsx")
	// 	oss.Matlab_mat_Dense(s.Acf_sET12, Folder5, "Acf_sET12"+".xlsx")
	// }
	return s
}
