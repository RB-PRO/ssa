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
	"github.com/RB-PRO/ssa/pkg/ssa"
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
	RGBs := LoadRGB("..\\..\\TelegramVideoNote/test/RGBTEST.txt")
	pw, _ := CalcPW(RGBs, "..\\..\\TelegramVideoNote/test/")
	createLineChart([]float64{}, pw, "..\\..\\TelegramVideoNote/test/"+"pw.png")
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
func LoadRGB(FileName string) []RGB_float64 {
	var RGBs []RGB_float64
	file, err := os.Open(FileName)
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
		panic(err)
	}
	return RGBs
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
	RGBs, errRGB := ExtractRGB("..\\..\\TelegramVideoNote/test/", "huawei.mp4")
	if errRGB != nil {
		t.Error(errRGB)
	}
	fmt.Println("len(RGBs)", len(RGBs))
}

func TestButter(t *testing.T) {
	FileName := "P1H1_edited_RGB.txt"
	FileName = strings.ReplaceAll(FileName, ".avi", "")
	FileName = strings.ReplaceAll(FileName, ".txt", "")
	FileName = strings.ReplaceAll(FileName, "_RGB", "")
	FileName = strings.ReplaceAll(FileName, "_pw", "")
	FileName = strings.ReplaceAll(FileName, "_but", "")
	Folder := "tests/"

	// RGB
	// RGB, _ := ExtractRGB(Folder, FileName+".avi")
	// RGB := LoadRGB(Folder + FileName + ".txt")
	// Butter
	// RGB_but := Butter(RGB)
	// SaveRGB(Folder+FileName+"_but.txt", RGB_but)
	// RGB2s := LoadRGB(Folder + FileName + ".txt")
	RGB := LoadRGB(Folder + FileName + "_RGB" + ".txt")
	// // PW
	pw, ErrPW := CalcPW(RGB, Folder)
	if ErrPW != nil {
		t.Error(ErrPW)
	}
	// SaveTXT(Folder+FileName+"_pw.txt", pw)
	// pw = LoadTXT(Folder + "tg10" + "_pw.txt")

	SSS_spw2(pw, []float64{})

	// Частоты

	// ssaAnalis, ErrNewSSA := ssa2.NewSSA(pw, ssa2.Setup{
	// 	Cad: 30, Win: 1024, NPart: 20,
	// 	FMi: 40.0 / 60.0, FMa: 240.0 / 60.0,
	// })
	// if ErrNewSSA != nil {
	// 	t.Error(ErrNewSSA)
	// }
	// ssaAnalis.Col()
	// ssaAnalis.SpwEstimation()
	// ssaAnalis.PwEstimation()
	// ssa2.CreateLineChart(ssaAnalis.Pto_fMAX, Folder+"pto.png")
}

func SSS_spw2(pw, fmp []float64) {
	s := ssa.New("File")
	s.Graph = true // Создавать графики
	s.Xlsx = true  // Сохранять в Xlsx
	s.Init(pw, fmp)
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

	// # 9
	// Визуализация СПМ сингулярных троек сегменов pw
	s.VisibleSPM()

	// # 10
	// Агрегирование сегментов очищенной пульсовой волны cpw
	s.AggregationPW()

}

// func LoadTXT()

func SaveTXT(FileName string, data []float64) {
	filePW, ErrOpenFile := os.Create(FileName)
	if ErrOpenFile != nil {
		panic(ErrOpenFile)
	}
	defer filePW.Close()
	for i := range data {
		if _, err := filePW.WriteString(fmt.Sprintf("%.8f\n", data[i])); err != nil {
			log.Println(err)
		}
	}
}
func SaveRGB(FileName string, data []RGB_float64) {
	filePW, ErrOpenFile := os.Create(FileName)
	if ErrOpenFile != nil {
		panic(ErrOpenFile)
	}
	defer filePW.Close()
	for i := range data {
		if _, err := filePW.WriteString(fmt.Sprintf("%.8f;%.8f;%.8f\n", data[i].R, data[i].G, data[i].B)); err != nil {
			log.Println(err)
		}
	}
}
