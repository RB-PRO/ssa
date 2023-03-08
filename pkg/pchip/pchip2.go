package pchip

import "math"

func Pchip2(x, y, new_x []float64, x_len, new_x_len int) ([]float64, []float64, PchipCoefs) {

	del := make([]float64, 101)

	slopes := make([]float64, 102)
	var d float64
	for k := 0; k < 101; k++ {
		del[k] = x[k+1] - x[k]
	}

	for k := 0; k < 100; k++ {
		slopes[k+1] = 0.0
		if del[k] < 0.0 {
			d = del[k+1]
			if d <= del[k] {
				slopes[k+1] = del[k] / (0.5*(del[k]/d) + 0.5)
			} else {
				if d < 0.0 {
					slopes[k+1] = d / (0.5 + 0.5*(d/del[k]))
				}
			}
		} else {
			if del[k] > 0.0 {
				d = del[k+1]
				if d >= del[k] {
					slopes[k+1] = del[k] / (0.5*(del[k]/del[k+1]) + 0.5)
				} else {
					if d > 0.0 {
						slopes[k+1] = del[k+1] / (0.5 + 0.5*(del[k+1]/del[k]))
					}
				}
			}
		}
	}

	slopes[0] = exteriorSlope2(del[0], del[1], 1.0, 1.0)
	slopes[101] = exteriorSlope2(del[100], del[99], 1.0, 1.0)
	return nil, nil, PchipCoefs{C: slopes}
}

/* Function Definitions */
func exteriorSlope2(d1, d2, h1, h2 float64) float64 {
	var s float64
	var signd1 float64
	var signs float64
	s = ((2.0*h1+h2)*d1 - h1*d2) / (h1 + h2)
	signd1 = d1
	if d1 < 0.0 {
		signd1 = -1.0
	} else if d1 > 0.0 {
		signd1 = 1.0
	} else {
		if d1 == 0.0 {
			signd1 = 0.0
		}
	}

	signs = s
	if s < 0.0 {
		signs = -1.0
	} else if s > 0.0 {
		signs = 1.0
	} else {
		if s == 0.0 {
			signs = 0.0
		}
	}

	if signs != signd1 {
		s = 0.0
	} else {
		signs = d2
		if d2 < 0.0 {
			signs = -1.0
		} else if d2 > 0.0 {
			signs = 1.0
		} else {
			if d2 == 0.0 {
				signs = 0.0
			}
		}

		if (signd1 != signs) && (math.Abs(s) > math.Abs(3.0*d1)) {
			s = 3.0 * d1
		}
	}

	return s
}
