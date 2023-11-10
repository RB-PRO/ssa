package ssa2

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"

	"github.com/RB-PRO/ssa/pkg/ssa2/pmtm"
)

// Структура загрузчика данных
type Setup struct {
	Cad   int     // Количество кадров в секунду
	NPart int     // Количество долей res
	Win   int     // Ширина окна
	FMi   float64 // Частота среза для 40 уд/мин (0.6667 Гц)
	FMa   float64 // Частота среза для 240 уд/мин (4.0 Гц)
}

type SSA struct {
	set Setup // Структура настройки Расчётов SSA

	len int // Количество отсчётов pw

	// Временные интервалы
	tim []float64 // Время в секундах
	dt  float64   // Интервал дискретизации времени, сек

	col int // Всего колонок

	// Сигнал фотоплетизмографии
	pw       []float64 // Сигнал фотоплетизмографии
	freq     []float64
	Pto_fMAX []float64

	// Оценки СПМ перекрывающихся сегменов pw
	iGmin int
	iGmax int

	res int
}

// Создать объект SSA для работы
func NewSSA(pw []float64, Setting Setup) (*SSA, error) {
	var ssa SSA
	ssa.pw = pw
	ssa.set = Setting

	// к-во отсчётов пульсовой волны
	ssa.len = len(ssa.pw)

	// интервал дискретизации времени, сек
	ssa.dt = 1.0 / float64(ssa.set.Cad)

	// время в секундах
	tim := make([]float64, ssa.len)
	for i := 0; i < ssa.len; i++ {
		tim[i] = float64(i) * ssa.dt
	}
	ssa.tim = tim

	return &ssa, nil
}

// Расчёт параметра col
func (ssa *SSA) Col() (*SSA, error) {
	col := 1
	Imin := 1
	Imax := ssa.set.Win
	resFloat64 := float64(ssa.len) - float64(ssa.set.Win)*math.Floor(float64(ssa.len)/float64(ssa.set.Win))
	res := int(math.Floor(resFloat64 / float64(ssa.set.NPart)))

	// col - кол-во перекрывающихся сегментов в пределах len
	for Imax <= ssa.len {
		Imin = Imin + res
		Imax = Imax + res
		col = col + 1
	}
	col--

	ssa.col = col
	ssa.res = res
	return ssa, nil
}

// текущий сегмент pw длинною win
func (ssa *SSA) Spw(column int) ([]float64, error) {
	if ssa.col == 0 {
		return nil, fmt.Errorf("spw: ssa.col is 0")
	}
	if column >= ssa.col {
		return nil, fmt.Errorf("spw: Условие column(%d) < ssa.col(%d) не выполнено", column, ssa.col)
	}
	if len(ssa.pw) == 0 {
		return nil, fmt.Errorf("spw: len(pw)=0")
	}
	return ssa.pw[column*ssa.res : column*ssa.res+ssa.set.Win], nil
}

// Оценки СПМ перекрывающихся сегменов pw
func (ssa *SSA) SpwEstimation() (*SSA, error) {
	df := float64(ssa.set.Cad) / float64(ssa.set.Win-1)
	Fmin := ssa.set.FMi - 10*df
	Fmax := ssa.set.FMa + 10*df
	row := 1 + ssa.set.Win/2

	tecalF := 0.0
	ssa.freq = make([]float64, row)
	for i := 0; i < row; i++ {
		if math.Abs(tecalF-Fmin) <= df {
			ssa.iGmin = i
		}
		if math.Abs(tecalF-Fmax) <= df {
			ssa.iGmax = i
		}
		ssa.freq[i] = tecalF
		tecalF += df
	}
	SaveTXT("ssafreq.txt", ssa.freq)

	return ssa, nil
}

// Оценки средних частот основного тона для сегментов pw
func (ssa *SSA) PwEstimation() (*SSA, error) {
	// BlackManHar := blackmanharris.Blackmanharris(ssa.set.Win, blackmanharris.Koef4_74db)

	filePW, _ := os.Create("tests/pto.txt")
	defer filePW.Close()

	// Выделяем память в слайс
	ssa.Pto_fMAX = make([]float64, ssa.col)
	for i := 0; i < ssa.col; i++ {

		// Получение колонки SPW
		ColumnSPW, ErrSpw := ssa.Spw(i)
		if ErrSpw != nil {
			return ssa, fmt.Errorf("SSA: PwEstimation: %w", ErrSpw)
		}
		// ColumnSPW := make([]float64, len(ColumnSPW2))
		// copy(ColumnSPW, ColumnSPW2)

		// Получение периодограммы колонки SPW
		// pg_spw := periodogram.Periodogram(ColumnSPW, BlackManHar, ssa.set.Win)

		pto_spw := pmtm.Pmtm(ColumnSPW, 3, ssa.set.Win)

		for j := range pto_spw {
			filePW.WriteString(fmt.Sprintf("%.8f;", pto_spw[j]))
		}
		filePW.WriteString(fmt.Sprintf("\n"))
		pto_spw = medianFilter(pto_spw, 30)
		// pto_spw = sredFilter(pto_spw, 15)

		// Сотртируем, чтобы получить по возрастанию порядковвые номера в слайсе по убыванию
		// _, SorterIndexts_pg_spw := InsertionSort(pto_spw)
		SorterIndexts_pg_spw := InsertionSort2(pto_spw)
		// _, SorterIndexts_pg_spw := InsertionSortInt(pto_spw)

		// fmt.Println(SorterIndexts_pg_spw[0], len(pto_spw))
		ssa.Pto_fMAX[i] = ssa.freq[SorterIndexts_pg_spw] // float64(SorterIndexts_pg_spw)
		// ssa.Pto_fMAX[i] = float64(SorterIndexts_pg_spw)
	}

	return ssa, nil
}

func SaveTXT(FileName string, data []float64) {
	filePW, ErrOpenFile := os.Create(FileName)
	if ErrOpenFile != nil {
		panic(ErrOpenFile)
	}
	defer filePW.Close()
	for i := range data {
		if _, err := filePW.WriteString(fmt.Sprintf("%.8f\n", data[i])); err != nil {
			log.Println(err)
		}
	}
}

// Сортировка с возвратом номеров изначальных элементов
func InsertionSort(array []float64) ([]float64, []int) {
	// var indexArray int
	// sort.Slice(array, func(i, j int) bool {
	// 	if array[i] > array[j] {
	// 		indexArray = j
	// 	}
	// 	return array[i] > array[j]
	// })
	// fmt.Println(indexArray, array[0], array[1], array[2])

	indexArray := make([]int, len(array))
	for ind := range indexArray {
		indexArray[ind] = (ind)
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

// Сортировка с возвратом номеров изначальных элементов
func InsertionSortInt(array []float64) ([]float64, int) {
	indexArrayint := 0
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 {
			if array[j-1] < array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				indexArrayint = i
			}
			j = j - 1
		}
	}
	return array, indexArrayint
}

// Сортировка с возвратом номеров изначальных элементов
func InsertionSort2(array []float64) int {
	return slices.Index(array, slices.Max(array))
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

func sredFilter(x []float64, n int) []float64 {
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
		// slises
		sum := 0.0
		for _, v := range window {
			sum += v
		}
		y[i] = sum / float64(n)

	}

	return y
}
