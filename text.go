package gui

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"

	ebitenText "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextView struct {
	identifier int
	text       string
	textColor  color.Color
	textSize   float64
	fontFace   font.Face
	x, y       float64
	img        *ebiten.Image
	onClick    func()
}

func NewTextView(text string) *TextView {
	defaultBackgroundColor := color.White
	defaultTextColor := color.Black
	defaultTextSize := 14.0
	defaultFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	defaultFontFace := truetype.NewFace(defaultFont, &truetype.Options{Size: defaultTextSize})
	defaultWidth := font.MeasureString(defaultFontFace, text)

	defaultImage := ebiten.NewImage(defaultWidth.Round(), defaultFontFace.Metrics().Height.Round())
	defaultImage.Fill(defaultBackgroundColor)

	faceToDraw := ebitenText.NewGoXFace(defaultFontFace)
	textOptions := &ebitenText.DrawOptions{}
	textOptions.ColorScale.ScaleWithColor(defaultTextColor)
	ebitenText.Draw(defaultImage, text, faceToDraw, textOptions)
	return &TextView{
		identifier: assignID(),
		text:       text,
		textColor:  defaultTextColor,
		textSize:   defaultTextSize,
		fontFace:   defaultFontFace,
		img:        defaultImage,
	}
}

func (tv *TextView) SetText(text string) {
	tv.text = text
}

func (tv *TextView) SetBackground(image *ebiten.Image) {
	// TODO
}

func (tv *TextView) SetOnClickListener(onClick func()) {
	tv.onClick = onClick
}

func (tv *TextView) Click() {
	if tv.onClick != nil {
		tv.onClick()
	}
}

func (tv *TextView) id() int {
	return tv.identifier
}

func (tv *TextView) image() *ebiten.Image {
	return tv.img
}

func (tv *TextView) coordinates() (float64, float64) {
	return tv.x, tv.y
}

func (tv *TextView) setPosition(x, y int) {
	tv.x = float64(x)
	tv.y = float64(y)
}
