package main

import (
	"fmt"
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
	for j := 0; j < S; j++ {                // цикл по сегментам  S
		//C, LBD, RC := SSA(win, M, spw[:][j], nET)
		_, _, RC := SSA(win, M, spw.ColView(j), nET)
		fmt.Println(j, S)
		RC_T := mat.DenseCopyOf(RC.T())

		sET12_sum2.SetCol(0, RC_T.RawRowView(0))
		sET12_sum2.SetCol(1, RC_T.RawRowView(1))
		sET12.SetCol(j, sum2(*sET12_sum2))
		sET12_sum2.Zero()
	}
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
