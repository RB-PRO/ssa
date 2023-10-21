package tg

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"log"
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

func ExtractRGB(Path, FileName string) (RGBs []RGB_float64, Err error) {

	MakeDir(Path + "in/")
	c := exec.Command(
		"ffmpeg", "-i", Path+FileName, Path+"in/%4d.png",
	) // "P1LC1-edited2.mp4" "P1LC1-edited.avi" "face.mp4"
	c.Stderr = os.Stderr
	c.Run()
	entries, err := os.ReadDir(Path + "in/")
	if err != nil {
		log.Fatal(err)
	}
	p := NewPigs()
	Bar := pb.StartNew(len(entries))
	fileRGB, ErrOpenFile := os.OpenFile(Path+"RGB.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if ErrOpenFile != nil {
		log.Println(ErrOpenFile)
	}
	defer fileRGB.Close()
	// rez := ""
	RGBs = make([]RGB_float64, len(entries))
	for ientries := range entries {
		FilePathIn := Path + "in/" + entries[ientries].Name()
		Dets := p.getCoords(FilePathIn)
		sort.Slice(Dets, func(i, j int) bool { // Сортировка по вероятности
			return Dets[i].Q > Dets[j].Q
		})
		var R, G, B uint32
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
					hsv := RGB2HSV(RGB{R: uint8(r), G: uint8(g), B: uint8(b)})
					// var hsvS float64 = math.Acos((float64(r-g) + float64(r-b)) / (2 * math.Sqrt(math.Pow(float64(r-g), 2)*float64(r-b)*float64(g-b))))
					if cb >= 98 && cb <= 142 &&
						cr >= 135 && cr <= 177 &&
						hsv.H >= 0.01 && hsv.H <= 0.1 {
						// rgbImage2.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
						R += (r >> 8)
						G += (g >> 8)
						B += (b >> 8)
						// fmt.Println(r>>8, g>>8, b>>8)
					}
				}
			}
			var sizes uint32 = uint32((yEnd - yStart) * (xEnd - xStart))
			// fmt.Printf("\n%d;%d;%d   ---   %d    ---   %15.10f;%15.10f;%15.10f\n", R, G, B, sizes, float64(R)/float64(sizes), float64(G)/float64(sizes), float64(B)/float64(sizes))
			// fmt.Println(R, G, B, sizes)
			// rez += fmt.Sprintf("%.8f;%.8f;%.8f\n", float64(R)/float64(sizes), float64(G)/float64(sizes), float64(B)/float64(sizes))
			RGBs[ientries] = RGB_float64{R: float64(R) / float64(sizes), G: float64(G) / float64(sizes), B: float64(B) / float64(sizes)}
			if _, err := fileRGB.WriteString(fmt.Sprintf("%.8f;%.8f;%.8f\n", float64(R)/float64(sizes), float64(G)/float64(sizes), float64(B)/float64(sizes))); err != nil {
				log.Println(err)
			}
		}
		Bar.Increment()
		//break
	}
	Bar.Finish()

	return RGBs, nil
}

/////////////////////////////

type Pigs struct {
	Classifier *pigo.Pigo
}

func NewPigs() *Pigs {
	// consumers.StartForwardStreamConsumer()
	// camtron.StartCam()
	cascadeFile, err := os.ReadFile("cascade/facefinder")
	if err != nil {
		log.Fatalf("Error reading the cascade file: %v", err)
	}

	pigo := pigo.NewPigo()
	// Unpack the binary file. This will return the number of cascade trees,
	// the tree depth, the threshold and the prediction from tree's leaf nodes.
	classifier, err := pigo.Unpack(cascadeFile)
	if err != nil {
		log.Fatalf("Error reading the cascade file: %s", err)
	}

	return &Pigs{
		Classifier: classifier,
	}
}

func (p Pigs) getCoords(filepath string) []pigo.Detection {

	src, err := pigo.GetImage(filepath)
	if err != nil {
		log.Fatalf("Cannot open the image file: %v", err)
	}

	pixels := pigo.RgbToGrayscale(src)
	cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     1000,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,

		ImageParams: pigo.ImageParams{
			Pixels: pixels,
			Rows:   rows,
			Cols:   cols,
			Dim:    cols,
		},
	}

	angle := 0.0 // cascade rotation angle. 0.0 is 0 radians and 1.0 is 2*pi radians

	// Run the classifier over the obtained leaf nodes and return the detection results.
	// The result contains quadruplets representing the row, column, scale and detection score.
	dets := p.Classifier.RunCascade(cParams, angle)
	// fmt.Printf("%+v\n", dets)

	// Calculate the intersection over union (IoU) of two clusters.
	dets = p.Classifier.ClusterDetections(dets, 0.2)

	// fmt.Printf("%+v\n", dets)
	// fmt.Println()

	return dets
}

///////////////////////////

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

//////////////////////////////

type RGB struct {
	R uint8 // Red
	G uint8 // Green
	B uint8 // Blue
}
type RGB_float64 struct {
	R float64 // Red
	G float64 // Green
	B float64 // Blue
}
type HSV struct {
	H float64 // Hue
	S float64 // Saturation
	V float64 // Lightness
}

// RGB2HSV converts RGB color to HSV (HSB)
func RGB2HSV(c RGB) HSV {
	R, G, B := float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0
	max := math.Max(math.Max(R, G), B)
	min := math.Min(math.Min(R, G), B)
	h, s, v := 0.0, 0.0, max
	if max != min {
		d := max - min
		s = d / max
		h = calcHUE(max, R, G, B, d)
	}
	return HSV{h, s, v}
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
