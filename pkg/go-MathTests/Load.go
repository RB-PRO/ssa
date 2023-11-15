package gomathtests

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Загрузить данные из файла
func LoadSlice(FilePathName string) (data []float64, Err error) {
	// Если грузим в txt
	if strings.Contains(FilePathName, ".txt") {
		file, ErrOpen := os.Open(FilePathName)
		if ErrOpen != nil {
			return nil, fmt.Errorf("os.Open: %v", ErrOpen)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fl, ErrParseFloat := strconv.ParseFloat(scanner.Text(), 64)
			if ErrParseFloat != nil {
				return nil, fmt.Errorf("strconv.ParseFloat: %v", ErrParseFloat)
			}
			data = append(data, fl)
		}
		if Err = scanner.Err(); Err != nil {
			return nil, fmt.Errorf("scanner.Err(): %v", Err)
		}
	}

	// // Если грузим в xlsx
	// if strings.Contains(FilePathName, ".xlsx") {

	// }
	return data, Err
}

// Загрузить RGB сигонал
func LoadRGB(FileName string) (R, G, B []float64, Err error) {
	file, ErrOpen := os.Open(FileName)
	if ErrOpen != nil {
		return nil, nil, nil, fmt.Errorf("os.Open: %v", ErrOpen)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()
		strs := strings.Split(str, ";")
		if len(strs) == 3 {
			strs[0] = strings.ReplaceAll(strs[0], ",", ".")
			strs[1] = strings.ReplaceAll(strs[1], ",", ".")
			strs[2] = strings.ReplaceAll(strs[2], ",", ".")

			r, _ := strconv.ParseFloat(strs[0], 64)
			g, _ := strconv.ParseFloat(strs[1], 64)
			b, _ := strconv.ParseFloat(strs[2], 64)

			R = append(R, r)
			G = append(R, g)
			B = append(R, b)
		}
	}
	if ErrScanner := scanner.Err(); ErrScanner != nil {
		return nil, nil, nil, fmt.Errorf("scanner.Err: %v", ErrScanner)
	}
	return R, G, B, nil
}

func Save(FileName string, opts ...[]float64) error {
	filePW, ErrOpenFile := os.Create(FileName)
	if ErrOpenFile != nil {
		panic(ErrOpenFile)
	}
	defer filePW.Close()

	// var content string
	for _, slice := range opts {
		// fmt.Println(slice)
		for i := range slice {
			strs := make([]string, len(opts))
			for j := range opts {
				strs[j] = fmt.Sprintf("%.8f", opts[j][i])
			}
			fmt.Println()
			if _, ErrWriteString := filePW.WriteString(strings.Join(strs, ";") + "\n"); ErrWriteString != nil {
				return fmt.Errorf("filePW.WriteString: %v", ErrWriteString)
			}
		}
	}
	return nil
}
