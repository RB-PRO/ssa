package pchip

import (
	"math"
)

func exteriorSlope(d1, d2, h1, h2 float64) float64 {
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

type PchipCoefs struct {
	a, b, c, d []float64
}

func Pchip(x, y, new_x []float64, x_len, new_x_len int) ([]float64, []float64, PchipCoefs) {
	new_y := make([]float64, new_x_len)
	var low_ip1 int
	var hs float64
	del := make([]float64, x_len-1)
	slopes := make([]float64, x_len)
	h := make([]float64, x_len-1)
	var hs3 float64
	var w1 float64
	//var ix int
	pp_coefs := make([]float64, (x_len-1)+(3*(x_len-1)))

	var low_i int
	var high_i int
	var mid_i int
	for low_ip1 := 0; low_ip1 < x_len-1; low_ip1++ {
		hs = x[low_ip1+1] - x[low_ip1]
		del[low_ip1] = (y[low_ip1+1] - y[low_ip1]) / hs
		h[low_ip1] = hs
	}

	for low_ip1 := 0; low_ip1 < x_len-2; low_ip1++ {
		hs = h[low_ip1] + h[low_ip1+1]
		hs3 = 3.0 * hs
		w1 = (h[low_ip1] + hs) / hs3
		hs = (h[low_ip1+1] + hs) / hs3
		hs3 = 0.0
		if del[low_ip1] < 0.0 {
			if del[low_ip1+1] <= del[low_ip1] {
				hs3 = del[low_ip1] / (w1*(del[low_ip1]/del[low_ip1+1]) + hs)
			} else {
				if del[low_ip1+1] < 0.0 {
					hs3 = del[low_ip1+1] / (w1 + hs*(del[low_ip1+1]/del[low_ip1]))
				}
			}
		} else {
			if del[low_ip1] > 0.0 {
				if del[low_ip1+1] >= del[low_ip1] {
					hs3 = del[low_ip1] / (w1*(del[low_ip1]/del[low_ip1+1]) + hs)
				} else {
					if del[low_ip1+1] > 0.0 {
						hs3 = del[low_ip1+1] / (w1 + hs*(del[low_ip1+1]/del[low_ip1]))
					}
				}
			}
		}

		slopes[low_ip1+1] = hs3
	}

	slopes[0] = exteriorSlope(del[0], del[1], h[0], h[1])
	slopes[x_len-1] = exteriorSlope(del[x_len-2], del[x_len-3], h[x_len-2], h[x_len-3])
	for low_ip1 := 0; low_ip1 < x_len-1; low_ip1++ {
		hs = (del[low_ip1] - slopes[low_ip1]) / h[low_ip1]
		hs3 = (slopes[low_ip1+1] - del[low_ip1]) / h[low_ip1]
		pp_coefs[low_ip1] = (hs3 - hs) / h[low_ip1]
		pp_coefs[low_ip1+x_len-1] = 2.0*hs - hs3
		pp_coefs[low_ip1+(2*(x_len-1))] = slopes[low_ip1]
		pp_coefs[low_ip1+(3*(x_len-1))] = y[low_ip1]
	}

	for ix := 0; ix < new_x_len; ix++ {
		low_i = 0
		low_ip1 = 2
		high_i = x_len
		for high_i > low_ip1 {
			mid_i = ((low_i + high_i) + 1) >> 1
			if new_x[ix] >= x[mid_i-1] {
				low_i = mid_i - 1
				low_ip1 = mid_i + 1
			} else {
				high_i = mid_i
			}
		}

		hs = new_x[ix] - x[low_i]
		hs3 = pp_coefs[low_i]
		for low_ip1 := 0; low_ip1 < 3; low_ip1++ {
			hs3 = hs*hs3 + pp_coefs[low_i+(low_ip1+1)*(x_len-1)]
		}

		new_y[ix] = hs3
		/*
			// my
			if new_y[ix] < 0.15 {
				new_y[ix] = new_y[ix-1]
				fmt.Println("YESS")
			}
		*/
	}

	var coefs PchipCoefs
	lenCoefs := len(pp_coefs) / 4
	coefs.a = make([]float64, lenCoefs)
	coefs.b = make([]float64, lenCoefs)
	coefs.c = make([]float64, lenCoefs)
	coefs.d = make([]float64, lenCoefs)
	for i := 0; i < lenCoefs; i++ {
		coefs.a[i] = pp_coefs[i]
		coefs.b[i] = pp_coefs[1*(len(pp_coefs)/4)+i]
		coefs.c[i] = pp_coefs[2*(len(pp_coefs)/4)+i]
		coefs.d[i] = pp_coefs[3*(len(pp_coefs)/4)+i]
	}

	return new_y, pp_coefs, coefs
}

// Функция, реализующая сглаживание с помощью фильтра Савицкого-Голея
