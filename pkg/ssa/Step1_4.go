package ssa

import (
	"fmt"

	"github.com/RB-PRO/ssa/pkg/graph"
	"github.com/RB-PRO/ssa/pkg/oss"
	"gonum.org/v1/gonum/mat"
)

// # 1, 2, 3, 4
func (s *SPW) SET_Form() *SPW {
	sET12_sum2 := mat.NewDense(s.Win, 2, nil) // НЕ ФАКТ, ЧТО К-во строк win
	sET34_sum2 := mat.NewDense(s.Win, 2, nil) // НЕ ФАКТ, ЧТО К-во строк win
	s.SET12 = mat.NewDense(s.Win, s.S, nil)   // НЕ ФАКТ, ЧТО К-во строк win
	s.SET34 = mat.NewDense(s.Win, s.S, nil)   // НЕ ФАКТ, ЧТО К-во строк win

	for j := 0; j < s.S; j++ {
		C, LBD, RC := SSA(s.Win, s.M, s.Spw.ColView(j), s.NET)

		RC_T := mat.DenseCopyOf(RC.T())

		sET12_sum2.SetCol(0, RC_T.RawRowView(0))
		sET12_sum2.SetCol(1, RC_T.RawRowView(1))
		s.SET12.SetCol(j, oss.Sum2(*sET12_sum2))

		sET34_sum2.SetCol(0, RC_T.RawRowView(2))
		sET34_sum2.SetCol(1, RC_T.RawRowView(3))
		s.SET34.SetCol(j, oss.Sum2(*sET34_sum2))

		/////////////////////////////
		// fmt.Printf("%f,%f\n",
		// 	sET12_sum2.At(0, 0), sET12_sum2.At(0, 1))
		// fmt.Printf("%f,%f\n",
		// 	sET12_sum2.At(1, 0), sET12_sum2.At(1, 1))
		// break

		sET12_sum2.Zero()
		sET34_sum2.Zero()

		// fmt.Println(">>>", j, s.Seg, s.S)
		if j == s.S/2 {
			fmt.Println(C.Dims())
			fmt.Println(s.Win, s.M, s.NET)
			// Если есть настрока формирования графика
			if s.Graph {
				FolderPNG := fmt.Sprintf("%s/png/", s.Dir.zeropath)
				oss.СreateFolderIfNotExists(FolderPNG)
				// Создаём график 1 и 2 коэффициента
				err_makeGraphYX_sET12 := graph.MakeGraphYX_VecDense(
					*mat.NewVecDense(s.Win, s.Tim[0:s.Win]),
					*(mat.VecDenseCopyOf(s.Spw.ColView(j))),
					*(mat.NewVecDense(len(oss.Vec_in_ArrFloat(s.SET12.ColView(j))), oss.Vec_in_ArrFloat(s.SET12.ColView(j)))),
					"origin", FolderPNG, "sET12")
				if err_makeGraphYX_sET12 != nil {
					fmt.Println(err_makeGraphYX_sET12)
				}
				// Создаём график 3 и 4 коэффициента
				err_makeGraphYX_sET34 := graph.MakeGraphYX_VecDense(
					*mat.NewVecDense(s.Win, s.Tim[0:s.Win]),
					*(mat.VecDenseCopyOf(s.Spw.ColView(j))),
					*(mat.NewVecDense(len(oss.Vec_in_ArrFloat(s.SET34.ColView(j))), oss.Vec_in_ArrFloat(s.SET34.ColView(j)))),
					"origin", FolderPNG, "sET34")
				if err_makeGraphYX_sET34 != nil {
					fmt.Println(err_makeGraphYX_sET34)
				}
			}
			// Если есть настрока сохранения данных в Xlsx
			if s.Xlsx {

				FolderSave := fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 5)
				oss.СreateFolderIfNotExists(FolderSave)

				graph.Imagesc(C, FolderSave, "C"+".xlsx")
				graph.MakeGraphOfArray(LBD, FolderSave, "LBD"+".png")

				oss.СreateFolderIfNotExists(fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 1))
				oss.СreateFolderIfNotExists(fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 2))
				oss.СreateFolderIfNotExists(fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 3))
				oss.СreateFolderIfNotExists(fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 4))

				oss.Matlab_mat_Dense(&C, fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 1), "C"+".xlsx")
				oss.Matlab_arr_float(LBD, fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 2), "LBD"+".xlsx")

				oss.Matlab_arr_float(s.Tim, fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 3), "tim"+".xlsx")
				oss.Matlab_mat_Dense(s.Spw, fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 3), "spw"+".xlsx")
				oss.Matlab_mat_Dense(s.SET12, fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 3), "sET12"+".xlsx")
				// log.Println("Original time series and reconstruction sET12")

				oss.Matlab_arr_float(s.Tim, fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 4), "tim"+".xlsx")
				oss.Matlab_mat_Dense(s.Spw, fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 4), "spw"+".xlsx")
				oss.Matlab_mat_Dense(s.SET34, fmt.Sprintf("%s/MatLab/%d/", s.Dir.zeropath, 4), "sET34"+".xlsx")
				// log.Println("Original time series and reconstruction sET34")
			}
		}
	}
	return s
}
