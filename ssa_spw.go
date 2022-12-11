package main

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
)

func ssa_spw(pw, fmp []float64) {
	// Сегменты отсчётов pw
	N := len(pw) // Количество отсчетов pw
	win := 1024
	res := N - win*int(math.Floor(float64(N)/float64(win)))
	nPart := 20 // Количество долей res
	res = int(math.Floor(float64(res) / float64(nPart)))
	//overlap := (float64(win) - float64(res)) / float64(win)
	S := 1
	Imin := 1
	Imax := win

	var ns []float64
	for Imax <= N {
		ns = append(ns, float64(S)) // номер текущего сегмента pw
		Imin = Imin + res
		Imax = Imax + res
		S++
	}
	S-- // кол-во перекрывающихся сегментов pw в пределах N
	//NSF := win + res*(S-1) // номер финального отсчета финального сегмента <= N

	spw := mat.NewDense(win, S, nil)
	//fmt.Println("Размеры spw:", win, S)
	for j := 0; j < S; j++ {
		for i := 0; i < win; i++ {
			k := (j) * res
			spw.Set(i, j, pw[k+i])
		}
	}

	// Set general parameters
	cad := 30              // 30 кадров/сек
	dt := 1 / float64(cad) // интервал дискретизации времени, сек
	tim := make([]float64, N)
	for index := 1; index < N; index++ {
		tim[index] = tim[index-1] + dt
	}

	ns2 := make([]int, S)
	for index := range ns2 {
		ns2[index] = (index + 1)
	}

	L := make([]float64, S)
	for index := range L { // цикл по сегментам pw
		L[index] = math.Floor(float64(cad) / fmp[index]) // кол-во отсчетов основного тона pw
	}

	K := 5
	M := int(float64(K) * max(L)) // параметр вложения в траекторное пространство
	// SSA - анализ сегментов pw
	seg := 100 // номер сегмента pw для визуализации
	nET := 4   // кол-во сингулярных троек для сегментов pw

	//var sET12 mat.Dense
	sET12_sum2 := mat.NewDense(win, 2, nil) // НЕ ФАКТ, ЧТО К-во строк win
	sET34_sum2 := mat.NewDense(win, 2, nil) // НЕ ФАКТ, ЧТО К-во строк win
	sET12 := mat.NewDense(win, S, nil)      // НЕ ФАКТ, ЧТО К-во строк win
	sET34 := mat.NewDense(win, S, nil)      // НЕ ФАКТ, ЧТО К-во строк win

	for j := 0; j < S; j++ {
		C, LBD, RC := SSA(win, M, spw.ColView(j), nET)
		//fmt.Println(j, S)
		RC_T := mat.DenseCopyOf(RC.T())

		sET12_sum2.SetCol(0, RC_T.RawRowView(0))
		sET12_sum2.SetCol(1, RC_T.RawRowView(1))
		sET12.SetCol(j, sum2(*sET12_sum2))
		sET12_sum2.Zero()

		sET34_sum2.SetCol(0, RC_T.RawRowView(2))
		sET34_sum2.SetCol(1, RC_T.RawRowView(3))
		sET34.SetCol(j, sum2(*sET34_sum2))
		sET34_sum2.Zero()

		if j == seg {
			imagesc(C, "C")
			matlab_mat_Dense(C, 1, "C")
			log.Println("Covariance matrix")
			makeGraphOfArray(LBD, "LBD")

			matlab_arr_float(LBD, 2, "LBD")
			log.Println("Eigenvalues")

			err_makeGraphYX_sET12 := makeGraphYX_VecDense(
				*mat.NewVecDense(win, tim[0:win]),
				*(mat.VecDenseCopyOf(spw.ColView(j))),
				*(mat.NewVecDense(len(vec_in_ArrFloat(sET12.ColView(j))), vec_in_ArrFloat(sET12.ColView(j)))),
				"sET12")
			matlab_arr_float(tim, 3, "tim")
			matlab_mat_Dense(*spw, 3, "spw")
			matlab_mat_Dense(*sET12, 3, "sET12")
			log.Println("Original time series and reconstruction sET12")

			err_makeGraphYX_sET34 := makeGraphYX_VecDense(
				*mat.NewVecDense(win, tim[0:win]),
				*(mat.VecDenseCopyOf(spw.ColView(j))),
				*(mat.NewVecDense(len(vec_in_ArrFloat(sET34.ColView(j))), vec_in_ArrFloat(sET34.ColView(j)))),
				"sET34")
			matlab_arr_float(tim, 4, "tim")
			matlab_mat_Dense(*spw, 4, "spw")
			matlab_mat_Dense(*sET34, 4, "sET34")
			log.Println("Original time series and reconstruction sET34")

			if err_makeGraphYX_sET12 != nil {
				fmt.Println(err_makeGraphYX_sET12)
			}
			if err_makeGraphYX_sET34 != nil {
				fmt.Println(err_makeGraphYX_sET34)
			}
		}
	}

	safeToXlsxMatrix(sET12, "sET12")
	safeToXlsxMatrix(sET34, "sET34")

	// *****************
	// Оценка АКФ сингулярных троек для сегментов pw
	lag := int(math.Floor(float64(win) / 10.0)) // % наибольший лаг АКФ <= win/10
	lagS := 2 * lag
	Acf_sET12 := ACF_estimation_of_singular_triples(lagS, win, S, *sET12)
	safeToXlsxM(Acf_sET12, "Acf_sET12")
	// *****************
	// Визуализация АКФ сингулярных троек для сегментов pw
	lgl := make([]float64, lag)
	for m := 0; m < len(lgl); m++ {
		lgl[m] = float64(m + 1)
	}
	time := make([]float64, lag)
	for m := 1; m < len(time); m++ {
		time[m] = time[m-1] + dt
	}
	matlab_arr_float(ns, 5, "ns")
	matlab_arr_float(time, 5, "time")
	matlab_mat_Dense(Acf_sET12, 5, "Acf_sET12")
	log.Println("Визуализация АКФ сингулярных троек для сегментов pw")

	// *****************
	// Огибающая по критерию локальных максимумов abs(acf_sET12)
	//power := 0.75 // параметр спрямляющего преобразования
	EnvAcf_sET12 := *mat.NewDense(lag, S, nil)
	AcfNrm_sET12 := *mat.NewDense(lag, S, nil)
	//for j := 16; j <= 16; j++ { // цикл по сегментам АКФ
	for j := 0; j < S; j++ { // цикл по сегментам АКФ
		Acf_sET12_col := *mat.VecDenseCopyOf(Acf_sET12.ColView(j))
		absTS := absVector(Acf_sET12_col)
		at1 := absTS.AtVec(0)
		at2 := absTS.AtVec(1)

		maxTS := *mat.NewVecDense(lag, nil)
		maxTS.SetVec(0, at1)

		maxN := *mat.NewVecDense(lag, nil)
		maxN.SetVec(0, 1)

		var Nmax int

		for m := 2; m < lag; m++ {
			at3 := absTS.AtVec(m)
			if (at1 <= at2) && (at2 >= at3) {
				Nmax++                        // номер очередного узла интерполяции (счетчик максимумов)
				maxN.SetVec(Nmax, float64(m)) // номер очередного максимума для ряда absTS
				maxTS.SetVec(Nmax, at2)       // отсчет очередного узла интерполяции
			}
			at1 = at2
			at2 = at3
		}
		Nmax++                               // количество узлов интерполяции
		maxN.SetVec(Nmax, float64(lag))      // номер отсчета absTS финального узла интерполяции
		maxTS.SetVec(Nmax, absTS.AtVec(lag)) // отсчет absTS финального узла интерполяции
		NumMax := maxN.SliceVec(0, Nmax+1)

		// Интерполяция огибающей АКФ
		/*
			acfEnvelope := pchip(vec_in_ArrFloat(NumMax),
				vec_in_ArrFloat(maxTS.SliceVec(0, Nmax+1)),
				(lgl),
				NumMax.Len(), len(lgl))
		*/

		acfEnvelope, _ := pchip(vec_in_ArrFloat(NumMax),
			vec_in_ArrFloat(maxTS.SliceVec(0, Nmax+1)),
			lgl,
			NumMax.Len(), len(lgl))

		EnvAcf_sET12.SetCol(j, acfEnvelope)

		// нормированные АКФ
		AcfNrm_sET12.SetCol(j, vecDense_in_float64(vector_DivElemVec((Acf_sET12.Slice(0, lag, j, j+1)), EnvAcf_sET12.ColView(j))))

		//fmt.Println(AcfNrm_sET12.At(lag-1, j)) // Тут возникает 850+ для 16-ти
	}

	// Обход ошибки вывода с 856, заменив последнюю строку
	EnvAcf_sET12 = editLastRow(EnvAcf_sET12)
	AcfNrm_sET12 = editLastRow(AcfNrm_sET12)

	// *****************
	safeToXlsxM(EnvAcf_sET12, "EnvAcf_sET12")
	safeToXlsxM(AcfNrm_sET12, "AcfNrm_sET12")

	// 6 - Огибающие АКФ сингулярных троек sET12 сегментов pw
	matlab_arr_float(ns, 6, "ns")
	matlab_arr_float(time, 6, "time")
	matlab_mat_Dense(EnvAcf_sET12, 6, "EnvAcf_sET12")
	log.Println("Огибающие АКФ сингулярных троек sET12 сегментов pw")

	// 7 - Нормированные АКФ сингулярных троек sET12 сегментов pw
	matlab_arr_float(ns, 7, "ns")
	matlab_arr_float(time, 7, "time")
	matlab_mat_Dense(AcfNrm_sET12, 7, "AcfNrm_sET12")
	log.Println("Нормированные АКФ сингулярных троек sET12 сегментов pw")

	// ********************************************************************
	// Мгновенная частота нормированной АКФ сингулярных троек sET12 для сегментов pw
	insFrc_AcfNrm := make([]float64, S)
	for j := 0; j < S; j++ {
		PhaAcfNrm := makePhaAcfNrm(AcfNrm_sET12.ColView(j))

		_, pCoef := pchip(vecDense_in_float64(PhaAcfNrm),
			lgl,
			lgl,
			PhaAcfNrm.Len(), len(lgl))

		//fmt.Println(pAcf[0], len(pCoef))

		FrcAcfNrm := make([]float64, lag)
		for m := 1; m < lag; m++ {
			//fmt.Println("pCoef[3*lag+m]", 2*lag+m, pCoef[2*lag+m])
			FrcAcfNrm[m] = math.Abs(pCoef[2*lag+m]) / (2.0 * math.Pi * dt)
		}
		FrcAcfNrm[0] = FrcAcfNrm[1]
		insFrc_AcfNrm[j] = median_floatArr(FrcAcfNrm) // средняя(медианная) мгновенная частотта j-го сегмента pw
	}

	smo_insFrc_AcfNrm := SavGolFilter(insFrc_AcfNrm, S, S/2, 0, 1.0)

	matlab_arr_float(ns, 8, "ns")
	matlab_arr_float(insFrc_AcfNrm, 8, "insFrc_AcfNrm")
	matlab_arr_float(smo_insFrc_AcfNrm, 8, "smo_insFrc_AcfNrm")
	err_insFrc_AcfNrm := makeGraphYX_float64(
		insFrc_AcfNrm,
		ns,
		"insFrc_AcfNrm")
	if err_insFrc_AcfNrm != nil {
		fmt.Println(err_insFrc_AcfNrm)
	}

}

func savitzky_goley(y []float64, f, k int) []float64 {

	return []float64{}
}

// Расчёты вектора PhaAcfNrm, модуль от Акосинуса.
func makePhaAcfNrm(vect mat.Vector) mat.VecDense {
	output := mat.VecDenseCopyOf(vect)

	for i := 0; i < output.Len(); i++ {
		output.SetVec(i, math.Abs(math.Acos(output.AtVec(i))))
	}

	return *output
}
