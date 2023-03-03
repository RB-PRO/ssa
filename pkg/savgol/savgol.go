package savgol

import (
	"fmt"
	"math"

	"github.com/RB-PRO/ssa/pkg/oss"
	"github.com/mjibson/go-dsp/fft"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/lapack/lapack64"
	"gonum.org/v1/gonum/mat"
)

// SavGolFilter implements Savitzky-Golay filter (https://docs.scipy.org/doc/scipy/reference/generated/scipy.signal.savgol_filter.html)
// based on: https://github.com/scipy/scipy/blob/v1.3.0rc1/scipy/signal/_savitzky_golay.py#L227
func SavGolFilter(x []float64, window_length int, polyorder int, deriv int /*=0*/, delta float64 /*=1.0*/) []float64 {
	// computing filter coefficients
	// the outputs of this step seem to be numerically same as the Python code ones
	coeffs := SavGolCoeffs(window_length, polyorder, deriv, delta, true)
	// convolving the original signal with the filter coefficients
	// note: the outputs of this step are not completely numerically same as the Python code ones (because the latter uses different convolution function)

	if len(x) < window_length {
		return nil
	}

	convolutionOutput := Convolve1d(x, coeffs)

	result := make([]float64, 0)
	for i := 0; i < len(convolutionOutput); i++ {
		result = append(result, real(convolutionOutput[i]))
	}
	return result
}

func Convolve1d(x []float64, coeffs []float64) []complex128 {
	xFFt := fft.FFTReal(x)
	coefftsFFt := fft.FFTReal(coeffs)

	if len(xFFt) != len(coefftsFFt) {
		for i := len(coefftsFFt); i < len(xFFt); i++ {
			coefftsFFt = append(coefftsFFt, 0+0i)
		}
	}
	matrixProduct := make([]complex128, 0)
	for i := 0; i < len(xFFt); i++ {
		matrixProduct = append(matrixProduct, xFFt[i]*coefftsFFt[i])
	}
	return fft.IFFT(matrixProduct)
}

// Computes Savitzky-Golay filter coefficients.
func SavGolCoeffs(window_length int, polyorder int, deriv int, delta float64, useInConv bool) []float64 {
	if polyorder >= window_length {
		panic("polyorder must be less than window_length.")
	}
	fmt.Println("window_length", window_length)
	if window_length%2 == 0 {
		panic("window_length must be odd.")
	}
	pos := window_length / 2
	if pos < 0 {
		panic("pos must be nonnegative.")
	}

	// Form the design matrix `A`. The columns of `A` are powers of the integers
	// from -pos to window_length - pos - 1.  The powers (i.e. rows) range
	// from 0 to polyorder.
	aRowTemplate := Arange(-pos, window_length-pos)
	if useInConv {
		// Reverse so that result can be used in a convolution.
		floats.Reverse(aRowTemplate)
	}
	a := oss.MakeMatrix(polyorder+1, len(aRowTemplate), func(i, j int) float64 {
		return math.Pow(aRowTemplate[j], float64(i))
	})

	// `b` determines which order derivative is returned.
	// The coefficient assigned to b[deriv] scales the result to take into
	// account the order of the derivative and the sample spacing.
	b := oss.MakeMatrix(polyorder+1, 1, func(i, j int) float64 {
		if i != deriv {
			return 0
		}
		return float64(factorial(deriv)) / math.Pow(delta, float64(deriv))
	})

	// finding the least-squares solution of A*x = b
	coeff := LstSq(a, b)
	if _, cols := coeff.Dims(); cols != 1 {
		panic(errors.Errorf("SHOULD NOT HAPPEN: LstSq result contains %d columns instead of 1", cols))
	}
	return coeff.RawMatrix().Data
}

// LstSq computes least-squares solution to equation A*x = b, i.e. computes a vector x such that the 2-norm “|b - A x|“ is minimized.
func LstSq(a, b *mat.Dense) *mat.Dense {
	// m is a number of columns in `a`, n is a number of rows in `a`
	m, n := a.Dims()
	if m == 0 || n == 0 {
		panic("zero-sized problem is not supported (confuses LAPACK)")
	}

	// nhrs (why is it called so?) is a number of rows in `b`
	m2, nhrs := b.Dims()
	if m2 != m {
		panic(errors.Errorf("shape mismatch: a and b should have the same number of rows: %d != %d", m, m2))
	}

	// LAPACK uses `b` as an output parameter as well - and therefore wants it to be resized from (m, nhrs) to (max(m,n), nhrs)
	// here we copy `b` anyway (even if it doesn't need to be resized) - to avoid overwriting the user-supplied `b`
	b = oss.MakeMatrix(max2(m, n), nhrs, func(i, j int) float64 {
		if i < m {
			return b.At(i, j)
		}
		return 0
	})

	// LAPACK function for computing least-squares solutions to linear equations
	gels := func(work []float64, lwork int) bool {
		return lapack64.Gels(blas.NoTrans, a.RawMatrix(), b.RawMatrix(), work, lwork)
	}

	// retrieving the size of work space needed (this is how LAPACK interfaces are designed:
	// if we call the function with lwork=-1, it returns the work size needed in work[0])
	work := make([]float64, 1)
	gels(work, -1)
	lwork := int(math.Ceil(work[0]))

	// solving the equation itself
	result := gels(make([]float64, lwork), lwork)
	if !result {
		panic(errors.Errorf("gels: computation didn't converge: A='%+v', b='%+v'", a, b))
	}

	// dgels writes a solution into b
	return b
}

// Arange implements `np.arange` - i.e. returns a list of integers (start, ..., stop - 1) in the form of []float64
func Arange(start int, stop int) []float64 {
	return Linspace(float64(start), float64(stop-1), stop-start)
}

// Zeroes returns an array of zeroes of specified size.
// It's encouraged to use it instead of just make() in case the further code relies on the fact that the array contains zeroes.
func Zeroes(size int) []float64 {
	return make([]float64, size)
}

// Ones return an array of ones of specified size.
func Ones(size int) []float64 {
	result := make([]float64, size)
	for i := range result {
		result[i] = 1
	}
	return result
}

// Linspace implements `np.linspace` - i.e. splits the interval [start, end] into `num - 1` equal intervals and returns `num` split points.
func Linspace(start, end float64, num int) []float64 {
	if num < 0 {
		panic(errors.Errorf("number of samples, %d, must be non-negative.", num))
	}
	result := make([]float64, num)
	step := (end - start) / float64(num-1)
	for i := range result {
		result[i] = start + float64(i)*step
	}
	return result
}

// maximum of two integers
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// computes `n!`
func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// *** my realyze

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

	q, _ := QRDec(v)

	ymid := filt(q, hf, y)

	ybegin := yBegin(q, hf, f, y)
	yend := yEnd(q, hf, f, y)
	for i := 0; i < len(ybegin); i++ {
		ymid[i] = ybegin[i]
	}

	// fmt.Println(len(ymid)-len(yend)-1, len(ymid), "len(yend)", len(yend))

	for i := len(ymid) - len(yend); i < len(ymid); i++ {
		//fmt.Println(i - (len(ymid) - len(yend)))
		//fmt.Println(ybegin[i-(len(ymid)-len(yend))])
		ymid[i] = ybegin[i-(len(ymid)-len(yend))]
	}

	return ymid
}

func yBegin(q mat.Dense, hf, f int, y []float64) []float64 {
	//ybegin = q(1:hf,:)*q'*y(1:f);
	_, col := q.Dims()

	sliseQ := q.Slice(0, hf, 0, col)
	var matr mat.Dense
	matr.Mul(sliseQ, q.T())

	var matVecDense mat.VecDense
	matVecDense.MulVec(mat.Matrix(&matr), mat.Vector(mat.NewVecDense(f, y[:f])))

	return oss.VecDense_in_float64(matVecDense)
}
func yEnd(q mat.Dense, hf, f int, y []float64) []float64 {
	//yend   = q((hf+2):end,:)*q'*y(n-f+1:n);
	row, col := q.Dims()

	sliseQ := q.Slice(hf+1, row, 0, col)
	var matr mat.Dense
	matr.Mul(sliseQ, q.T())

	var matVecDense mat.VecDense
	matVecDense.MulVec(mat.Matrix(&matr), mat.Vector(mat.NewVecDense(f, y[len(y)-f:])))
	return oss.VecDense_in_float64(matVecDense)
}

func filt(q mat.Dense, hf int, x []float64) []float64 {
	// b=q*q(hf+1,:)'
	var b mat.VecDense
	elem := q.RowView(hf + 1)            // q(hf+1,:)'
	var matr mat.Matrix = mat.Matrix(&q) // q
	b.MulVec(matr, elem)                 // q*q(hf+1,:)'

	y := make([]float64, len(x))
	y[0] = b.AtVec(0) * x[0]

	var y_tmp, y_b1, y_j float64

	for i := 1; i < len(x); i++ {
		y_tmp = 0
		y_b1 = b.AtVec(1) * x[i]

		for j := 1; j < b.Len(); j++ {
			if i-j == 0 {
				y_j = b.AtVec(j) * x[i-j+1]
				y_tmp = y_tmp + y_j
				y[i] = y_b1 + y_tmp
				break
			} else {
				y_j = b.AtVec(j) * x[i-j+1]
				y_tmp = y_tmp + y_j
				if i-j >= 1 {
					y[i] = y_b1 + y_tmp
				}
			}
		}

	}

	return y
}

// QR Decomposition - Некорректный вывод
func QRDec(a mat.Dense) (mat.Dense, mat.Dense) {
	_, col := a.Dims()
	q := a
	var r mat.Dense = *mat.NewDense(col, col, nil)
	var matVecDense mat.VecDense
	for i := 0; i < col; i++ {
		for j := 0; j < i; j++ {
			r.Set(i, j, oss.MulVecToVec(q.ColView(j), q.ColView(i)))
			matVecDense.ScaleVec(r.At(j, i), q.ColView(j))
			matVecDense.SubVec(q.ColView(i), &matVecDense)
			q.SetCol(i, oss.VecDense_in_float64(matVecDense))
		}
		matVecDense = *mat.VecDenseCopyOf(q.ColView(i))
		r.Set(i, i, matVecDense.Norm(2.0))
		//r.Set(i, i, p_norm(q.ColView(i), 2.0))

		if r.At(i, i) == 0.0 {
			break
		}
		matVecDense.ScaleVec(1/r.At(i, i), q.ColView(i))
		q.SetCol(i, oss.VecDense_in_float64(matVecDense))
	}
	return q, r
}

func p_norm(arr mat.Vector, p float64) float64 {
	//func p_norm(arr []float64, p float64) float64 {
	// The general definition for the p-norm of a vector v that has N elements is
	// If p = 1, then the resulting 1-norm is the sum of the absolute values of the vector elements.
	// If p = 2, then the resulting 2-norm gives the vector magnitude or Euclidean length of the vector.
	// If p = Inf, then v = max(arr)
	// If p = -Inf, then v = min(arr)
	var sum float64
	if p == 1 {
		for ind := 0; ind < arr.Len(); ind++ {
			//for _, value := range arr {
			sum += arr.AtVec(ind)
		}
	} else if p == 2 {
		for ind := 0; ind < arr.Len(); ind++ {
			//for _, value := range arr {
			sum += math.Pow(arr.AtVec(ind), 2.0)
		}
		sum = math.Sqrt(sum)
	} else {
		for ind := 0; ind < arr.Len(); ind++ {
			//for _, value := range arr {
			sum += math.Pow(arr.AtVec(ind), p)
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
