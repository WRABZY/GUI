package main

import (
	"log"
	"os"

	"github.com/wrabzy/gui"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	var f = gui.NewForm("WRABZY/GUI demo")
	f.Width = 400
	f.Heigth = 300

	ivImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/100x200.png")
	if err != nil {
		log.Fatal(err)
	}
	var iv = gui.NewImageView()
	iv.SetImage(ivImage)
	iv.SetOnClickListener(
		func() {
			f.SwitchFullscreen()
		},
	)
	f.AddView(iv)

	iv2Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/200x200.png")
	if err != nil {
		log.Fatal(err)
	}
	var iv2 = gui.NewImageView()
	iv2.SetImage(iv2Image)
	iv2.SetPosition(100, 0)
	iv2.SetOnClickListener(
		func() {
			os.Exit(0)
		},
	)
	f.AddView(iv2)

	avImg, _, err := ebitenutil.NewImageFromFile("test/test_assets/a100x200.png")
	if err != nil {
		log.Fatal(err)
	}
	var avSpriteSheetId = gui.AddAnimationSpriteSheet(avImg, 0, 0, 100, 200, 30)
	var avAnimatiion = gui.NewAnimation(avSpriteSheetId, 30).WithTimes(1000 / 60)
	var av = gui.NewAnimationView()
	av.SetAnimation(avAnimatiion)
	av.SetPosition(300, 0)
	av.SetOnClickListener(
		func() {
			if avAnimatiion.IsRunning() {
				avAnimatiion.Pause()
			} else {
				avAnimatiion.Resume()
			}
		},
	)
	f.AddView(av)

	var tv = gui.NewTextView("github.com/WRABZY/gui")
	tv.SetPosition(0, 200)
	f.AddView(tv)

	f.OpenWithExecute(avAnimatiion.Start)
}
