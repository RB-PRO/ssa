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

	fmt.Println(1)
}
