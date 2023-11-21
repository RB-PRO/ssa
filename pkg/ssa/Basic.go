package ssa

import (
	"math"
	"slices"

	"gonum.org/v1/gonum/mat"
)

// Структура с данными спектрального анализа
type SPW struct {
	N       int // Количество отсчетов pw
	Win     int // Размер окна
	Res     int
	NPart   int // Количество долей res
	Overlap int
	S       int
	Imin    int
	Imax    int
	Ns      []float64
	NSF     int // номер финального отсчета финального сегмента <= N

	// Настройки конфигурации
	Graph bool // Создавать графики, где true - создавать
	Xlsx  bool // Сохранять графики в Xlsx, где true - создавать

	// Set general parameters
	Cad int // 30 кадров/сек
	Dt  float64
	Tim []float64 // Set general parameters
	Ns2 []int
	L   []float64

	lag  int // к-во отсчётов
	lgl  []float64
	time []float64 // Слайс со временем

	K int
	M int // параметр вложения в траекторное пространство
	// SSA - анализ сегментов pw
	Seg int // номер сегмента pw для визуализации
	NET int // кол-во сингулярных троек для сегментов pw

	InsFrc_AcfNrm     []float64 // Это для сохранения графика
	Smo_insFrc_AcfNrm []float64 // Это для сохранения графика

	Dir *Direction

	Spw *mat.Dense

	SET12 *mat.Dense
	SET34 *mat.Dense

	Acf_sET12 *mat.Dense

	EnvAcf_sET12 *mat.Dense
	AcfNrm_sET12 *mat.Dense
}
type Direction struct {
	zeropath string // Путь к рабочей папке
}

func New(Path string) *SPW {
	Dir, _ := NewDirection(Path)
	spw := SPW{Dir: Dir}
	return &spw
}

// Инициализируем сигнал pw
func (s *SPW) Var(pw, fmp []float64) *SPW {
	s.N = len(pw)
	s.Win = 1024
	s.Res = s.N - s.Win*int(math.Floor(float64(s.N)/float64(s.Win)))
	s.NPart = 20 // Количество долей res
	s.Res = int(math.Floor(float64(s.Res) / float64(s.NPart)))
	// fmt.Println("s.Res", s.Res)
	//s.Res = 40
	//overlap := (float64(win) - float64(res)) / float64(win)
	s.S = 1
	s.Imin = 1
	s.Imax = s.Win

	for s.Imax <= s.N {
		s.Ns = append(s.Ns, float64(s.S)) // номер текущего сегмента pw
		s.Imin += s.Res
		s.Imax += s.Res
		s.S++
	}
	s.S-- // кол-во перекрывающихся сегментов pw в пределах N

	// Различные глобальные переменные
	s.Cad = 30                // 30 кадров/сек
	s.Dt = 1 / float64(s.Cad) // интервал дискретизации времени, сек
	tim := make([]float64, s.N)
	for index := 1; index < s.N; index++ {
		tim[index] = tim[index-1] + s.Dt
	}
	s.Tim = tim

	ns2 := make([]int, s.S)
	for index := range ns2 {
		ns2[index] = (index + 1)
	}
	s.Ns2 = ns2

	L := make([]float64, s.S)
	for index := 0; index < len(L); index++ {
		// for index := range L { // цикл по сегментам pw
		// L[index] = math.Floor(float64(s.Cad) / fmp[index]) // кол-во отсчетов основного тона pw
		L[index] = float64(s.Cad) / 1.5
	}
	s.L = L

	s.K = 5
	s.M = int(float64(s.K) * slices.Max(s.L)) // параметр вложения в траекторное пространство

	// SSA - анализ сегментов pw
	s.Seg = 100 // номер сегмента pw для визуализации
	s.NET = 4   // кол-во сингулярных троек для сегментов pw

	return s
}

func (s *SPW) Spw_Form(pw []float64) *SPW {
	s.Spw = mat.NewDense(s.Win, s.S, nil)
	for j := 0; j < s.S; j++ {
		for i := 0; i < s.Win; i++ {
			k := j * s.Res
			s.Spw.Set(i, j, pw[k+i])
		}
	}
	return s
}
