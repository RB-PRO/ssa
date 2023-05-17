package graph_test

import (
	"testing"

	"github.com/RB-PRO/ssa/pkg/graph"
)

func Test3D(t *testing.T) {
	graph.TreeXDXD([]float64{}, []float64{}, []float64{})
}

func TestSplotMatrixFromDat(t *testing.T) {
	Gragh7 := graph.Option3D{ // Задаём настройки 3D графика
		FileNameDat: "tests/" + "AcfNrm_sET12.dat",
		FileNameOut: "tests/" + "AcfNrm_sET12.png",
		Titile:      "Нормированные АКФ сингулярных троек sET12 сегментов pw",
		Xlabel:      "ns",
		Ylabel:      "lag,s",
		Zlabel:      "Acf_Nrm",
	}
	graph.SplotMatrixFromFile(Gragh7) // Делаем график
}
