package gui

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type form struct {
	Name            string
	Fullscreen      bool
	Width, Height   int
	icons           []image.Image
	layout          *layout
	elements        map[int]view
	clickableBounds map[ClickableBound]int // TODO: 2'slice with sort, fast search

	mouseState *MouseState

	drawImageOptions *ebiten.DrawImageOptions
}

func NewForm(name string) *form {
	return &form{
		Name:             name,
		Fullscreen:       false,
		elements:         make(map[int]view),
		clickableBounds:  make(map[ClickableBound]int),
		mouseState:       &MouseState{},
		drawImageOptions: &ebiten.DrawImageOptions{},
	}
}

func (f *form) SetIcons(icons ...image.Image) {
	f.icons = icons
}

func (f *form) OpenWithExecute(fun func()) {
	go fun()

	f.Open()
}

func (f *form) Open() {
	ebiten.SetWindowTitle(f.Name)
	if len(f.icons) > 0 {
		ebiten.SetWindowIcon(f.icons)
	}

	if f.Fullscreen {
		ebiten.SetFullscreen(true)
	} else {
		if f.layout != nil {
			f.Width, f.Height = f.layout.SizeOfContent()
			f.DistributeViewsByLayout()
		}
		ebiten.SetWindowSize(f.Width, f.Height)
	}

	if err := ebiten.RunGame(f); err != nil {
		log.Fatal(err)
	}
}

func (f *form) SetLayout(l *layout) {
	f.layout = l
}

func (f *form) DistributeViewsByLayout() {
	if f.layout == nil {
		return
	}

	for {
		elem, last := f.layout.ThroughLayout()
		if elem != nil {
			f.elements[elem.id()] = elem
			x0, y0 := elem.coordinates()
			f.clickableBounds[ClickableBound{
				X0: int(x0),
				Y0: int(y0),
				X1: int(x0) + elem.image().Bounds().Dx(),
				Y1: int(y0) + elem.image().Bounds().Dy(),
			}] = elem.id()
		}
		if last {
			return
		}
	}

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
	return f.Width, f.Height
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
