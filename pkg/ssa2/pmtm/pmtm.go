package pmtm

import "math"

// Pxx = PMTM(X,NW,NFFT) указывает длину БПФ, используемую для вычисления оценок PSD.
// Для реального X Pxx имеет (NFFT/2+1) строк, если NFFT четное, и (NFFT + 1)/2 строк,
// если NFFT нечетное. Для сложного X значение Pxx всегда имеет длину NFFT.
// Если значение NFFT указано как пустое, значение NFFT устанавливается равным либо 256,
// либо следующей степени 2, превышающей длину X, в зависимости от того, что на % больше.
func Pmtm(x []float64, NW int, NFFT int) []float64 {

	// // Размер массива
	// N := len(x)
	// // Расчет преобразования Фурье
	// X := make([]float64, N)
	// for k := 0; k < N; k++ {
	// 	for i := 0; i < N; i++ {
	// 		X[k] += x[i] * math.Cos(2*math.Pi*float64(k)*float64(i)/float64(N))
	// 	}
	// }
	// // Расчет периодограммы Томсона
	// P := make([]float64, (NFFT/2 + 1))
	// for k := 0; k < (NFFT/2 + 1); k++ {
	// 	P[k] = math.Pow(X[k], 2) / float64(N)
	// 	for i := 1; i < N/2; i++ {
	// 		// fmt.Println("len(P)", len(P), "len(X)", len(X), k, k+i)
	// 		P[k] += 2 * math.Pow(X[k+i], 2) / float64(N)
	// 	}
	// }
	// return P

	//Fs := 1000.0 // Sampling frequency
	//T := 1 / Fs  // Sampling period

	// // fft
	// fft_signal := fft.FFTReal(x)
	// //fmt.Println(len(fft_signal))
	// P2 := make([]float64, len(fft_signal))
	// for i := 0; i < len(P2); i++ {
	// 	// P2[i] = cmplx.Abs((fft_signal[i])*(fft_signal[i])) / 1024.0
	// 	P2[i] = cmplx.Abs((fft_signal[i]) / 1024.0)
	// }
	// // Возвести в квадрат
	// P1 := make([]float64, (NFFT/2 + 1))
	// for i := 0; i < len(P1); i++ {
	// 	P1[i] = 2 * P2[i]
	// }
	// return P2

	// Вычислите размер Pxx в зависимости от типа nfft.
	var PxxLen int
	if NFFT%2 == 0 {
		PxxLen = NFFT/2 + 1
	} else {
		PxxLen = (NFFT + 1) / 2
	}

	// Создайте слайс Pxx с правильной длиной.
	Pxx := make([]float64, PxxLen)
	alpha := 2.0

	for k := 0; k < PxxLen; k++ {
		freq := float64(k) / float64(NFFT)
		sinSum := 0.0
		cosSum := 0.0

		for n := 0; n < len(x); n++ {
			sinSum += x[n] * math.Sin(2*math.Pi*freq*float64(n))
			cosSum += x[n] * math.Cos(2*math.Pi*freq*float64(n))
		}

		Sk := (sinSum*sinSum + cosSum*cosSum) / float64(len(x))

		Pxx[k] = Sk / (alpha * float64(NW))
	}

	return Pxx
}
