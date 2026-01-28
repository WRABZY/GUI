package gui

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type form struct {
	Name            string
	Fullscreen      bool
	Width, Heigth   int
	elements        map[int]view
	clickableBounds map[ClickableBound]int // TODO: 2'slice with sort, fast search

	mouseState *MouseState

	drawImageOptions *ebiten.DrawImageOptions
}

func NewForm(name string) *form {
	return &form{
		Name:             name,
		Fullscreen:       false,
		Width:            100,
		Heigth:           100,
		elements:         make(map[int]view),
		clickableBounds:  make(map[ClickableBound]int),
		mouseState:       &MouseState{},
		drawImageOptions: &ebiten.DrawImageOptions{},
	}
}

func (f *form) OpenWithExecute(fun func()) {
	go fun()

	f.Open()
}

func (f *form) Open() {
	ebiten.SetWindowTitle(f.Name)
	if f.Fullscreen {
		ebiten.SetFullscreen(true)
	} else {
		ebiten.SetWindowSize(f.Width, f.Heigth)
	}

	if err := ebiten.RunGame(f); err != nil {
		log.Fatal(err)
	}
}

func (f *form) AddView(v view) {
	f.elements[v.id()] = v
	x0, y0 := v.coordinates()
	f.clickableBounds[ClickableBound{
		X0: int(x0),
		Y0: int(y0),
		X1: int(x0) + v.image().Bounds().Dx(),
		Y1: int(y0) + v.image().Bounds().Dy(),
	}] = v.id()
}

func (f *form) SwitchFullscreen() {
	f.Fullscreen = !f.Fullscreen
	ebiten.SetFullscreen(f.Fullscreen)
}

func (f *form) Draw(screen *ebiten.Image) {
	for _, elem := range f.elements {
		x, y := elem.coordinates()
		f.drawImageOptions.GeoM.Reset()
		f.drawImageOptions.GeoM.Translate(x, y)
		screen.DrawImage(elem.image(), f.drawImageOptions)
	}
}

func (f *form) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if f.Fullscreen {
		return ebiten.Monitor().Size()
	}
	return f.Width, f.Heigth
}

func (f *form) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !f.mouseState.Pressed() {
		f.mouseState.Press()
		x, y := ebiten.CursorPosition()
		for bound, id := range f.clickableBounds {
			if x >= bound.X0 && x < bound.X1 && y >= bound.Y0 && y < bound.Y1 {
				clickable, ok := f.elements[id].(Clickable)
				if ok {
					go clickable.Click()
				}
			}
		}
	} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && f.mouseState.Pressed() {
		f.mouseState.Release()
	}
	return nil
}
