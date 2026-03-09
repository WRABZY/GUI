package gui

import "github.com/hajimehoshi/ebiten/v2"

type ImageView struct {
	identifier int
	img        *ebiten.Image
	x, y       float64
	onClick    func()
	onHover    func()
	onHoverOff func()
}

func NewImageView() *ImageView {
	return &ImageView{identifier: assignID()}
}

func (iv *ImageView) SetImage(image *ebiten.Image) *ImageView {
	iv.img = image
	return iv
}

func (iv *ImageView) SetOnClickListener(onClick func()) {
	iv.onClick = onClick
}

func (iv *ImageView) Click() {
	if iv.onClick != nil {
		iv.onClick()
	}
}

func (iv *ImageView) SetOnHoverListener(onHover func()) {
	iv.onHover = onHover
}

func (iv *ImageView) SetOnHoverOffListener(onHoverOff func()) {
	iv.onHoverOff = onHoverOff
}

func (iv *ImageView) Hover() {
	if iv.onHover != nil {
		iv.onHover()
	}
}

func (iv *ImageView) HoverOff() {
	if iv.onHoverOff != nil {
		iv.onHoverOff()
	}
}

func (iv *ImageView) id() int {
	return iv.identifier
}

func (iv *ImageView) image() *ebiten.Image {
	return iv.img
}

func (iv *ImageView) coordinates() (float64, float64) {
	return iv.x, iv.y
}

func (iv *ImageView) setPosition(x, y int) {
	iv.x = float64(x)
	iv.y = float64(y)
}
