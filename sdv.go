package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

type SDVs struct {
	V *mat.Dense
	U *mat.Dense
	S *mat.Dense
}

func aaT(matr *mat.Dense) *mat.Dense { // Multipy matrix AT*A
	a := mat.Matrix(matr)
	aT := a.T()
	ad := mat.DenseCopyOf(a)
	aTd := mat.DenseCopyOf(aT)
	n1, _ := aTd.Dims()
	_, m2 := ad.Dims()
	output := mat.NewDense(n1, m2, nil)
	fmt.Print("X: ")
	fmt.Println(ad.Dims())
	fmt.Print("XT: ")
	fmt.Println(aTd.Dims())
	safeToXlsxMatrix(ad, "ad")
	safeToXlsxMatrix(aTd, "aTd")
	output.Mul(ad, aTd)
	return output
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
	safeToXlsxMatrix(matT, "matT")

	var SDVout SDVs
	var svdMat mat.SVD
	ok := svdMat.Factorize(matT, mat.SVDFull)
	if !ok {
		log.Fatal("failed to factorize A")
	}

	SDVout.V = new(mat.Dense)
	SDVout.S = new(mat.Dense)
	svdMat.VTo(SDVout.V)

	svdMat.UTo(SDVout.S)
	lenX_s, lenY_s := matT.Dims()
	//fmt.Println(lenY_s)
	valuesMat := make([]float64, lenX_s)
	//fmt.Println(len(valuesMat))
	svdMat.Values(valuesMat)

	SDVout.U = mat.NewDense(lenX_s, lenY_s, nil)
	for ind, val := range valuesMat {
		SDVout.U.Set(ind, ind, val)
	}

	//fmt.Println(SDVout.S.Dims())

	SDVout.U = mat.DenseCopyOf(SDVout.U.T())

	//fmt.Println(SDVout.S.Dims())

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
