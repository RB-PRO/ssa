package pmtm

import (
	"fmt"
	"math"
)

func Pmtm(x []float64, n int) []float64 {
	// Размер массива
	N := len(x)
	// Расчет преобразования Фурье
	X := make([]float64, N)
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			X[k] += x[i] * math.Cos(2*math.Pi*float64(k)*float64(i)/float64(N))
		}
	}
	// Расчет периодограммы Томсона
	P := make([]float64, n)
	for k := 0; k < n; k++ {
		P[k] = math.Pow(X[k], 2) / float64(N)
		for i := 1; i < N/2; i++ {
			fmt.Println("len(P)", len(P), "len(X)", len(X), k, k+i)
			P[k] += 2 * math.Pow(X[k+i], 2) / float64(N)
		}
	}
	return P
}
