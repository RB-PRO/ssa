package main

import (
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

const OpSystemFilder string = "/" // "\\" for Windows, "/" for Linux

const rcond = 1e-15

func main() {
	fmp, _ := make_singnal_xn("fmp") // Загрузить сигнал из файла fmp.xlsx
	pw, _ := make_singnal_xn("pw")   // Загрузить сигнал из файла pw.xlsx

	safeToXlsx(fmp, "fmp") // Сохранить данные сигнала fmp в xlsx
	safeToXlsx(pw, "pw")   // Сохранить данные сигнала pw в xlsx

	//ssa_spw(pw, fmp)

	sgl, _ := make_singnal_xn("sgl") // Загрузить сигнал из файла pw.xlsx

	safeToXlsx(sgl, "sgl")
	sgl2 := savitzky_goley(sgl, 33, 2)
	safeToXlsx(sgl2, "sgl2")

}

func savitzky_goley(y []float64, f, k int) []float64 {
	// Функция, реализующая сглаживание с помощью фильтра Савицкого-Голея.
	// f - окно сглаживания, желательно, чтобы оно было нечётным числом.
	// k - степень полинома, оно должно быть меньше чем f

	x := make([]int, len(y))
	for ind := range x {
		x[ind] = ind
	}
	n := len(x)
	f = int(math.Floor(float64(f)))
	f = min2(f, n)
	hf := (f - 1) / 2

	//v := dinamicMatrix_float64(f, k+1)
	var v mat.Dense = *(mat.NewDense(f, k+1, nil))

	t := make([]int, hf*2+1)
	for ind := range t {
		t[ind] = -hf + ind
	}

	for i := 0; i < f; i++ {
		for j := 0; j <= k; j++ {
			v.Set(i, j, math.Pow(float64(t[i]), float64(j)))
		}
	}

	q, r := QRDec(v)

	realyPrint2(q, "q")
	realyPrint2(r, "r")

	return y
}

func QRDec(a mat.Dense) (mat.Dense, mat.Dense) {
	row, col := a.Dims()
	q := a
	//r := dinamicMatrix_float64(col, col)
	var r mat.Dense = *mat.NewDense(col, col, nil)

	var matVecDense mat.VecDense

	fmt.Println("row", row, "col", col)

	for i := 0; i < row; i++ {

		for j := 0; j < i-1; j++ {

			//r[j][i], _ = multipleArray(q[j][:], q[i][:])
			matVecDense.MulVec(q.ColView(j), q.ColView(i))
			r.Set(j, i, matVecDense.AtVec(0))

			matVecDense.ScaleVec(r.At(j, i), q.ColView(j))

			matVecDense.SubVec(q.ColView(i), &matVecDense)
			q.SetCol(i, vecDense_in_float64(matVecDense))

			fmt.Println(i, j, "r[j][i]", r.At(j, i))

			//q[i], _ = subVectors(q[i][:], multipleConstArray(q[j][:], r[i][j]))

		}

		//fmt.Println("r", r)
		//fmt.Println("q", q)

		matVecDense = *mat.VecDenseCopyOf(q.ColView(i))

		r.Set(i, i, matVecDense.Norm(2))

		if r.At(i, i) == 0.0 {
			break
		}
		//q[:][i] = divisionConstArray(q[:][i], r[i][i])
		matVecDense.ScaleVec(1/r.At(i, i), q.ColView(i))
		q.SetCol(i, vecDense_in_float64(matVecDense))
	}
	return q, r
}

func p_norm(arr []float64, p float64) float64 {
	// The general definition for the p-norm of a vector v that has N elements is
	// If p = 1, then the resulting 1-norm is the sum of the absolute values of the vector elements.
	// If p = 2, then the resulting 2-norm gives the vector magnitude or Euclidean length of the vector.
	// If p = Inf, then v = max(arr)
	// If p = -Inf, then v = min(arr)
	var sum float64
	if p == 1 {
		for _, value := range arr {
			sum += value
		}
	} else if p == 2 {
		for _, value := range arr {
			sum += math.Pow(value, 2.0)
		}
		sum = math.Sqrt(sum)
	} else {
		for _, value := range arr {
			sum += math.Pow(value, p)
		}
		sum = math.Pow(sum, 1/p)
	}
	return sum
}

func multipleArray(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0.0, errors.New("Length vector is different")
	}
	var sum float64
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum, nil
}

// Вычитание вектора из вектора
func subVectors(a, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, errors.New("Length vector is different")
	}
	for i := 0; i < len(a); i++ {
		a[i] -= b[i]
	}
	return a, nil
}

// Умножение вектора на константу
func multipleConstArray(a []float64, multipleConstant float64) []float64 {
	for i := 0; i < len(a); i++ {
		a[i] *= multipleConstant
	}
	return a
}

// Умножение вектора на константу
func divisionConstArray(a []float64, multipleConstant float64) []float64 {
	for i := 0; i < len(a); i++ {
		a[i] /= multipleConstant
	}
	return a
}

// Создать матрицу с размерами row, col. int
func dinamicMatrix(row, col int) [][]int {
	matrix := make([][]int, row)
	for ind := range matrix {
		matrix[ind] = make([]int, col)
	}
	return matrix
}

// Создать матрицу с размерами row, col. float64
func dinamicMatrix_float64(row, col int) [][]float64 {
	matrix := make([][]float64, row)
	for ind := range matrix {
		matrix[ind] = make([]float64, col)
	}
	return matrix
}

// Минимальное число из двух
func min2(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
