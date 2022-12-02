package main

/*
Пакет для работы с матрицами.
ПОдстоянно должен дорабатываться, в соответствии с gonum
*/

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
)

// Горизонтальная сумма массива
func sum2(a mat.Dense) []float64 {
	var output []float64
	r, _ := a.Dims()
	for i := 0; i < r; i++ {
		output = append(output, mat.Sum(a.RowView(i)))
	}
	return output
}

// МАксимальный элемент массива
func max(arr []float64) float64 {
	max_num := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max_num {
			max_num = arr[i]
		}
	}
	return max_num
}

func aTa(matr mat.Dense) mat.Dense { // Multipy matrix AT*A
	a := mat.Matrix(&matr)
	aT := a.T()
	ad := mat.DenseCopyOf(a)
	aTd := mat.DenseCopyOf(aT)
	n1, _ := aTd.Dims()
	_, m2 := ad.Dims()
	output := mat.NewDense(n1, m2, nil)
	output.Mul(aTd, ad)
	return *output
}

// модуль от всех значений вектора
func absVector(vect mat.VecDense) mat.VecDense {
	for i := 0; i < vect.Len(); i++ {
		vect.SetVec(i, math.Abs(vect.AtVec(i)))
	}
	return vect
}

/*

// Среднее значение каждого столбца
func mean(m mat.Dense) []float64 {
	_, c_m := m.Dims()
	outputArray := make([]float64, c_m)
	for ind := range outputArray {
		vect := colDense(m, ind)
		outputArray[ind] = averge(vect)
	}
	return outputArray
}
*/

// Среднее массива float64
func averge(array []float64) float64 {
	var sum float64
	for _, val := range array {
		sum += val
	}
	return sum / float64(len(array))
}

// Вернуть subdiagonal
func subdiagonal(m mat.Dense, k int) ([]float64, error) {
	var outputArray []float64
	r_m, c_m := m.Dims()
	if k == 0 {
		for i := 0; i < r_m && i < c_m; i++ {
			outputArray = append(outputArray, m.At(i, i))
		}
	} else if k < 0 {
		if k > r_m {
			return nil, errors.New("k > matrix ROW")
		}
		for i := 0; i-k < r_m && i < c_m; i++ {
			outputArray = append(outputArray, m.At(i-k, i))
		}
	} else if k > 0 {
		if k > c_m {
			return nil, errors.New("k > matrix COL")
		}
		for i := 0; i+k < c_m && i < r_m; i++ {
			outputArray = append(outputArray, m.At(i, i+k))
		}
	}
	return outputArray, nil
}

// Вернуть колонку из матрицы
func colDense(m mat.Dense, ind int) []float64 {
	row, _ := m.Dims()
	outputArray := make([]float64, row)
	for i := 0; i < row; i++ {
		outputArray[i] = m.At(i, ind)
	}
	return outputArray
}

// Вернуть строку из матрицы
func rowDense(m mat.Dense, ind int) []float64 {
	_, col := m.Dims()
	outputArray := make([]float64, col)
	for j := 0; j < col; j++ {
		outputArray[j] = m.At(ind, j)
	}
	return outputArray
}

func diag(mat mat.Dense, R int) []float64 {
	ret := make([]float64, R)
	for ind := range ret {
		ret[ind] = mat.At(ind, ind)
	}
	return ret
}

func make_diag_danse(arr []float64) mat.Dense {
	lensOfArray := len(arr)
	dens := mat.NewDense(lensOfArray, lensOfArray, nil)
	for i := 0; i < len(arr); i++ {
		dens.Set(i, i, arr[i])
	}
	return *dens
}

// AsSymDense attempts return a SymDense from the provided Dense.
func AsSymDense(m *mat.Dense) (*mat.SymDense, error) {
	r, c := m.Dims()
	if r != c {
		return nil, errors.New("matrix must be square")
	}
	mT := m.T()
	vals := make([]float64, r*c)
	idx := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mT.At(i, j) != m.At(i, j) {
				return nil, errors.New("matrix is not symmetric")
			}
			vals[idx] = m.At(i, j)
			idx++
		}
	}
	return mat.NewSymDense(r, vals), nil
}
func realyPrint(matr *mat.Dense, name string) {
	fmatr := mat.Formatted(matr, mat.Prefix(string(strings.Repeat(" ", 2+len(name)))), mat.Squeeze())
	fmt.Printf(name+" =%.3v\n", fmatr)
}

func realyPrint2(matr mat.Dense, name string) {
	fmatr := mat.Formatted(&matr, mat.Prefix(string(strings.Repeat(" ", 2+len(name)))), mat.Squeeze())
	fmt.Printf(name+" =%.3v\n", fmatr)
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

// Диагональ матрицы в зависимости от корреляции k // reference MatLab diag(A,n)
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

// Поэлементно разделить нулевое значение столбца Matrix на Vector на вектор
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
