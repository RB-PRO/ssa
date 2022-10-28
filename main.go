package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

const rcond = 1e-15

type SDVs struct {
	V *mat.Dense
	U *mat.Dense
	S *mat.Dense
}

func main() {
	XX := mat.NewDense(4, 5, []float64{1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0})

	//XX := mat.NewDense(4, 5, []float64{1, 0, 0, 0, 0, 0, 0, 2, 0, 3, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0})
	SDVsEL := SDV_single(XX)
	fmt.Print("XX ")
	fmt.Println(XX.Dims())
	realyPrint(XX, "XX")
	fmt.Print("SDVsEL.U ")
	fmt.Println(SDVsEL.U.Dims())
	realyPrint(SDVsEL.U, "SDVsEL.U")
	fmt.Print("SDVsEL.S ")
	fmt.Println(SDVsEL.S.Dims())
	realyPrint(SDVsEL.S, "SDVsEL.S")
	fmt.Print("SDVsEL.V ")
	fmt.Println(SDVsEL.V.Dims())
	realyPrint(SDVsEL.V, "SDVsEL.V")

	var multip *mat.Dense
	multip.Mul(SDVsEL.U, SDVsEL.S)
	multip.Mul(multip, SDVsEL.V)
	fmt.Print("multip ")
	fmt.Println(multip.Dims())
	realyPrint(multip, "multip")

	var L int = 40
	var N int = 300
	sig := make_singnal_xn(N) // Создать сигнал с N
	autoSSA(sig, 1, L, N)

	safeToXlsx(sig, "signal") // Сохранить данные в xlsx

	makeGraph2(N, "png"+OpSystemFilder+"sig.png")

}

func autoSSA(s []float64, r int, L int, N int) []float64 {

	X := BuildTrajectoryMatrix(s, L, N) // Создать матрицу траекторий
	safeToXlsxMatrix(X, "X")
	R := makeRank(X)

	SUV_arr := SDV(X, R)

	safeToXlsxMatrix(SUV_arr[0].U, "U")
	safeToXlsxMatrix(SUV_arr[0].V, "V")
	safeToXlsxMatrix(SUV_arr[0].S, "S")

	x := make([]float64, R)
	for i := 0; i < R; i++ {
		sumMatrix := makeSumMatrix(SUV_arr[i])
		x[i] = DiagAveraging(sumMatrix, i, R)
	}

	// тестовая штука сравнить результат с матлабом
	sumsX := 0.0
	xx := x
	for _, val := range xx {
		sumsX += val
	}
	for ind, val := range xx {
		xx[ind] = (sumsX / val) * 100.0
	}
	safeToXlsx(x, "xzx")
	safeToXlsx(xx, "xx")

	Ik := HCA(x, R)
	fmt.Println("HCA:", Ik)
	safeToXlsx(Ik, "Ik")

	yk := make([]float64, r)
	//fmt.Println(len(Ik), r)
	for k := 0; k < r; k++ {
		yk[k] = sumOfIk(Ik[k])
	}

	makeGraphOfArray(x, "png"+OpSystemFilder+"x.png")

	safeToXlsx(x, "DiagAveraging") // Сохранить данные в xlsx
	makeGraph2(R, "png"+OpSystemFilder+"DiagAveraging.png")
	return yk
}

func HCA(x []float64, R int) []float64 {

	return []float64{0.0, 1.0, 2.0}
}
func sumOfIk(a float64) float64 {

	return 0.0
}

func makeSumMatrix(SUV SDVs) *mat.Dense {
	var output mat.Dense
	a := mat.DenseCopyOf(SUV.U)
	b := mat.DenseCopyOf(SUV.V)
	c := mat.DenseCopyOf(SUV.S)

	fmt.Println("\n************************************************************************************")
	fmt.Println(a.Dims())
	fmt.Println(b.Dims())
	fmt.Println(c.Dims())
	output.Add(a, b)
	output.Add(a, c)
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
func SDV(X *mat.Dense, rank int) []SDVs {
	SDVsout := make([]SDVs, rank)
	for i := 0; i < rank; i++ {
		/*
			// Это с дроблением на матрицы
			X_y, _ := X.Dims()
			kk := X.Slice(0, X_y, i, i+X_y) // Взять часть матрицы X
			kek := mat.DenseCopyOf(kk)      // Преобразовать в Dense
			SDVsout[i] = SDV_single(kek)    // Сохранить значение
		*/

		// это без дробления на матрицы
		SDVsout[i] = SDV_single(X)
	}
	return SDVsout
}
func SDV_single(matT *mat.Dense) SDVs {
	//matT := mat.NewDense(5, 4, []float64{1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0})
	//matT := mat.NewDense(5, 4, []float64{1, 0, 0, 0, 0, 0, 0, 2, 0, 3, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0})
	safeToXlsxMatrix(matT, "matT")

	var SDVout SDVs
	var svdMat mat.SVD
	ok := svdMat.Factorize(matT, mat.SVDFull)
	if !ok {
		log.Fatal("failed to factorize A")
	}

	SDVout.V = new(mat.Dense)
	SDVout.U = new(mat.Dense)
	svdMat.VTo(SDVout.V)
	SDVout.V = mat.DenseCopyOf(SDVout.V.T())
	svdMat.UTo(SDVout.U)
	lenX_s, lenY_s := matT.Dims()
	//fmt.Println(lenY_s)
	valuesMat := make([]float64, lenX_s)
	//fmt.Println(len(valuesMat))
	svdMat.Values(valuesMat)

	SDVout.S = mat.NewDense(lenX_s, lenY_s, nil)
	for ind, val := range valuesMat {
		SDVout.S.Set(ind, ind, val)
	}

	return SDVout
}
func makeRank(matr *mat.Dense) int {
	var svd mat.SVD
	ok := svd.Factorize(matr, mat.SVDFull)
	if !ok {
		log.Fatal("failed to factorize A")
	}
	rank := svd.Rank(rcond)
	if rank == 0 {
		log.Fatal("zero rank system")
	}
	return (rank)
}

func BuildTrajectoryMatrix(s []float64, L int, N int) *mat.Dense {
	K := N - L + 1
	matr := mat.NewDense(L, K, nil)
	n, m := matr.Dims()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//fmt.Println(i, "*", L, "+", j, "=", i*L+j, "//", s[i*L+j])
			matr.Set(i, j, s[i+j])
		}
	}
	return matr
}
func make_diag_danse(arr []float64) *mat.Dense {
	lensOfArray := len(arr)
	dens := mat.NewDense(lensOfArray, lensOfArray, nil)
	for i := 0; i < len(arr); i++ {
		dens.Set(i, i, arr[i])
	}
	return dens
}
