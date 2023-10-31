package ssa

import (
	"log"

	"github.com/RB-PRO/ssa/pkg/oss"
	"gonum.org/v1/gonum/mat"
)

func SSA(N int, M int, X mat.Vector, nET int) (mat.Dense, []float64, mat.Dense) {
	//  Calculate covariance matrix (trajectory approach)
	// it ensures a positive semi-definite covariance matrix
	Y := BuildTrajectoryMatrix(X, M, N) // Создать матрицу траекторий

	// fmt.Println("------", N-M+1, M)

	var Cemb mat.Dense
	Cemb.Mul(Y, Y.T())
	Cemb.Scale(1.0/float64(N-M+1), &Cemb)

	C := Cemb

	// Choose covariance estimation
	RHO, LBD := eig(C)
	LBD_diag := oss.Diag(LBD, LBD.DiagView().Diag())
	LBD_sort, _ := InsertionSort(LBD_diag)

	// Перевернуть матрицу по вертикали
	_, col_RHO := RHO.Dims()
	for j := 0; j < col_RHO/2; j++ {
		a := oss.ColDense(RHO, j)
		b := oss.ColDense(RHO, col_RHO-1-j)
		RHO.SetCol(j, b)
		RHO.SetCol(col_RHO-1-j, a)
	}

	// Calculate principal components PC
	// The principal components are given as the scalar product
	// between Y, the time-delayed embedding of X, and the eigenvectors RHO
	var PC mat.Dense
	PC.Mul(Y.T(), &RHO)

	// Calculate reconstructed components RC
	// In order to determine the reconstructed components RC,
	// we have to invert the projecting PC = Y*RHO;i.e. RC = Y*RHO*RHO'=PC*RHO'
	// Averaging along anti-diagonals gives the RCs for the original input X

	RC := mat.NewDense(N, nET, nil)
	r_PC, _ := PC.Dims()
	r_RHO, _ := RHO.Dims()
	for m := 0; m < nET; m++ {
		// invert projection
		var buf mat.Dense
		b1 := mat.NewDense(r_PC, 1, oss.ColDense(PC, m))
		b2 := mat.NewDense(r_RHO, 1, oss.ColDense(RHO, m))
		buf.Mul(b1, b2.T())

		// Перевернуть матрицу по горизонтали
		row_buf, _ := buf.Dims()
		for j := 0; j < row_buf/2; j++ {
			a := oss.RowDense(buf, j)
			b := oss.RowDense(buf, row_buf-1-j)
			buf.SetRow(j, b)
			buf.SetRow(row_buf-1-j, a)
		}

		// Anti-diagonal averaging
		for n := 0; n < N; n++ {
			diag_buf, error_subdiagonal := oss.Subdiagonal(buf, -(N-M+1)+n+1)
			if error_subdiagonal != nil {
				panic(error_subdiagonal)
			}
			RC.Set(n, m, oss.Averge(diag_buf))
		}

	}
	return C, LBD_sort, *RC
}

// Матрица траекторий
func BuildTrajectoryMatrix(s mat.Vector, L, N int) *mat.Dense {
	//fmt.Println(N, M)
	matr := mat.NewDense(L, N-L+1, nil)
	n, m := matr.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//fmt.Println(i, "*", L, "+", j, "=", i*L+j, "//", s[i*L+j])
			matr.Set(i, j, s.AtVec(i+j))
		}
	}
	return matr
}

// Матрица траекторий
func BuildTrajectoryMatrix222(s mat.VecDense, lagS, win int) mat.Dense {
	matr := mat.NewDense(win-lagS+1, lagS, nil)
	_, c := matr.Dims()
	//fmt.Println("Len matr", r, c)
	for m := 0; m < c; m++ {
		//fmt.Println(m, win-lagS+m+1)
		matr.SetCol(m, oss.Vec_in_ArrFloat(s.SliceVec(m, win-lagS+m+1)))
	}
	return *matr
}

// Returns diagonal matrix D of eigenvalues and matrix V whose columns are the corresponding right eigenvectors, so that A*V = V*D
func eig(matr mat.Dense) (mat.Dense, mat.Dense) {
	a, err := oss.AsSymDense(&matr)
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
	return ev, oss.Make_diag_danse(eigsym.Values(nil))
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
