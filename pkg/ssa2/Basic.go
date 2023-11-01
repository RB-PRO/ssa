package ssa2

import (
	"fmt"
	"math"

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

	return ssa, nil
}

// Оценки средних частот основного тона для сегментов pw
func (ssa *SSA) PwEstimation() (*SSA, error) {
	// BlackManHar := blackmanharris.Blackmanharris(ssa.set.Win, blackmanharris.Koef4_74db)

	// Выделяем память в слайс
	ssa.Pto_fMAX = make([]float64, ssa.col)
	for i := 0; i < ssa.col; i++ {

		// Получение колонки SPW
		ColumnSPW, ErrSpw := ssa.Spw(i)
		if ErrSpw != nil {
			return ssa, fmt.Errorf("SSA: PwEstimation: %w", ErrSpw)
		}

		// Получение периодограммы колонки SPW
		// pg_spw := periodogram.Periodogram(ColumnSPW, BlackManHar, ssa.set.Win)

		pto_spw := pmtm.Pmtm(ColumnSPW, ssa.res)

		// Сотртируем, чтобы получить по возрастанию порядковвые номера в слайсе по убыванию
		_, SorterIndexts_pg_spw := InsertionSort(pto_spw)

		ssa.Pto_fMAX[i] = ssa.freq[SorterIndexts_pg_spw[0]]
	}

	return ssa, nil
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
