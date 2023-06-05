package ssa

import (
	"strconv"

	"github.com/RB-PRO/ssa/pkg/graph"
	"github.com/RB-PRO/ssa/pkg/oss"
	"github.com/RB-PRO/ssa/pkg/pchip"
	"gonum.org/v1/gonum/mat"
)

// # 6, 7
// Огибающая по критерию локальных максимумов abs(acf_sET12)
// Огибающие АКФ сингулярных троек sET12 сегментов pw
// Нормированные АКФ сингулярных троек sET12 сегментов pw
func (s *SPW) Envelope() *SPW {
	//power := 0.75 // параметр спрямляющего преобразования

	EnvAcf_sET12 := *mat.NewDense(s.lag, s.S, nil)
	AcfNrm_sET12 := *mat.NewDense(s.lag, s.S, nil)
	for j := 0; j < s.S; j++ { // цикл по сегментам АКФ
		Acf_sET12_col := *mat.VecDenseCopyOf(s.Acf_sET12.ColView(j))
		absTS := oss.AbsVector(Acf_sET12_col)
		at1 := absTS.AtVec(0)
		at2 := absTS.AtVec(1)

		maxTS := *mat.NewVecDense(s.lag, nil)
		maxTS.SetVec(0, at1)

		maxN := *mat.NewVecDense(s.lag, nil)
		maxN.SetVec(0, 1)

		var Nmax int

		for m := 2; m < s.lag; m++ {
			at3 := absTS.AtVec(m)
			if (at1 <= at2) && (at2 >= at3) {
				Nmax++                        // номер очередного узла интерполяции (счетчик максимумов)
				maxN.SetVec(Nmax, float64(m)) // номер очередного максимума для ряда absTS
				maxTS.SetVec(Nmax, at2)       // отсчет очередного узла интерполяции
			}
			at1 = at2
			at2 = at3
		}
		Nmax++                                 // количество узлов интерполяции
		maxN.SetVec(Nmax, float64(s.lag))      // номер отсчета absTS финального узла интерполяции
		maxTS.SetVec(Nmax, absTS.AtVec(s.lag)) // отсчет absTS финального узла интерполяции
		NumMax := maxN.SliceVec(0, Nmax+1)

		// Интерполяция огибающей АКФ
		acfEnvelope, _, _ := pchip.Pchip(oss.Vec_in_ArrFloat(NumMax), oss.Vec_in_ArrFloat(maxTS.SliceVec(0, Nmax+1)), s.lgl, NumMax.Len(), len(s.lgl))

		EnvAcf_sET12.SetCol(j, acfEnvelope)

		// нормированные АКФ
		AcfNrm_sET12.SetCol(j, oss.VecDense_in_float64(oss.Vector_DivElemVec((s.Acf_sET12.Slice(0, s.lag, j, j+1)), EnvAcf_sET12.ColView(j))))

	}

	// Обход ошибки вывода с 856, заменив последнюю строку
	EnvAcf_sET12 = oss.EditLastRow(EnvAcf_sET12)
	AcfNrm_sET12 = oss.EditLastRow(AcfNrm_sET12)
	s.EnvAcf_sET12 = &EnvAcf_sET12
	s.AcfNrm_sET12 = &AcfNrm_sET12

	// *****************
	if s.Xlsx {
		oss.SafeToXlsxM(EnvAcf_sET12, "EnvAcf_sET12")
		oss.SafeToXlsxM(AcfNrm_sET12, "AcfNrm_sET12")
	}
	// 6 - Огибающие АКФ сингулярных троек sET12 сегментов pw
	if s.Graph {
		oss.Matlab_arr_float(s.Ns, 6, "ns")
		oss.Matlab_arr_float(s.time, 6, "time")
		oss.Matlab_mat_Dense(s.EnvAcf_sET12, 6, "EnvAcf_sET12")
		graph.SaveDat_2(EnvAcf_sET12, "File_For_MatLab"+oss.OpSystemFilder+strconv.Itoa(6)+oss.OpSystemFilder+"EnvAcf_sET12"+".dat")
		graph.SaveDat(s.Ns, "File_For_MatLab"+oss.OpSystemFilder+strconv.Itoa(6)+oss.OpSystemFilder+"ns"+".dat")
		graph.SaveDat(s.time, "File_For_MatLab"+oss.OpSystemFilder+strconv.Itoa(6)+oss.OpSystemFilder+"time"+".dat")
	}

	// 7 - Нормированные АКФ сингулярных троек sET12 сегментов pw
	if s.Graph {
		oss.Matlab_arr_float(s.Ns, 7, "ns")
		oss.Matlab_arr_float(s.time, 7, "time")
		oss.Matlab_mat_Dense(s.AcfNrm_sET12, 7, "AcfNrm_sET12")
		Folder7 := "File_For_MatLab" + oss.OpSystemFilder + strconv.Itoa(7) + oss.OpSystemFilder
		graph.SaveDat_2(AcfNrm_sET12, Folder7+"AcfNrm_sET12"+".dat")
		graph.SaveDat(s.Ns, Folder7+"ns"+".dat")
		graph.SaveDat(s.time, Folder7+"time"+".dat")
		graph.SplotMatrixFromFile(graph.Option3D{ // Задаём настройки 3D графика
			FileNameDat: Folder7 + "AcfNrm_sET12.dat",
			FileNameOut: Folder7 + "AcfNrm_sET12.png",
			Titile:      "Нормированные АКФ сингулярных троек sET12 сегментов pw",
			Xlabel:      "ns",
			Ylabel:      "lag,s",
			Zlabel:      "Acf_Nrm",
		}) // Делаем график
	}
	return s
}
