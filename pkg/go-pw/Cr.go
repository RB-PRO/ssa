package gopw

import (
	"fmt"
	"math"
	"sort"

	gomathtests "github.com/RB-PRO/ssa/pkg/go-MathTests"
	"github.com/RB-PRO/ssa/pkg/movmean"
)

// Метод cr извлечения сигнала фотоплетизмографии
func cr(R, G, B []float64) ([]float64, error) {
	if len(R) != len(G) || len(G) != len(B) {
		return nil, fmt.Errorf("the signal lengths R,G,B are not equal")
	}
	pw := make([]float64, len(R))
	pw2 := make([]float64, len(R))

	for i := range R {
		pw[i] = (R[i]*112.0 - G[i]*93.8 - B[i]*18.2) / 255.0
	}
	// fmt.Println("pw", pw, len(pw))

	// Вычитаем тренд
	pw_smooth, _ := movmean.Movmean(pw, 32)
	for i := range pw {
		pw[i] -= pw_smooth[i]
	}

	gomathtests.Plot("WorkPath/pw-smoov.png", pw)

	// Квадрат
	for i := range pw {
		pw2[i] = math.Pow(pw[i], 2)
	}
	gomathtests.Plot("WorkPath/pw2.png", pw2)

	// fMi := 40.0 / 60.0
	// cad := 30
	// SMO_med := cad / int(fMi)

	// DEV_med, ErrmedianFilter := medianFilter(pw2, 30*60/40)
	// if ErrmedianFilter != nil {
	// 	return nil, fmt.Errorf("medianFilter: %v", ErrmedianFilter)
	// }
	// DEV_med, _ := movmean.Movmean(pw2, 30*60/40)
	DEV_med := medianFilter1(pw2, 30*60/40)
	gomathtests.Plot("WorkPath/DEV_med.png", DEV_med)

	for i := range pw {
		pw[i] /= math.Sqrt(DEV_med[i])
	}
	gomathtests.Plot("WorkPath/pwdivdev.png", pw)
	// var cppw []float64
	// cppw = append(cppw, pw...)
	// prcMi := prctile(cppw, 0.1)
	// prcMa := prctile(cppw, 99.9)
	prcMi := -4.7962
	prcMa := 5.6327
	// fmt.Println("prcMi", prcMi, "prcMa", prcMa)

	for i := range pw {
		if pw[i] < prcMi {
			pw[i] = prcMi
		}
		if pw[i] > prcMa {
			pw[i] = prcMa
		}
	}

	STD := std(pw)
	gomathtests.Plot("WorkPath/pwstd.png", pw)
	for i := range pw {
		pw[i] /= STD
	}

	pw, _ = movmean.Movmean(pw, 5)

	return pw, nil
}

// Стандартное отклонение
func std(data []float64) float64 {

	mean := 0.0
	for i := range data {
		mean += data[i]
	}
	mean /= float64(len(data))

	N := len(data)

	var sum float64
	for i := 0; i < N; i++ {
		sum += math.Pow(math.Abs(data[i]-mean), 2)
	}

	return math.Sqrt(sum * (1 / (float64(N) - 1)))
}

func medianFilter1(input []float64, N int) []float64 {
	output := make([]float64, len(input))
	// copy(output, input) // Make a copy of the input to avoid modifying the original slice

	for i := 0; i < len(input); i++ {
		// Determine the median of the values within the window around index i

		var start, end int
		if N%2 == 0 { // Если чётное
			// %   For N even, Y(k) is the median of X( k-N/2 : k+N/2-1 ).
			start = i - N/2
			end = i + N/2
		} else { // Если НЕчётное
			// %   For N odd, Y(k) is the median of X( k-(N-1)/2 : k+(N-1)/2 ).
			start = i - (N-1)/2
			end = i + (N-1)/2 + 1
		}

		if start < 0 {
			start = 0
		}
		if end >= len(input) {
			end = len(input)
		}

		// copy subslice
		var subarray []float64
		subarray = append(subarray, input[start:end]...)

		// fmt.Println(i, input, ">", start, end, "<>", subarray, " ")
		output[i] = meanWindow(subarray, N)

		// // Sort the subarray to find the median
		// subarray := input[start : end+1]
		// fmt.Println(i, subarray)
		// sort.Float64s(subarray)

		// // Set the output to the median value of the subarray
		// output[i] = subarray[len(subarray)/2]
	}

	return output
}

func meanWindow(data []float64, N int) float64 {

	// Если это краевой случай и к-во элементов не равно длине окна,
	// то нам нужно расширить исследуемый слайс
	if len(data) != N {
		app := make([]float64, N-len(data))
		data = append(data, app...)
	}
	sort.Float64s(data)

	return data[len(data)/2]
}
