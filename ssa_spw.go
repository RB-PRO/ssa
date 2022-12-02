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

			makeGraphOfArray(LBD, "LBD")
			matlab_arr_float(LBD, 2, "LBD")

			err_makeGraphYX_sET12 := makeGraphYX_VecDense(
				*mat.NewVecDense(win, tim[0:win]),
				*(mat.VecDenseCopyOf(spw.ColView(j))),
				*(mat.NewVecDense(len(vec_in_ArrFloat(sET12.ColView(j))), vec_in_ArrFloat(sET12.ColView(j)))),
				"sET12")
			matlab_arr_float(tim, 3, "tim")
			matlab_mat_Dense(*spw, 3, "spw")
			matlab_mat_Dense(*sET12, 3, "sET12")

			err_makeGraphYX_sET34 := makeGraphYX_VecDense(
				*mat.NewVecDense(win, tim[0:win]),
				*(mat.VecDenseCopyOf(spw.ColView(j))),
				*(mat.NewVecDense(len(vec_in_ArrFloat(sET34.ColView(j))), vec_in_ArrFloat(sET34.ColView(j)))),
				"sET34")
			matlab_arr_float(tim, 4, "tim")
			matlab_mat_Dense(*spw, 4, "spw")
			matlab_mat_Dense(*sET34, 4, "sET34")

			if err_makeGraphYX_sET12 != nil {
				fmt.Println(err_makeGraphYX_sET12)
			}
			if err_makeGraphYX_sET34 != nil {
				fmt.Println(err_makeGraphYX_sET34)
			}
		}
	}

	/*
		lag := math.Floor(float64(win) / 10.0)
		lagS := 2 * lag

		var Acf_sET12 mat.Dense
		for j := 0; j < S; j++ {
			Acf_sET12.SetCol(j)
			//Acf_sET12(:,j) = AcfMed(lagS,win,sET12(:,j))//; %
		}
	*/

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
	// *****************
	// Огибающая по критерию локальных максимумов abs(acf_sET12)
	//power := 0.75 // параметр спрямляющего преобразования
	EnvAcf_sET12 := *mat.NewDense(lag, S, nil)
	AcfNrm_sET12 := *mat.NewDense(lag, S, nil)
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
		acfEnvelope := pchip(vec_in_ArrFloat(NumMax),
			vec_in_ArrFloat(maxTS.SliceVec(0, Nmax+1)),
			(lgl),
			NumMax.Len(), len(lgl))
		EnvAcf_sET12.SetCol(j, acfEnvelope)

		// нормированные АКФ
		AcfNrm_sET12.SetCol(j, vecDense_in_float64(vector_DivElemVec((Acf_sET12.Slice(0, lag, j, j+1)), EnvAcf_sET12.ColView(j))))
	}
	// *****************
	safeToXlsxM(EnvAcf_sET12, "EnvAcf_sET12")
	safeToXlsxM(AcfNrm_sET12, "AcfNrm_sET12")

	// 6 - Огибающие АКФ сингулярных троек sET12 сегментов pw
	matlab_arr_float(ns, 6, "ns")
	matlab_arr_float(time, 6, "time")
	matlab_mat_Dense(EnvAcf_sET12, 6, "EnvAcf_sET12")

	// 7 - Нормированные АКФ сингулярных троек sET12 сегментов pw
	matlab_arr_float(ns, 7, "ns")
	matlab_arr_float(time, 7, "time")
	matlab_mat_Dense(AcfNrm_sET12, 7, "AcfNrm_sET12")

}

// Поэлементно разделить вектор на вектор
func vector_DivElemVec(a mat.Matrix, b mat.Vector) mat.VecDense {
	var div_vectors mat.VecDense
	var div_Dense mat.Dense
	div_Dense.CloneFrom(a)
	//fmt.Println(div_Dense.Dims())
	asd := div_Dense.ColView(0)
	//fmt.Println(">", asd.Len(), b.Len())
	div_vectors.DivElemVec(asd, b)
	return div_vectors
}

// func interpl(NumMax, maxTS mat.Vector, lgl []float64) []float64 {
func interpl(NumMax, maxTS mat.Vector, lgl []float64) []float64 {

	// https://russianblogs.com/article/26831691722/
	// https://github.com/liaoyinan/pchip

	/*
		 //	function y = hermite( x0,y0,y1,x )
		 // Интерполяционный полином Эрмита
		 // x0 - абсцисса точки
		 // y0 - значение функции
		 // y1 - значение производной
		 // m точек интерполяции вводятся с массивом x
		n=length(x0);m=length(x);
		for k=1:m
		    yy=0.0;
		    for i=1:n
		     h=1.0;
		     a=0.0;
		      for j=1:n
		         if j~=i
		           h=h*((x(k)-x0(j))/(x0(i)-x0(j)))^2;
		           a=1/(x0(i)-x0(j))+a;
		         end
		      end
		      yy=yy+h*((x0(i)-x(k))*(2*a*y0(i)-y1(i))+y0(i));
		end
		y(k)=yy;
	*/

	return []float64{}
}

// модуль от всех значений вектора
func absVector(vect mat.VecDense) mat.VecDense {
	for i := 0; i < vect.Len(); i++ {
		vect.SetVec(i, math.Abs(vect.AtVec(i)))
	}
	return vect
}

// Оценка АКФ сингулярных троек для сегментов pw
func ACF_estimation_of_singular_triples(lagS, win, S int, sET12 mat.Dense) mat.Dense {
	//var Acf_sET12 mat.Dense
	Acf_sET12 := mat.NewDense(lagS, S, nil)
	for j := 0; j < S; j++ {
		Acf_sET12.SetCol(j, AcfMed(lagS, win, *mat.VecDenseCopyOf(sET12.ColView(j))))
	}
	return *Acf_sET12
}
func AcfMed(lagS, win int, sET12_vec mat.VecDense) []float64 {
	// lagS - параметр погружения временного ряда (ВР) TS в траекторное пространство
	// win  - количество отсчетов ВР TS
	// TS   - ВР, содержащий win отсчетов

	acf := make([]float64, lagS)

	Y := BuildTrajectoryMatrix222(sET12_vec, lagS, win)

	cor := aTa(Y) // lagS*lagS матрица корреляц-х произведений
	lon := lagS

	CorPro := diag_of_Dense(cor, 0) // ВР корреляц-го произведения для лага 0
	acf[0] = median(CorPro)         // медиана главной диагонали CorPro
	for m := 1; m < lagS; m++ {
		lon--
		diag_cor_minus_1 := diag_of_Dense(cor, m)
		if m < lagS {
			acf[m] = median(diag_cor_minus_1) / acf[0]
		}
	}
	acf[0] = 1.0
	return acf
}

// Диагональ матрицы в зависимости от корреляции k
func diag_of_Dense(matr mat.Dense, k int) mat.VecDense {
	r, c := matr.Dims()
	var matr2 mat.Matrix
	switch {
	case k > 0:
		matr2 = matr.Slice(0, r, k, c)
		break
	case k < 0:
		matr2 = matr.Slice(-k, r, 0, c)
		break
	default:
		matr2 = matr.Slice(0, r, 0, c)
		break
	}

	vect := mat.NewVecDense(mat.DenseCopyOf(matr2).DiagView().Diag(), nil)

	for i := 0; i < vect.Len(); i++ {
		vect.SetVec(i, matr2.At(i, i))
	}
	return *vect
}

// Получить медианное значение массива
func median(dataVect mat.VecDense) float64 {
	dataVect = sortVecDense(dataVect)
	var median float64
	l := dataVect.Len()
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataVect.AtVec(l/2-1) + dataVect.AtVec(l/2)) / 2
	} else {
		median = dataVect.AtVec(l / 2)
	}
	return median
}

// Сортировка вектора массива по возрастанию.
func sortVecDense(dataVect mat.VecDense) mat.VecDense {
	dataVectLength := dataVect.Len()
	for i := 1; i < dataVectLength; i++ {
		j := i - 1
		for j >= 0 && dataVect.AtVec(j) > dataVect.AtVec(j+1) {
			vspom := dataVect.AtVec(j)
			dataVect.SetVec(j, dataVect.AtVec(j+1))
			dataVect.SetVec(j, vspom)
			j--
		}
	}
	return dataVect
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
