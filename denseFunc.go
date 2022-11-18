package main

/*
Пакет для работы с матрицами.
ПОдстоянно должен дорабатываться, в соответствии с gonum
*/

import (
	"errors"

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

/*
// Создать динамический двумерный массив
func dinamicArray(r, c int) [][]float64 {
	spw := make([][]float64, r)
	for index := range spw {
		spw[index] = make([]float64, c)
	}
	return spw
}

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
