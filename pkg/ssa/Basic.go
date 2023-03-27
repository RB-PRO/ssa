package ssa

import (
	"fmt"
	"log"
	"math"

	"github.com/RB-PRO/ssa/pkg/graph"
	"github.com/RB-PRO/ssa/pkg/oss"
	"gonum.org/v1/gonum/mat"
)

// Структура с данными спектрального анализа
type SPW struct {
	N       int // Количество отсчетов pw
	Win     int // Размер окна
	Res     int
	NPart   int // Количество долей res
	Overlap int
	S       int
	Imin    int
	Imax    int
	Ns      []float64
	NSF     int // номер финального отсчета финального сегмента <= N

	// Настройки конфигурации
	Graph bool // Создавать графики, где true - создавать
	Xlsx  bool // Сохранять графики в Xlsx, где true - создавать

	// Set general parameters
	Cad int // 30 кадров/сек
	Dt  float64
	Tim []float64 // Set general parameters
	Ns2 []int
	L   []int

	K int
	M int // параметр вложения в траекторное пространство
	// SSA - анализ сегментов pw
	Seg int // номер сегмента pw для визуализации
	NET int // кол-во сингулярных троек для сегментов pw

	Spw *mat.Dense

	SET12 *mat.Dense
	SET34 *mat.Dense
}

func New() *SPW {
	return &SPW{}
}
func (s *SPW) Init(pw []float64) {
	s.N = len(pw)
	s.Win = 1024
	s.Res = s.N - s.Win*int(math.Floor(float64(s.N)/float64(s.Win)))
	s.NPart = 20 // Количество долей res
	s.Res = int(math.Floor(float64(s.Res) / float64(s.NPart)))
	//overlap := (float64(win) - float64(res)) / float64(win)
	s.S = 1
	s.Imin = 1
	s.Imax = s.Win

	for s.Imax <= s.N {
		s.Ns = append(s.Ns, float64(s.S)) // номер текущего сегмента pw
		s.Imin += s.Res
		s.Imax += s.Res
		s.S++
	}
	s.S-- // кол-во перекрывающихся сегментов pw в пределах N
}

func (s *SPW) Spw_Form(pw []float64) {
	s.Spw = mat.NewDense(s.Win, s.S, nil)
	for j := 0; j < s.S; j++ {
		for i := 0; i < s.Win; i++ {
			k := (j) * s.Res
			s.Spw.Set(i, j, pw[k+i])
		}
	}
}

func (s *SPW) SET_Form() {
	sET12_sum2 := mat.NewDense(s.Win, 2, nil) // НЕ ФАКТ, ЧТО К-во строк win
	sET34_sum2 := mat.NewDense(s.Win, 2, nil) // НЕ ФАКТ, ЧТО К-во строк win
	s.SET12 = mat.NewDense(s.Win, s.S, nil)   // НЕ ФАКТ, ЧТО К-во строк win
	s.SET34 = mat.NewDense(s.Win, s.S, nil)   // НЕ ФАКТ, ЧТО К-во строк win

	for j := 0; j < s.S; j++ {
		C, LBD, RC := SSA(s.Win, s.M, s.Spw.ColView(j), s.NET)
		//fmt.Println(j, S)
		RC_T := mat.DenseCopyOf(RC.T())

		sET12_sum2.SetCol(0, RC_T.RawRowView(0))
		sET12_sum2.SetCol(1, RC_T.RawRowView(1))
		s.SET12.SetCol(j, oss.Sum2(*sET12_sum2))
		sET12_sum2.Zero()

		sET34_sum2.SetCol(0, RC_T.RawRowView(2))
		sET34_sum2.SetCol(1, RC_T.RawRowView(3))
		s.SET34.SetCol(j, oss.Sum2(*sET34_sum2))
		sET34_sum2.Zero()

		if j == s.Seg {
			// Если есть настрока формирования графика
			if s.Graph {
				graph.Imagesc(C, "C")
				graph.MakeGraphOfArray(LBD, "LBD")

				// Создаём график 1 и 2 коэффициента
				err_makeGraphYX_sET12 := graph.MakeGraphYX_VecDense(
					*mat.NewVecDense(s.Win, s.Tim[0:s.Win]),
					*(mat.VecDenseCopyOf(s.Spw.ColView(j))),
					*(mat.NewVecDense(len(oss.Vec_in_ArrFloat(s.SET12.ColView(j))), oss.Vec_in_ArrFloat(s.SET12.ColView(j)))),
					"origin", "sET12")
				if err_makeGraphYX_sET12 != nil {
					fmt.Println(err_makeGraphYX_sET12)
				}

				// Создаём график 3 и 4 коэффициента
				err_makeGraphYX_sET34 := graph.MakeGraphYX_VecDense(
					*mat.NewVecDense(s.Win, s.Tim[0:s.Win]),
					*(mat.VecDenseCopyOf(s.Spw.ColView(j))),
					*(mat.NewVecDense(len(oss.Vec_in_ArrFloat(s.SET34.ColView(j))), oss.Vec_in_ArrFloat(s.SET34.ColView(j)))),
					"origin", "sET34")
				if err_makeGraphYX_sET34 != nil {
					fmt.Println(err_makeGraphYX_sET34)
				}
			}
			// Если есть настрока сохранения данных в Xlsx
			if s.Xlsx {
				oss.Matlab_mat_Dense(C, 1, "C")
				oss.Matlab_arr_float(LBD, 2, "LBD")

				oss.Matlab_arr_float(s.Tim, 3, "tim")
				oss.Matlab_mat_Dense(*s.Spw, 3, "spw")
				oss.Matlab_mat_Dense(*s.SET12, 3, "sET12")
				log.Println("Original time series and reconstruction sET12")

				oss.Matlab_arr_float(s.Tim, 4, "tim")
				oss.Matlab_mat_Dense(*s.Spw, 4, "spw")
				oss.Matlab_mat_Dense(*s.SET34, 4, "sET34")
				log.Println("Original time series and reconstruction sET34")
			}
		}
	}
}
