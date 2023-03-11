package blackmanharris_test

import (
	"fmt"
)

func Blackmanharris(dv []float64) ([]float64, error) {

	var Xx = [...]complex128{1 + 8i, 2 + 8i, 3 + 8i, 4 + 8i}

	var Xx_re float64
	lenSignal := len(dv)
	res := 513
	fmt.Println("res", res)

	xw := make([]float64, len(dv))
	for i := 0; i < lenSignal; i++ {
		xw[i] = dv[i] * dv[i]
	}
	var y float64
	for i := 0; i < lenSignal; i++ {
		y += dv[i] * dv[i]
	}

	for i := 0; i < lenSignal; i++ {
		Xx_re = real(Xx[i])*real(Xx[i]) - imag(Xx[i])*-imag(Xx[i])
		if real(Xx[i])*-imag(Xx[i])+imag(Xx[i])*real(Xx[i]) == 0.0 {
			Xx_re /= y
		} else if Xx_re == 0.0 {
			Xx_re = 0.0
		} else {
			Xx_re /= y
		}

		xw[i] = Xx_re
	}

	output := make([]float64, 513)
	output[0] = xw[0] / 6.2831853071795862
	for i := 0; i < len(output)-2; i++ {
		output[i+1] = 2.0 * xw[i+1] / 6.2831853071795862
	}

	output[512] = xw[512] / 6.2831853071795862
	return nil, nil
}
