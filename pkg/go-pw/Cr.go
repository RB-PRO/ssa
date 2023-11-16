package gopw

import (
	"fmt"
	"math"

	"github.com/RB-PRO/ssa/pkg/movmean"
)

// Метод cr извлечения сигнала фотоплетизмографии
func cr(R, G, B []float64) ([]float64, error) {
	if len(R) != len(G) || len(G) != len(B) {
		return nil, fmt.Errorf("the signal lengths R,G,B are not equal")
	}
	pw := make([]float64, len(R))
	pw2 := make([]float64, len(R))

	for i := range R {
		pw[i] = (R[i]*112.0 - R[i]*93.8 - R[i]*18.2) / 255.0
	}

	// Вычитаем тренд
	pw_smooth, _ := movmean.Movmean(pw, 32)
	for i := range pw {
		pw[i] -= pw_smooth[i]
	}

	// Квадрат
	for i := range pw {
		pw2[i] = math.Pow(pw[i], 2)
	}

	// fMi := 40.0 / 60.0
	// cad := 30
	// SMO_med := cad / int(fMi)

	DEV_med, ErrmedianFilter := medianFilter(pw2, 30*60/40)
	if ErrmedianFilter != nil {
		return nil, fmt.Errorf("medianFilter: %v", ErrmedianFilter)
	}

	for i := range pw {
		pw[i] /= math.Sqrt(DEV_med[i])
	}
	prcMi := prctile(pw, 0.1)
	prcMa := prctile(pw, 99.9)
	prcMi = -0.0197
	prcMa = 0.0207
	// fmt.Println("prcMi", prcMi, "prcMa", prcMa)

	for i := range pw {
		if pw[i] < prcMi {
			pw[i] = prcMi
		}
		if pw[i] > prcMa {
			pw[i] = prcMa
		}
	}

	return pw, nil
}
