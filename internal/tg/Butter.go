package tg

import (
	"fmt"
	"math"

	"github.com/jfcg/butter"
)

// Нормирование слайса RGB
//
// Сейчас работает очень неэффективно, требуется модификация
func Butter(RGB []RGB_float64) []RGB_float64 {

	dt := 1. / 30.0
	wc := 2 * math.Pi * dt

	// fl := []butter.Filter{
	// 	// butter.NewHighPass1(4 * wc),
	// 	// butter.NewLowPass1(0.6 * wc),
	// 	// butter.NewHighPass2(4 * wc),
	// 	// butter.NewLowPass2(0.6 * wc),
	// 	// butter.NewBandPass2(wc*0.6, wc*4),
	// 	butter.NewBandPass2(wc*0.3, wc*2),
	// 	// butter.NewBandStop2(wc*0.6, wc*2'),
	// 	// butter.NewBandStop2(wc*0.6, wc*4),
	// 	// butter.NewRateLimit(0, wc),
	// }
	a, b := wc*1.2, wc*2.5
	Rfilter := butter.NewBandPass2(a, b)
	Gfilter := butter.NewBandPass2(a, b)
	Bfilter := butter.NewBandPass2(a, b)

	for i := range RGB {
		RGB[i].R = Rfilter.Next(RGB[i].R)
		RGB[i].G = Gfilter.Next(RGB[i].G)
		RGB[i].B = Bfilter.Next(RGB[i].B)
	}
	return RGB
}

////////////////

func ButterworthLowpassFilter(signal []float64, sampleRate float64, cutoffFreq float64) ([]float64, error) {
	// Вычислите параметры фильтра
	order := 4 // Порядок фильтра (вы можете выбрать другой порядок)
	nyquistFreq := 0.5 * sampleRate
	normalCutoff := cutoffFreq / nyquistFreq

	// Создайте нормализованный фильтр Баттерворта
	b, a, err := butterworthLowpass(order, normalCutoff)
	if err != nil {
		return nil, err
	}

	// Примените фильтр к сигналу
	filteredSignal := make([]float64, len(signal))
	for i := 0; i < len(signal); i++ {
		for j := 0; j <= order; j++ {
			if i-j >= 0 {
				filteredSignal[i] += b[j] * signal[i-j]
			}
		}
		for j := 1; j <= order; j++ {
			if i-j >= 0 {
				filteredSignal[i] -= a[j] * filteredSignal[i-j]
			}
		}
	}

	return filteredSignal, nil
}

func butterworthLowpass(order int, normalCutoff float64) (b, a []float64, err error) {
	if order <= 0 {
		return nil, nil, fmt.Errorf("Порядок фильтра должен быть положительным числом")
	}
	if normalCutoff <= 0 || normalCutoff >= 1.0 {
		return nil, nil, fmt.Errorf("Частота среза должна быть между 0 и 1")
	}

	// Вычислите коэффициенты фильтра
	b, a = make([]float64, order+1), make([]float64, order+1)
	wc := 2.0 * math.Pi * normalCutoff

	a[0] = 1
	for i := 0; i < order; i++ {
		a[i+1] = a[i] * (-wc / float64(i+1))
	}

	for i := 0; i <= order; i++ {
		b[i] = 0
	}
	b[order] = 1.0

	return b, a, nil
}
func Butter2(RGB []RGB_float64) []RGB_float64 {
	// dt := 1. / 30 // 64 hz sample rate
	RGBsOut := make([]RGB_float64, len(RGB))
	r := make([]float64, len(RGB))
	g := make([]float64, len(RGB))
	b := make([]float64, len(RGB))
	for i := range RGB {
		r[i] = RGB[i].R
		g[i] = RGB[i].G
		b[i] = RGB[i].B
	}
	sampleRate := 1.0 / 30.0 // Частота дискретизации в Гц
	cutoffFreq := 3.0        // Частота среза в Гц
	r, _ = ButterworthLowpassFilter(r, sampleRate, cutoffFreq)
	g, _ = ButterworthLowpassFilter(g, sampleRate, cutoffFreq)
	b, _ = ButterworthLowpassFilter(b, sampleRate, cutoffFreq)
	for i := range RGB {
		RGBsOut[i].R = r[i]
		RGBsOut[i].G = g[i]
		RGBsOut[i].B = b[i]
	}

	return RGBsOut
}
