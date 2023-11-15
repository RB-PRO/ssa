package gomathtests

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Отрисовка графика с сохранением
func Plot(FilePathName string, opts ...[]float64) error {

	p := plot.New()

	p.Title.Text = FilePathName

	if len(opts) == 0 {
		return fmt.Errorf("nil plot data")
	}

	// Создание точек данных
	X := make([]float64, len(opts[0]))
	for i := range X {
		X[i] = float64(i)
	}
	for _, Y := range opts {

		pts := make(plotter.XYs, len(Y))
		for i := range pts {
			pts[i].X = X[i]
			pts[i].Y = Y[i]
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

	}

	// Установка названий осей
	// p.X.Label.Text = "ns"
	// p.Y.Label.Text = "insFrc_AcfNrm, Hz"

	// Сохранение графика в файл
	if ErrSave := p.Save(6*vg.Inch, 4*vg.Inch, FilePathName); ErrSave != nil {
		return ErrSave
	}

	return nil
}
