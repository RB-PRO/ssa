package gopw

import (
	"errors"
	"sort"
)

// Медианная фильрация
func medianFilter(x []float64, n int) ([]float64, error) {

	// Сравнения окна и количества отсчётов в сигнале
	if n > len(x) {
		return nil, errors.New("window N more lenght signal")
	}

	// Проверка на нечетное значение n
	// в противном случае из нечётного делаем чётное
	// %   For N odd, Y(k) is the median of X( k-(N-1)/2 : k+(N-1)/2 ).
	// %   For N even, Y(k) is the median of X( k-N/2 : k+N/2-1 ).
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
		// т.е. округление в нижнюю сторону, хотя нам нужно в старшую степень
		// Поэтому из нечётного делаем чётное,
		// а в случае получения нечётного не имеет разницы
		medianIndex := (len(sortedWindow) + 1) / 2

		y[i] = sortedWindow[medianIndex] // Сохравнение результата
	}

	return y, nil
}
