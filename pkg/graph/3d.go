package graph

import (
	"fmt"

	"github.com/sbinet/go-gnuplot"
)

func TreeXDXD(x, y, z []float64) {

	fname := ""
	persist := false
	debug := true

	p, err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.CheckedCmd("plot %f*x", 23.0)
	p.CheckedCmd("plot %f * cos(%f * x)", 32.0, -3.0)
	//p.CheckedCmd("save foo.ps")
	p.CheckedCmd("set terminal png")
	p.CheckedCmd("set output 'plot001.png'")
	p.CheckedCmd(`set terminal png size 1280,720 enhanced font "Helvetica,20"`)
	p.CheckedCmd("set output 'output.png'")
	p.CheckedCmd("replot")
	p.CheckedCmd("q")

}

// Стуктура-настройка для реализации графика
type Option3D struct {
	FileNameDat string // Файл формата .dat, который содержит данные, которые необходимо построить
	FileNameOut string // Файл формата .png, который будет сохранён
	Titile      string // Подпись графика
	Xlabel      string // Подпись оси X
	Ylabel      string // Подпись оси Y
	Zlabel      string // Подпись оси Z
}

// Сделать 3D mash
func SplotMatrixFromFile(opt Option3D) error {

	// Запускаем термина
	p, err := gnuplot.NewPlotter("", false, false)
	if err != nil {
		return err
	}
	defer p.Close()

	p.CheckedCmd(`set terminal png size 1024,768 font "Helvetica,15.0"`)
	p.CheckedCmd(`set output "` + opt.FileNameOut + `"`)

	p.CheckedCmd(`splot "` + opt.FileNameDat + `" matrix w l`)
	p.CheckedCmd("set pm3d")
	p.CheckedCmd("unset surface")
	p.CheckedCmd(`set title "` + opt.Titile + `"`)
	p.CheckedCmd(`set xlabel "` + opt.Xlabel + `"`)
	p.CheckedCmd(`set ylabel "` + opt.Ylabel + `"`)
	p.CheckedCmd(`set zlabel "` + opt.Zlabel + `"`)

	p.CheckedCmd(`set terminal png size 1024,768 font "Helvetica,15.0"`)
	p.CheckedCmd(`set output "` + opt.FileNameOut + `"`)
	p.CheckedCmd("replot")

	p.CheckedCmd("set view map")

	p.CheckedCmd(`set terminal png size 1024,768 font "Helvetica,15.0"`)
	p.CheckedCmd(`set output "` + opt.FileNameOut + `"`)
	p.CheckedCmd("replot")

	return nil
}
