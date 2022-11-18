package main

import (
	"fmt"
	"strings"

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

func aTa(matr *mat.Dense) *mat.Dense { // Multipy matrix AT*A
	a := mat.Matrix(matr)
	aT := a.T()
	ad := mat.DenseCopyOf(a)
	aTd := mat.DenseCopyOf(aT)
	n1, _ := aTd.Dims()
	_, m2 := ad.Dims()
	output := mat.NewDense(n1, m2, nil)
	output.Mul(aTd, ad)
	return output
}

func realyPrint(matr *mat.Dense, name string) {
	fmatr := mat.Formatted(matr, mat.Prefix(string(strings.Repeat(" ", 2+len(name)))), mat.Squeeze())
	fmt.Printf(name+" =%.3v\n", fmatr)
}

func realyPrint2(matr mat.Dense, name string) {
	fmatr := mat.Formatted(&matr, mat.Prefix(string(strings.Repeat(" ", 2+len(name)))), mat.Squeeze())
	fmt.Printf(name+" =%.3v\n", fmatr)
}
