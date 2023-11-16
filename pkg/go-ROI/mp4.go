package goroi

import (
	"image"
	"image/color"
	"image/draw"
	"log"
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
	c := exec.Command(
		// ffmpeg -an -y -threads 0 -i in.mp4 -vf "select=not(mod(n\,1))" -vsync vfr in/%4d.png
		"ffmpeg", "-an", "-y", "-threads", "0", "-i", Path+FileName, "-vf", `select=not(mod(n\,1))`, "-vsync", "vfr", Path+"in/%4d.png",
		//"ffmpeg", "-i", Path+FileName, "-r", "15", Path+"in/%4d.png",
	) // "P1LC1-edited2.mp4" "P1LC1-edited.avi" "face.mp4"
	c.Stderr = os.Stderr
	c.Run()
	entries, err := os.ReadDir(Path + "in/")
	if err != nil {
		log.Fatal(err)
	}
	p := NewPigs()
	Bar := pb.StartNew(len(entries))

	// Выделяем память под слайсы R G B
	R = make([]float64, len(entries))
	G = make([]float64, len(entries))
	B = make([]float64, len(entries))
	for ientries := range entries {
		FilePathIn := Path + "in/" + entries[ientries].Name()
		Dets := p.getCoords(FilePathIn)
		sort.Slice(Dets, func(i, j int) bool { // Сортировка по вероятности
			return Dets[i].Q > Dets[j].Q
		})

		var Ruint32, Guint32, Buint32 uint32
		var sizes uint32
		if len(Dets) > 0 {
			src, err := getImageFromFilePath(FilePathIn)
			if err != nil {
				panic(err.Error())
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
		// RGBs[ientries] = RGB_float64{R: , G: float64(G) / float64(sizes), B: float64(B) / float64(sizes)}
		R[ientries] = float64(Ruint32) / float64(sizes)
		G[ientries] = float64(Guint32) / float64(sizes)
		B[ientries] = float64(Buint32) / float64(sizes)

		// fmt.Println(FilePathIn, RGBs[ientries])

		Bar.Increment()
		//break
	}
	Bar.Finish()

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
