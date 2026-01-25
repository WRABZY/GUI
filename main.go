package main

import (
	"GUI/view"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	var f = view.NewForm("WRABZY/GUI demo")
	f.Width = 300
	f.Heigth = 200

	ivImage, _, err := ebitenutil.NewImageFromFile("test_assets/100x200.png")
	if err != nil {
		log.Fatal(err)
	}
	var iv = view.NewImageView()
	iv.SetImage(ivImage)
	iv.SetOnClickListener(
		func() {
			f.SwitchFullscreen()
		},
	)
	f.AddView(iv)

	iv2Image, _, err := ebitenutil.NewImageFromFile("test_assets/200x200.png")
	if err != nil {
		log.Fatal(err)
	}
	var iv2 = view.NewImageView()
	iv2.SetImage(iv2Image)
	iv2.SetPosition(100, 0)
	iv2.SetOnClickListener(
		func() {
			os.Exit(0)
		},
	)
	f.AddView(iv2)

	f.Open()
}
