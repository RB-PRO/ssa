package ssa

import (
	"fmt"
	"math"

	"github.com/RB-PRO/ssa/pkg/blackmanharris"
	"github.com/RB-PRO/ssa/pkg/oss"
	"github.com/RB-PRO/ssa/pkg/periodogram"
	"gonum.org/v1/gonum/mat"
)

// # 9
// Визуализация СПМ сингулярных троек сегменов pw
func (s *SPW) VisibleSPM() *SPW {
	// Оценки СПМ сингулярных троек для сегменов pw
	var iGmin, iGmax int
	smopto := 3                                                         // параметр сглаживания периодограммы Томсона
	fmi := 40.0 / 60.0                                                  // частота среза для 40 уд/мин (0.6667 Гц)
	fma := 240.0 / 60.0                                                 // частота среза для 240 уд/мин (4.0 Гц)
	Nf := 1 + s.Win/2                                                   // кол-во отсчетов частоты
	df := float64(s.Cad) / float64(s.Win-1)                             // интервал дискретизации частоты, Гц
	Fmin := fmi - float64(10*df)                                        // частота в Гц, min
	Fmax := fma + float64(10*df)                                        // частота в Гц, max
	pto_sET12 := pto_sET12_init(s.SET12, s.Spw, smopto, s.Win, Nf, s.S) // Расчёт оценки СПМ сингулярных троек для сегменов pw

	f := make([]float64, Nf)
	for i := 2; i < Nf; i++ {
		f[i] = f[i-1] + df // частота в герцах
		if math.Abs(f[i]-Fmin) <= df {
			iGmin = i
		}
		if math.Abs(f[i]-Fmax) <= df {
			iGmax = i
		}
	}

	fG := make([]float64, iGmax)
	for i := 0; i < iGmax; i++ {
		fG[i] = f[i] // сетка частот 3D-графика
	}

	Folder9 := fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 9)
	oss.СreateFolderIfNotExists(Folder9)
	if s.Xlsx {
		oss.SafeToXlsxMatrix(pto_sET12, Folder9, "pto_sET12"+".xlsx") // Сохранить в Xlsx матрицу оценки СПМ
	}
	if s.Graph {
		oss.Matlab_arr_float(s.Ns, Folder9, "ns"+".xlsx")
		oss.Matlab_arr_float(fG, Folder9, "fG"+".xlsx")
		oss.Matlab_mat_Dense(pto_sET12, Folder9, "pto_sET12"+".xlsx")
		oss.Matlab_variable(iGmin, Folder9, "iGmin"+".txt")
		oss.Matlab_variable(iGmax, Folder9, "iGmax"+".txt")
	}
	return s
}

// Формирование оценки СПМ сингулярных троек для сегменов pw
func pto_sET12_init(sET12 *mat.Dense, spw *mat.Dense, smopto, win, Nf, S int) *mat.Dense {
	pto_sET12 := mat.NewDense(Nf, S, nil)

	// Расчёт окна Блэкмана_Харриса шириной win
	// и с заданными коэффициентами
	BlackManHar := blackmanharris.Blackmanharris(win, blackmanharris.Koef4_74db)

	for j := 0; j < S; j++ {
		// Периодограмма Блэкмана_Харриса
		// pto_sET12(:,j) = periodogram(spw(:,j),blackmanharris(win),win); % Блэкмана-Харриса
		pto_sET12.SetCol(j, periodogram.Periodogram(oss.Vec_in_ArrFloat(spw.ColView(j)), BlackManHar, win))

		// Периодограмма Томсона
		// pto_sET12.SetCol(j, pmtmMy(sET12.ColView(j), smopto, win))
	}
	return pto_sET12
}
