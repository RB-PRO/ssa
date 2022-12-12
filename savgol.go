package main

import (
	"fmt"
	"math"

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
	coeffs := savGolCoeffs(window_length, polyorder, deriv, delta, true)
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
func savGolCoeffs(window_length int, polyorder int, deriv int, delta float64, useInConv bool) []float64 {
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
	a := makeMatrix(polyorder+1, len(aRowTemplate), func(i, j int) float64 {
		return math.Pow(aRowTemplate[j], float64(i))
	})

	// `b` determines which order derivative is returned.
	// The coefficient assigned to b[deriv] scales the result to take into
	// account the order of the derivative and the sample spacing.
	b := makeMatrix(polyorder+1, 1, func(i, j int) float64 {
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

// Makes a dense matrix of size r*c and fills it with a user-defined function.
func makeMatrix(r int, c int, value func(i, j int) float64) *mat.Dense {
	data := make([]float64, r*c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			data[c*i+j] = value(i, j)
		}
	}
	return mat.NewDense(r, c, data)
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
	b = makeMatrix(max2(m, n), nhrs, func(i, j int) float64 {
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
