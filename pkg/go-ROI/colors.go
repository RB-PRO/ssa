package goroi

import "math"

// RGB2HSV converts RGB color to HSV (HSB)
func RGB2HSV(R, G, B float64) (H, S, V float64) {
	max := math.Max(math.Max(R, G), B)
	min := math.Min(math.Min(R, G), B)
	h, s, v := 0.0, 0.0, max
	if max != min {
		d := max - min
		s = d / max
		h = calcHUE(max, R, G, B, d)
	}
	return h, s, v
}
func calcHUE(max, r, g, b, d float64) float64 {
	var h float64
	switch max {
	case r:
		if g < b {
			h = (g-b)/d + 6.0
		} else {
			h = (g - b) / d
		}
	case g:
		h = (b-r)/d + 2.0
	case b:
		h = (r-g)/d + 4.0
	}
	return h / 6
}
