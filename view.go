package gui

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var nextViewID = math.MinInt

func assignID() int {
	id := nextViewID
	nextViewID++
	return id
}

type view interface {
	id() int
	image() *ebiten.Image
	coordinates() (float64, float64)
}
