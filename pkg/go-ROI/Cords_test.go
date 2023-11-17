package goroi

import (
	"fmt"
	"image/color"
	"math"
	"testing"
)

func TestCoords(t *testing.T) {
	// "cascade/facefinder"
	p, ErrPig := NewPigs("../../cascade/facefinder")
	if ErrPig != nil {
		t.Error(ErrPig)
	}
	R, G, B, ErrCoords := p.Coords2("tests/0200.png")
	if ErrCoords != nil {
		t.Error(ErrCoords)
	}
	// 200 - 138.4217 111.4517 92.7798
	// 300 - 138.3724 111.1482 92.4330
	fmt.Println(R, G, B)

	var r uint32 = 26728
	var g uint32 = 23387
	var b uint32 = 17476
	fmt.Println(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	// _, cb, cr := ycbcr(r, g, b)
	_, cb, cr := color.YCbCrToRGB(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	h, _, _ := RGBAToHSV(r, g, b, math.MaxUint32)
	fmt.Println("cb,cr,h", cb, cr, h)

	fmt.Println("math.MaxUint8 ", math.MaxUint8)
	fmt.Println("math.MaxUint16", math.MaxUint16)
	fmt.Println("math.MaxUint32", math.MaxUint32)

}
