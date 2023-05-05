package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type GUI struct {
	a   fyne.App
	win fyne.Window
}

func New() *GUI {
	a := app.New()
	win := a.NewWindow("Hello World")
	win.SetContent(widget.NewVBox(

		widget.NewLabel("Hello World!"),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))
	win.SetFixedSize(true)
	win.Resize(fyne.NewSize(480, 320))

	image := canvas.NewImageFromFile("png/LBD.png")
	win.SetContent(container.NewGridWithColumns(1, image))

	win.ShowAndRun()
	return &GUI{a: a, win: win}
}

func (g *GUI) VievImage() {
	image := canvas.NewImageFromFile("png/LBD.png")
	g.win.SetContent(container.NewGridWithColumns(1, image))
	// g.win.
}
