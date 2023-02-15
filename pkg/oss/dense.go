package oss

import "gonum.org/v1/gonum/mat"

// минимальное значение элемента матрицы
func MinDense(matr mat.Dense) float64 {
	var min float64 = matr.At(0, 0)
	r, c := matr.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if min > matr.At(i, j) {
				min = matr.At(i, j)
			}
		}
	}
	return min
}

// максимальное значение элемента матрицы
func MaxDense(matr mat.Dense) float64 {
	var max float64 = matr.At(0, 0)
	r, c := matr.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if max < matr.At(i, j) {
				max = matr.At(i, j)
			}
		}
	}
	return max
}

// Makes a dense matrix of size r*c and fills it with a user-defined function.
func MakeMatrix(r int, c int, value func(i, j int) float64) *mat.Dense {
	data := make([]float64, r*c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			data[c*i+j] = value(i, j)
		}
	}
	return mat.NewDense(r, c, data)
}
