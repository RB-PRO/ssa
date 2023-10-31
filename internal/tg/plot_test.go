package tg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/RB-PRO/ssa/pkg/oss"
)

func TestPlot(t *testing.T) {
	createLineChart([]float64{1, 2, 3, 4, 5}, []float64{2, 3, 4, 5, 6}, "Sample.png")
}

func TestSSA_tgbot(t *testing.T) {
	Folder := "..\\..\\TelegramVideoNote/test/"
	pw, _ := oss.Make_singnal_xn(Folder + "pw")
	SSA_tgbot(Folder, pw)
}

func TestPW(t *testing.T) {
	var RGBs []RGB_float64
	file, err := os.Open("..\\..\\TelegramVideoNote/test/RGB.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		strs := strings.Split(str, ";")
		if len(strs) == 3 {
			R, _ := strconv.ParseFloat(strs[0], 64)
			G, _ := strconv.ParseFloat(strs[1], 64)
			B, _ := strconv.ParseFloat(strs[2], 64)
			RGBs = append(RGBs, RGB_float64{
				R: R,
				G: G,
				B: B,
			})
		}
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	pw, _ := CalcPW(RGBs, "..\\..\\TelegramVideoNote/test/")
	createLineChart([]float64{}, pw, "..\\..\\TelegramVideoNote/test/"+"pw.png")

}

func TestSSA(t *testing.T) {
	var pw []float64
	file, err := os.Open("..\\..\\TelegramVideoNote/test/pw.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		strfloat64, _ := strconv.ParseFloat(str, 64)
		pw = append(pw, strfloat64)
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	SSA_tgbot("..\\..\\TelegramVideoNote/test/", pw)
}

func TestMedianFilter(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Пример входных данных
	n := 3                                        // Порядок фильтра
	result := medianFilter(x, n)                  // Применение медианного фильтра
	fmt.Println(result)                           // Вывод результата
}

func TestExtractRGB(t *testing.T) {
	RGBs, errRGB := ExtractRGB("..\\..\\TelegramVideoNote/test/",
		"huawei.mp4")
	if errRGB != nil {
		t.Error(errRGB)
	}
	fmt.Println("len(RGBs)", len(RGBs))
}
