package GameUI

import (
	"fyne.io/fyne"
	// "fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	// "fyne.io/fyne/widget"
	"image/color"
)

type Canvas struct {
	// App fyne.App
	// Window    fyne.Window
	Container *fyne.Container
	Width     float64
	Height    float64
}

func (c *Canvas) Image(x, y float64, w, h int, name string) {
	// x, y = dimen(x, y, c.Width, c.Height)
	AbsImage(c.Container, int(x), int(y), w, h, name)
}

func NewCanvas(name string, w, h int) Canvas {
	c := Canvas{
		// App: app.New(),
		// Window:    App.NewWindow(name),
		Container: fyne.NewContainer(iRect(w/2, h/2, w, h, color.RGBA{0, 0, 0, 255})),
		Width:     float64(w),
		Height:    float64(h),
	}
	return c
}

func pct(p float64, m float64) float64 {
	return ((p / 100.0) * m)
}

func dimen(xp, yp, w, h float64) (float64, float64) {
	// return pct(xp, w), pct(100-yp, h)
	return pct(xp, w), pct(yp, h)
}

func iRect(x, y, w, h int, color color.RGBA) *canvas.Rectangle {
	r := &canvas.Rectangle{FillColor: color}
	r.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	r.Resize(fyne.Size{Width: w, Height: h})
	return r
}

func AbsImage(cont *fyne.Container, x, y, w, h int, name string) {
	i := canvas.NewImageFromFile(name)
	// i.Move(fyne.Position{X: x - (w / 2), Y: y - (h / 2)})
	i.Move(fyne.Position{X: x, Y: y})
	i.Resize(fyne.Size{Width: w, Height: h})
	cont.AddObject(i)
}
