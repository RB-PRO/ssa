package ssa

import (
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

	if s.Xlsx {
		oss.SafeToXlsxMatrix(s.Acf_sET12, "Acf_sET12")
	}
	if s.Graph {
		oss.Matlab_arr_float(s.Ns, 5, "ns")
		oss.Matlab_arr_float(time, 5, "time")
		oss.Matlab_mat_Dense(s.Acf_sET12, 5, "Acf_sET12")
	}
	return s
}
