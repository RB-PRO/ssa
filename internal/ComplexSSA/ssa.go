package complexssa

import (
	gomathtests "github.com/RB-PRO/ssa/pkg/go-MathTests"
	"github.com/RB-PRO/ssa/pkg/ssa"
)

func SSA(Folder string, pw []float64) error {
	s := ssa.New(Folder + "ssa")
	s.Graph = true // Создавать графики
	s.Xlsx = true  // Сохранять в Xlsx
	s.Var(pw, []float64{})
	s.Spw_Form(pw) // Создать spw

	// # 1, 2, 3, 4
	s.SET_Form() // SSA - анализ сегментов pw

	// # 5
	// Оценка АКФ сингулярных троек для сегментов pw
	// Визуализация АКФ сингулярных троек для сегментов pw
	s.AKF_Form() // Оценка АКФ сингулярных троек для сегментов pw

	// # 6, 7
	// Огибающие АКФ сингулярных троек sET12 сегментов pw
	// Нормированные АКФ сингулярных троек sET12 сегментов pw
	s.Envelope()

	// # 8
	// Мгновенная частота нормированной АКФ сингулярных троек sET12 для сегментов pw
	s.MomentFrequency()
	gomathtests.Plot(Folder+"smo.png", s.Smo_insFrc_AcfNrm)

	// # 9
	// Визуализация СПМ сингулярных троек сегменов pw
	s.VisibleSPM()

	// # 10
	// Агрегирование сегментов очищенной пульсовой волны cpw
	s.AggregationPW()
	return nil
}
