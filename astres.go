package main

import (
	"errors"
	"log"

	"gonum.org/v1/gonum/mat"
)

/*
func autoSSA(s []float64, r int, L int, N int) []float64 {
	// % Step 1: Build trajectory matrix
	X := BuildTrajectoryMatrix(s, L, N) // Создать матрицу траекторий
	safeToXlsxMatrix(X, "X")

	// % Step 2: SVD

	var S mat.Dense
	S.Mul(mat.Matrix(X), X.T())
	safeToXlsxM(S, "S")

	R := makeRank(X)
	fmt.Println("Rank:", R)

	// ***************************************************

	EigenVectors, EigenValues := eig(S)
	safeToXlsxM(EigenVectors, "EigenVectors")
	safeToXlsxM(EigenValues, "EigenValues")

	// ***************************************************

	EigenValues.Scale(-1.0, &EigenValues)
	aa := diag(EigenValues, R)
	i_R := make([]int, R) // Пока что примитивная сортировка. без I
	for ind := range i_R {
		i_R[ind] = 40 - ind
	}

	fmt.Println("diag", aa)
	sort.Float64s(aa)
	fmt.Println("sort", aa)
	for ind := range aa {
		aa[ind] *= -1.0
	}
	fmt.Println("minusOdin", aa)
	// ***************************************************

	return []float64{0.0, 1.0, 2.0}
}
*/
// Returns diagonal matrix D of eigenvalues and matrix V whose columns are the corresponding right eigenvectors, so that A*V = V*D
func eig(matr mat.Dense) (mat.Dense, mat.Dense) {
	a, err := AsSymDense(&matr)
	if err != nil {
		panic(err)
	}
	var eigsym mat.EigenSym
	ok := eigsym.Factorize(a, true)
	if !ok {
		log.Fatal("Symmetric eigendecomposition failed")
	}
	var ev mat.Dense
	eigsym.VectorsTo(&ev)
	return ev, make_diag_danse(eigsym.Values(nil))
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
func HCA(x []float64, r int) []float64 {

	return []float64{0.0, 1.0, 2.0}
}
func sumOfIk(a float64) float64 {

	return 0.0
}

func makeSumMatrix(SUV SDVs) *mat.Dense {
	var output mat.Dense
	u := mat.DenseCopyOf(SUV.U)
	v := mat.DenseCopyOf(SUV.V)
	s := mat.DenseCopyOf(SUV.S)

	output.Mul(v, s)
	output.Mul(&output, u)
	return mat.DenseCopyOf(&output)
}

// Computation- and space-eﬃcient implementation of SSA - Anton Korobeynikov
func DiagAveraging(SUV *mat.Dense, k int, N int) float64 {
	var gk float64
	K, L := SUV.Dims()
	//fmt.Println("***\nK", K, "L", L)
	//realyPrint(SUV.U, "SUV.U")
	if 0 < k && k < L {
		for j := 0; j < k; j++ {
			//fmt.Println("j", j, "k-j", k-j)
			gk += SUV.At(j, k-j)
		}
		gk *= 1 / float64(k)
	}
	if L <= k && k <= K {
		for j := 0; j < k; j++ {
			//fmt.Println("j", j, "k-j", k-j)
			gk += SUV.At(j, k-j)
		}
		gk *= 1 / float64(L)
	}
	if K < k && k < N {
		for j := k - K + 1; j < k; j++ {
			//fmt.Println("j", j, "k-j", k-j)
			gk += SUV.At(j, k-j)
		}
		gk *= 1 / (float64(N) - float64(k) + 1)
	}
	return gk
}

func BuildTrajectoryMatrix(s mat.Vector, L int, N int) *mat.Dense {
	K := N - L + 1
	matr := mat.NewDense(L, K, nil)
	n, m := matr.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//fmt.Println(i, "*", L, "+", j, "=", i*L+j, "//", s[i*L+j])
			matr.Set(i, j, s.AtVec(i+j))
		}
	}
	return matr
}
func make_diag_danse(arr []float64) mat.Dense {
	lensOfArray := len(arr)
	dens := mat.NewDense(lensOfArray, lensOfArray, nil)
	for i := 0; i < len(arr); i++ {
		dens.Set(i, i, arr[i])
	}
	return *dens
}

func diag(mat mat.Dense, R int) []float64 {
	ret := make([]float64, R)
	for ind := range ret {
		ret[ind] = mat.At(ind, ind)
	}
	return ret
}
