package tg

import (
	"fmt"
	"log"
	"os"
)

func CalcPW(RGBs []RGB_float64, Path string) (pw []float64, Err error) {
	pw = make([]float64, len(RGBs))

	filePW, ErrOpenFile := os.OpenFile(Path+"pw.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if ErrOpenFile != nil {
		return nil, ErrOpenFile
	}
	defer filePW.Close()
	for i := range RGBs {
		pw[i] = (RGBs[i].R*112.0 -
			RGBs[i].G*93.8 -
			RGBs[i].B*18.2) / 255.0
		if _, err := filePW.WriteString(fmt.Sprintf("%.8f\n", pw[i])); err != nil {
			log.Println(err)
		}
	}
	return pw, nil
}
