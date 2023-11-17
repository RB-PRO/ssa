package goroi

import (
	"fmt"
	"log"
	"os"

	pigo "github.com/esimov/pigo/core"
)

type Pigs struct {
	Classifier *pigo.Pigo
}

func NewPigs(FileCascade string) (*Pigs, error) {
	// consumers.StartForwardStreamConsumer()
	// camtron.StartCam()
	cascadeFile, err := os.ReadFile(FileCascade) // "cascade/facefinder"
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: Error reading the cascade file: %v", err)
	}

	pigo := pigo.NewPigo()
	// Unpack the binary file. This will return the number of cascade trees,
	// the tree depth, the threshold and the prediction from tree's leaf nodes.
	classifier, err := pigo.Unpack(cascadeFile)
	if err != nil {
		return nil, fmt.Errorf("pigo.Unpack: Error reading the cascade file: %v", err)
	}

	return &Pigs{
		Classifier: classifier,
	}, nil
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
	dets = p.Classifier.ClusterDetections(dets, 0.1)

	// fmt.Printf("%+v\n", dets)
	// fmt.Println()

	return dets
}
