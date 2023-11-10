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
	pw := LoadTXT("EUT_P1H1_RGB_pw.txt") // "P1H1_edited_pw.txt"
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
	filePW, ErrOpenFile := os.Create("pto2.txt")
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

func LoadTXT(FileName string) []float64 {
	var data []float64
	file, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fl, _ := strconv.ParseFloat(scanner.Text(), 64)
		data = append(data, fl)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return data
}
