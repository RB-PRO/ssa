package goroi

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"os"
	"os/exec"
	"sort"

	"github.com/cheggaaa/pb"
	pigo "github.com/esimov/pigo/core"
)

type Frame struct {
	FileName string
	Dets     []pigo.Detection
}

func ExtractRGB(Path, FileName string) (R, G, B []float64, Err error) {

	MakeDir(Path + "in/")
	fmt.Println((Path + "in/"))
	fmt.Println((Path + FileName))
	c := exec.Command(
		// ffmpeg -an -y -threads 0 -i in.mp4 -vf "select=not(mod(n\,1))" -vsync vfr in/%4d.png
		"ffmpeg", "-an", "-y", "-threads", "0", "-i", Path+FileName, "-vf", `select=not(mod(n\,1))`, "-vsync", "vfr", Path+"in/%4d.png",
		//"ffmpeg", "-i", Path+FileName, "-r", "15", Path+"in/%4d.png",
	) // "P1LC1-edited2.mp4" "P1LC1-edited.avi" "face.mp4"
	c.Stderr = os.Stderr
	c.Run()
	entries, err := os.ReadDir(Path + "in/")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("os.ReadDir: %v", err)
	}
	p, err := NewPigs("cascade/facefinder")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("NewPigs: %v", err)
	}

	// Выделяем память под слайсы R G B
	R = make([]float64, len(entries))
	G = make([]float64, len(entries))
	B = make([]float64, len(entries))

	Bar := pb.StartNew(len(entries))
	Bar.Prefix(Path + "in/")
	for ientries := range entries {
		FilePathIn := Path + "in/" + entries[ientries].Name()

		r, g, b, ErrCoord := p.Coords2(FilePathIn)
		if ErrCoord != nil {
			fmt.Printf("p.Coords: %v for image [%d/%d] for file %v\n", ErrCoord, ientries, len(entries), entries[ientries].Name())
		}

		R[ientries] = r
		G[ientries] = g
		B[ientries] = b

		Bar.Increment()
		//break
	}
	Bar.Finish()

	return R, G, B, nil
}

func (p *Pigs) Coords(FilePathName string) (R float64, G float64, B float64, Err error) {
	Dets := p.getCoords(FilePathName)
	sort.Slice(Dets, func(i, j int) bool { // Сортировка по вероятности
		return Dets[i].Q > Dets[j].Q
	})

	var Ruint32, Guint32, Buint32 uint32
	var sizes uint32
	if len(Dets) > 0 {
		src, Err := getImageFromFilePath(FilePathName)
		if Err != nil {
			return 0.0, 0.0, 0.0, fmt.Errorf("getImageFromFilePath: %v", Err)
		}
		pix := Dets[0].Scale / 2
		yStart, yEnd := Dets[0].Row-pix, Dets[0].Row+pix
		xStart, xEnd := Dets[0].Col-pix, Dets[0].Col+pix

		for y := yStart; y < yEnd; y++ {
			for x := xStart; x < xEnd; x++ {
				rgb := src.At(x, y)
				r, g, b, _ := rgb.RGBA()
				_, cb, cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
				hsvH, _, _ := RGB2HSV(float64(r), float64(g), float64(b))
				if cb >= 98 && cb <= 142 &&
					cr >= 135 && cr <= 177 &&
					hsvH >= 0.01 && hsvH <= 0.1 {
					Ruint32 += (r >> 8)
					Guint32 += (g >> 8)
					Buint32 += (b >> 8)
				}
			}
		}
		sizes = uint32((yEnd - yStart) * (xEnd - xStart))
	}

	R = float64(Ruint32) / float64(sizes)
	G = float64(Guint32) / float64(sizes)
	B = float64(Buint32) / float64(sizes)
	return R, G, B, nil
}

func getImageFromFilePath(filePath string) (draw.Image, error) {

	// read file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// convert as image.Image
	orig, _, err := image.Decode(f)

	// convert as usable image
	b := orig.Bounds()
	img := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(img, img.Bounds(), orig, b.Min, draw.Src)

	return img, err
}

func (p *Pigs) Coords2(FilePathName string) (R float64, G float64, B float64, Err error) {

	Dets := p.getCoords(FilePathName)
	sort.Slice(Dets, func(i, j int) bool { // Сортировка по вероятности
		return Dets[i].Q > Dets[j].Q
	})

	if len(Dets) > 0 {
		f, ErrOpen := os.Open(FilePathName)
		if ErrOpen != nil {
			return 0.0, 0.0, 0.0, fmt.Errorf("os.Open: picture: %v", ErrOpen)
		}
		defer f.Close()

		img, _, ErrDecode := image.Decode(f)
		if ErrDecode != nil {
			return 0.0, 0.0, 0.0, fmt.Errorf("image.Decode: picture: %v", ErrDecode)
		}

		// pix := Dets[0].Scale / 2
		// face_image := img.(interface {
		// 	SubImage(r image.Rectangle) image.Image
		// }).SubImage(image.Rect(Dets[0].Col-pix, Dets[0].Row-pix,
		// 	Dets[0].Col+pix, Dets[0].Row+pix))

		// fmt.Println(img.Bounds().Max.X - img.Bounds().Min.X)
		// fmt.Println(img.Bounds().Max.Y - img.Bounds().Min.Y)

		// fmt.Println(img.Bounds().Min.X, img.Bounds().Max.X)
		// fmt.Println(img.Bounds().Min.Y, img.Bounds().Max.Y)

		pix := Dets[0].Scale / 2
		// yStart, yEnd := Dets[0].Row-pix, Dets[0].Row+pix
		// xStart, xEnd := Dets[0].Col-pix, Dets[0].Col+pix

		var count int
		for y := Dets[0].Row - pix; y < Dets[0].Row+pix; y++ {
			// fmt.Println(y)
			for x := Dets[0].Col - pix; x < Dets[0].Col+pix; x++ {
				rgb := img.At(x, y)
				r, g, b, _ := rgb.RGBA()

				// fmt.Println(uint8(r>>8), uint8(g>>8), uint8(b>>8))
				// _, cb, cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
				// _, cb, cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))

				// _, cb, cr := ycbcr(r, g, b)

				// hsvH, _, _ := RGB2HSV(float64(r), float64(g), float64(b))
				// h, _, _ := RGBAToHSV(r, g, b, math.MaxInt32)

				// fmt.Println(r, g, b, "--", cb, cr, h, "=", cr>>8, cr>>16)

				// h, _, _ := RGBAToHSV(r, g, b, math.MaxUint32)

				_, cb, cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
				h, _, _ := RGB2HSV(float64(r), float64(g), float64(b))

				if cb >= (98) && cb <= (142) &&
					cr >= (135) && cr <= (177) &&
					h >= 0.01 && h <= 0.1 {
					R += float64(r >> 8)
					G += float64(g >> 8)
					B += float64(b >> 8)
					count++
				}
			}
		}

		// for y := yStart; y < yEnd; y++ {
		// 	for x := xStart; x < xEnd; x++ {
		// 		rgb := src.At(x, y)
		// 		r, g, b, _ := rgb.RGBA()
		// 		_, cb, cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
		// 		hsvH, _, _ := RGB2HSV(float64(r), float64(g), float64(b))
		// 		if cb >= 98 && cb <= 142 &&
		// 			cr >= 135 && cr <= 177 &&
		// 			hsvH >= 0.01 && hsvH <= 0.1 {
		// 			Ruint32 += (r >> 8)
		// 			Guint32 += (g >> 8)
		// 			Buint32 += (b >> 8)
		// 		}
		// 	}
		// }
		// output_file, _ := os.Create("tests/output.jpeg")
		// jpeg.Encode(output_file, face_image, nil)

		// sizes := float64((yEnd - yStart) * (xEnd - xStart))
		count64 := float64(count)
		return R / count64, G / count64, B / count64, nil
	}

	return 0.0, 0.0, 0.0, nil
}

func ycbcr(r, g, b uint32) (uint32, uint32, uint32) {
	y := (19595*r + 38470*g + 7471*b + 1<<15) >> 16

	// The bit twiddling below is equivalent to
	//
	// cb := (-11056*r - 21712*g + 32768*b + 257<<15) >> 16
	// if cb < 0 {
	//     cb = 0
	// } else if cb > 0xff {
	//     cb = ^int32(0)
	// }
	//
	// but uses fewer branches and is faster.
	// Note that the uint8 type conversion in the return
	// statement will convert ^int32(0) to 0xff.
	// The code below to compute cr uses a similar pattern.
	//
	// Note that -11056 - 21712 + 32768 equals 0.
	cb := 32768*b - 11056*r - 21712*g + 257<<15
	if cb&0xff000000 == 0 {
		cb >>= 16
	} else {
		cb = ^(cb >> 31)
	}

	// Note that 32768 - 27440 - 5328 equals 0.
	cr := 32768*r - 27440*g - 5328*b + 257<<15
	if cr&0xff000000 == 0 {
		cr >>= 16
	} else {
		cr = ^(cr >> 31)
	}
	return y, cb, cr
}

func RGBAToHSV(rValue, gValue, bValue, aValue uint32) (h, s, v float64) {

	// The RGBA color components are scaled by the Alpha value, as per:
	// https://golang.org/src/image/color/color.go?s=2394:2435#L21
	// Since we need RGB values in the [0-1] range, we need to divide
	// them by A, making sure it's not 0.
	if aValue == 0 {
		return h, s, v
	}

	a := float64(aValue)
	r := float64(rValue) / a
	g := float64(gValue) / a
	b := float64(bValue) / a

	maxValue := math.Max(r, math.Max(g, b))

	// They're all 0s
	if maxValue == 0 {
		return 0, 0, 0
	}

	minValue := math.Min(r, math.Min(g, b))
	delta := maxValue - minValue

	// Greyscale, only V can be != 0
	if delta == 0 {
		return 0, 0, math.Round(maxValue * 100)
	}

	//hue
	switch maxValue {
	case r:
		h = 60 * ((g - b) / delta)
	case g:
		h = 60 * (((b - r) / delta) + 2)
	case b:
		h = 60 * (((r - g) / delta) + 4)
	}

	if h < 0 {
		h += 360
	}

	h = math.Round(h)

	//saturation
	s = math.Round(100 * delta / maxValue)

	//value
	v = math.Round(maxValue * 100)
	return h, s, v

}
