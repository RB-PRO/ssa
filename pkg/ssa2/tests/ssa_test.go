package ssa2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/RB-PRO/ssa/pkg/ssa2"
)

func TestSSA(t *testing.T) {
	var pw []float64
	file, err := os.Open("EUT_P1H1.txt") // "P1H1_edited_pw.txt"
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		strfloat64, _ := strconv.ParseFloat(str, 64)
		pw = append(pw, strfloat64)
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}
	ssaAnalis, ErrNewSSA := ssa2.NewSSA(pw, ssa2.Setup{
		Cad:   30,
		Win:   1024,
		NPart: 20,
		FMi:   40.0 / 60.0,
		FMa:   240.0 / 60.0,
	})
	if ErrNewSSA != nil {
		t.Error(ErrNewSSA)
	}
	ssaAnalis.Col()
	col, ErrSpw := ssaAnalis.Spw(1)
	if ErrSpw != nil {
		t.Error(ErrSpw)
	}
	fmt.Println(col[0], col[1], col[2], len(col))
	ssaAnalis.SpwEstimation()
	ssaAnalis.PwEstimation()
	ssa2.CreateLineChart(ssaAnalis.Pto_fMAX, "pto.png")
	filePW, ErrOpenFile := os.Create("pto.txt")
	if ErrOpenFile != nil {
		t.Error(ErrOpenFile)
	}
	defer filePW.Close()
	for i := range ssaAnalis.Pto_fMAX {
		if _, err := filePW.WriteString(fmt.Sprintf("%.8f\n", ssaAnalis.Pto_fMAX[i])); err != nil {
			log.Println(err)
		}
	}
}
