package tg

import (
	"fmt"
	"math"
	"sort"

	"github.com/RB-PRO/ssa/pkg/oss"
)

func CalcPW(RGBs []RGB_float64, Path string) (pw []float64, Err error) {
	pw = make([]float64, len(RGBs))
	pw2 := make([]float64, len(RGBs))

	// CR
	for i := range RGBs {
		pw[i] = (RGBs[i].R*112.0 -
			RGBs[i].G*93.8 -
			RGBs[i].B*18.2) / 255.0
		pw2[i] = math.Pow(pw[i], 2)
		// if _, err := filePW.WriteString(fmt.Sprintf("%.8f\n", pw[i])); err != nil {
		// 	log.Println(err)
		// }
	}

	// fMi := 40.0 / 60.0
	// cad := 30
	// SMO_med := cad / int(fMi)

	DEV_med := medianFilter(pw2, 30*60/40)
	createLineChart([]float64{}, DEV_med, Path+"DEV_med.png")
	// filePWmedianFilter, _ := os.Create(Path + "medianFilter.txt")
	// for i := range DEV_med {
	// 	if _, err := filePWmedianFilter.WriteString(fmt.Sprintf("%.8f\n", DEV_med[i])); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	for i := range pw {
		pw[i] /= math.Sqrt(DEV_med[i])
	}
	prcMi := oss.Prctile(pw, 0.1)
	prcMa := oss.Prctile(pw, 99.9)
	prcMi = -0.0197
	prcMa = 0.0207
	fmt.Println("prcMi", prcMi, "prcMa", prcMa)
	createLineChart([]float64{}, pw, Path+"pw2.png")
	for i := range pw {
		if pw[i] < prcMi {
			pw[i] = prcMi
		}
		if pw[i] > prcMa {
			pw[i] = prcMa
		}
	}
	createLineChart([]float64{}, pw, Path+"pw.png")

	return pw, nil
}

func medianFilter(x []float64, n int) []float64 {
	// Проверка на нечетное значение n
	if n%2 == 0 {
		n++
	}

	// Длина входного массива
	length := len(x)

	// Результат фильтрации
	y := make([]float64, length)

	for i := 0; i < length; i++ {
		// Индексы для сбора значений для медианного фильтра
		start := i - n/2
		end := i + n/2

		// Гарантия, что индексы не выходят за пределы массива
		if start < 0 {
			start = 0
		}
		if end >= length {
			end = length - 1
		}

		// Извлечение значений для медианы
		window := x[start : end+1]

		// Сортировка окна значений и выбор медианы
		sortedWindow := make([]float64, len(window))
		copy(sortedWindow, window)
		sort.Float64s(sortedWindow)
		// Хитрый мув. При делении int(5) на int(2), получается int(2),
		// т.е. округление в нижнюю сторону, хотя нам нужно в старшую степень.
		// Поэтому из нечётного делаем чётное, а в случае получения нечётного не имеет разницы
		medianIndex := (len(sortedWindow) + 1) / 2
		// fmt.Println("medianIndex", medianIndex, "-", sortedWindow)
		y[i] = sortedWindow[medianIndex]
	}

	return y
}
