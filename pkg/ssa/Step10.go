package ssa

import (
	"github.com/RB-PRO/ssa/pkg/oss"
	"gonum.org/v1/gonum/mat"
)

// # 10
// Агрегирование сегментов очищенной пульсовой волны cpw
func (s *SPW) AggregationPW() *SPW {
	NSF := s.Win + s.Res*(s.S-1) // номер финального отсчета финального сегмента <= N
	NumS, cpw_avr, cpw_med, cpw_iqr := wav(NSF, s.S, s.Win, s.Res, s.SET12)

	oss.SafeToXlsx(NumS, "NumS")
	oss.Matlab_variable(NSF, 10, "NSF")
	oss.Matlab_arr_float(s.Tim, 10, "tim")
	oss.Matlab_arr_float(cpw_avr, 10, "cpw_avr")
	oss.Matlab_arr_float(cpw_med, 10, "cpw_med")
	oss.Matlab_arr_float(cpw_iqr, 10, "cpw_iqr")

	return s
}

func wav(N, S, W, res int, sET *mat.Dense) ([]float64, []float64, []float64, []float64) {
	NS := make([]float64, N)
	w_avr := make([]float64, N)
	w_med := make([]float64, N)
	w_iqr := make([]float64, N)

	ET := mat.NewDense(N, S, nil)

	for j := 0; j < S; j++ { // цикл по сегментам
		for i := 0; i < W; i++ {
			k := (j) * res
			ET.Set(i+k, j, sET.At(i, j)) // сдвинутый сегмент ET(:,j)
		}
	}

	Smp := make([]float64, N*S)
	for i := 0; i < N; i++ {
		var nSi int
		for j := 0; j < S; j++ {
			if ET.At(i, j) != 0.0 {
				nSi++
				Smp[nSi] = ET.At(i, j)
			}
		}
		NS[i] = float64(nSi)                      // кол-во сегментов для текущего i
		w_avr[i] = oss.Mean(Smp[:nSi])            // выборочная средняя
		w_med[i] = oss.Median_floatArr(Smp[:nSi]) // медиана
		w_iqr[i] = (oss.Prctile(Smp[:nSi], 75) - oss.Prctile(Smp[:nSi], 25)) / 2.0
	}

	return NS, w_avr, w_med, w_iqr
}
