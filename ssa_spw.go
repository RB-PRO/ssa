package main

import (
	"fmt"
	"math"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

func ssa_spw(pw, fmp []float64) {
	// Сегменты отсчётов pw
	N := len(pw) // Количество отсчетов pw
	win := 1024
	res := N - win*int(math.Floor(float64(N)/float64(win)))
	nPart := 20 // Количество долей res
	res = int(math.Floor(float64(res) / float64(nPart)))
	overlap := (float64(win) - float64(res)) / float64(win)
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
	S--                    // кол-во перекрывающихся сегментов pw в пределах N
	NSF := win + res*(S-1) // номер финального отсчета финального сегмента <= N

	spw := mat.NewDense(win, S, nil)
	fmt.Println("Размеры spw:", win, S)
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
	sET12 := mat.NewDense(win, S, nil)      // НЕ ФАКТ, ЧТО К-во строк win

	fmt.Println(win)
	fmt.Println(M)
	fmt.Println(nET)
	fmt.Println("***************")

	for j := 100; j <= 100; j++ {
		//for j := 0; j < S; j++ { // цикл по сегментам  S
		//C, LBD, RC := SSA(win, M, spw[:][j], nET)
		C, LBD, RC := SSA(win, M, spw.ColView(j), nET)
		fmt.Println(j, S)
		RC_T := mat.DenseCopyOf(RC.T())

		sET12_sum2.SetCol(0, RC_T.RawRowView(0))
		sET12_sum2.SetCol(1, RC_T.RawRowView(1))
		sET12.SetCol(j, sum2(*sET12_sum2))
		sET12_sum2.Zero()

		if j == seg {
			imagesc(C, "C")

			makeGraphOfArray(LBD, "LBD-"+strconv.Itoa(j))

			fmt.Println("GRAPH")
			err_makeGraphYX := makeGraphYX_VecDense(
				*mat.NewVecDense(win, tim[0:win]),
				*(mat.VecDenseCopyOf(spw.ColView(j))),
				*(mat.NewVecDense(len(vec_in_ArrFloat(sET12.ColView(j))), vec_in_ArrFloat(sET12.ColView(j)))),
				"sET12")
			if err_makeGraphYX != nil {
				fmt.Println(err_makeGraphYX)
			}
		}
	}

	/*
		lag := math.Floor(float64(win) / 10.0)
		lagS := 2 * lag

		var Acf_sET12 mat.Dense
		for j := 0; j < S; j++ {
			Acf_sET12.SetCol(j)
			//Acf_sET12(:,j) = AcfMed(lagS,win,sET12(:,j))//; % ������������� ��� j-�� ��������
		}
	*/

	safeToXlsxMatrix(sET12, "sET12")

	// *****************
	// Оценка АКФ сингулярных троек для сегментов pw
	lag := int(math.Floor(float64(win) / 10.0)) // % наибольший лаг АКФ <= win/10
	lagS := 2 * lag
	fmt.Println("--- sET12 ---")
	fmt.Println(sET12.At(0, 0), sET12.At(0, 1), "\n", sET12.At(1, 0), sET12.At(1, 1))
	Acf_sET12 := ACF_estimation_of_singular_triples(lagS, win, S, *sET12)
	safeToXlsxM(Acf_sET12, "Acf_sET12")
	// *****************

	safeToXlsxMatrix(sET12, "sET12")

	fmt.Println("dt", cad)
	fmt.Println("dt", dt)
	fmt.Println("Imax", Imax)
	fmt.Println("Imin", Imin)
	fmt.Println("K", K)
	fmt.Println("M", M)
	fmt.Println("N", N)
	fmt.Println("nET", nET)
	fmt.Println("nPart", nPart)
	fmt.Println("ns", ns)
	fmt.Println("NSF", NSF)
	fmt.Println("overlap", overlap)
	fmt.Println("res", res)
	fmt.Println("S", S)
	safeToXlsx(tim, "tim")
	safeToXlsx(L, "L")
	fmt.Println("win", win)
	fmt.Println("seg", seg)

}

// Оценка АКФ сингулярных троек для сегментов pw
func ACF_estimation_of_singular_triples(lagS, win, S int, sET12 mat.Dense) mat.Dense {
	//var Acf_sET12 mat.Dense
	Acf_sET12 := mat.NewDense(lagS, S, nil)
	for j := 0; j < S; j++ {
		Acf_sET12.SetCol(j, AcfMed(lagS, win, sET12.ColView(j)))
	}
	return *Acf_sET12
}
func AcfMed(lagS, win int, sET12_vec mat.Vector) []float64 {
	// lagS - параметр погружения временного ряда (ВР) TS в траекторное пространство
	// win  - количество отсчетов ВР TS
	// TS   - ВР, содержащий win отсчетов

	/*
		fmt.Println("----- AcfMed ------")
		TS := mat.VecDenseCopyOf(sET12_vec)
		fmt.Println(lagS, win, TS.Len())

		Y := mat.NewDense(win-lagS+1, lagS, nil)
		fmt.Println("Y.Dims()")
		fmt.Println(Y.Dims())

		for m := 0; m < lagS; m++ {
			matV := TS.SliceVec(m, win-lagS+m+1)
			floa := vec_in_ArrFloat(matV)
			fmt.Println(m, "floa: ", len(floa))
			Y.SetCol(m, floa) // vector in []float
		}
	*/
	//Y := BuildTrajectoryMatrix(sET12_vec, win-lagS+1, win)
	//fmt.Println(sET12_vec)
	Y := BuildTrajectoryMatrix222(sET12_vec, lagS, win)
	safeToXlsxM(Y, "YYYY")
	fmt.Println("Y.Dims()")
	fmt.Println(Y.Dims())
	fmt.Println(Y.At(0, 0), Y.At(0, 1))
	fmt.Println(Y.At(1, 0), Y.At(1, 1))

	//Y(:,m) = TS(m:win-lagS+m)//; % m-й столбец траекторной матрица ВР TS
	return []float64{1.0, 2.0, 3.0}
}

func vec_in_ArrFloat(a mat.Vector) []float64 {
	b := make([]float64, a.Len())
	for i := 0; i < a.Len(); i++ {
		b[i] = a.AtVec(i)
	}
	return b
}

// Сортировка с возвратом номеров изначальных элементов
func InsertionSort(array []float64) ([]float64, []int) {
	indexArray := make([]int, len(array))
	for ind := range indexArray {
		indexArray[ind] = (ind) + 1
	}
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 {
			if array[j-1] < array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				indexArray[j-1], indexArray[j] = indexArray[j], indexArray[j-1]
			}
			j = j - 1
		}
	}
	return array, indexArray
}
