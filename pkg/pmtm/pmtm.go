package pmtm

import (
	"math/cmplx"

	"github.com/mjibson/go-dsp/fft"
)

func Pmtm(x []float64, n int) []float64 {
	/*
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
	*/
	//Fs := 1000.0 // Sampling frequency
	//T := 1 / Fs  // Sampling period
	L := 1024 // Length of signal

	// fft
	fft_signal := fft.FFTReal(x)
	//fmt.Println(len(fft_signal))
	P2 := make([]float64, len(fft_signal))
	for i := 0; i < len(P2); i++ {
		// P2[i] = cmplx.Abs((fft_signal[i])*(fft_signal[i])) / 1024.0
		P2[i] = cmplx.Abs((fft_signal[i]) / 1024.0)
	}
	// Возвести в квадрат
	P1 := make([]float64, L/2+1)
	for i := 0; i < len(P1); i++ {
		P1[i] = 2 * P2[i]
	}

	return P2
}
