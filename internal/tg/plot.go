package tg

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func createLineChart(x, y []float64, FileName string) error {

	p := plot.New()

	p.Title.Text = "Нормированные АКФ сингулярных троек sET12 сегментов pw"

	// Создание точек данных
	pts := make(plotter.XYs, len(y))
	xx := make([]float64, len(y))
	for i := range xx {
		xx[i] = float64(i)
	}

	for i := range pts {
		pts[i].X = xx[i]
		pts[i].Y = y[i]
	}

	// Создание линейного графика
	line, ErrNewLine := plotter.NewLine(pts)
	if ErrNewLine != nil {
		return ErrNewLine
	}
	line.LineStyle.Width = vg.Points(1)
	line.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Добавление графика к графическому контексту
	p.Add(line)

	// Установка названий осей
	p.X.Label.Text = "ns"
	p.Y.Label.Text = "insFrc_AcfNrm, Hz"

	// Сохранение графика в файл
	if ErrSave := p.Save(6*vg.Inch, 4*vg.Inch, FileName); ErrSave != nil {
		return ErrSave
	}

	return nil
}
