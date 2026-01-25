package view

import "github.com/hajimehoshi/ebiten/v2"

type imageView struct {
	identifier int
	img        *ebiten.Image
	x, y       float64
	onClick    func()
}

func NewImageView() *imageView {
	return &imageView{identifier: assignID()}
}

func (iv *imageView) SetPosition(x, y int) {
	iv.x = float64(x)
	iv.y = float64(y)
}

func (iv *imageView) SetImage(image *ebiten.Image) {
	iv.img = image
}

func (iv *imageView) SetOnClickListener(onClick func()) {
	iv.onClick = onClick
}

func (iv *imageView) Click() {
	if iv.onClick != nil {
		iv.onClick()
	}
}

func (iv *imageView) id() int {
	return iv.identifier
}

func (iv *imageView) image() *ebiten.Image {
	return iv.img
}

func (iv *imageView) coordinates() (float64, float64) {
	return iv.x, iv.y
}
