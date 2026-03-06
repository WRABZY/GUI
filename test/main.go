package main

import (
	"log"
	"os"

	"github.com/wrabzy/gui"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	var f = gui.NewForm("CALCULATOR")
	//icon16, _, err := ebitenutil.NewImageFromFile("test/test_assets/icons/icon16.png")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//icon32, _, err := ebitenutil.NewImageFromFile("test/test_assets/icons/icon32.png")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//icon48, _, err := ebitenutil.NewImageFromFile("test/test_assets/icons/icon48.png")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//f.SetIcons(icon16, icon32, icon48)

	rootLayout := gui.NewVerticalLayout()
	f.SetLayout(rootLayout)

	systemLayout := gui.NewHorizontalLayout()
	displayLayout := gui.NewHorizontalLayout()
	otherLayout := gui.NewHorizontalLayout()

	rootLayout.AddLayout(systemLayout)
	rootLayout.AddLayout(displayLayout)
	rootLayout.AddLayout(otherLayout)

	ivImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/48x48.png")
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
	systemLayout.AddView(iv)

	displayEmptyDigitImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/empty_display_digit.png")
	if err != nil {
		log.Fatal(err)
	}
	var displayDigitView0 = gui.NewImageView()
	var displayDigitView1 = gui.NewImageView()
	var displayDigitView2 = gui.NewImageView()
	var displayDigitView3 = gui.NewImageView()
	var displayDigitView4 = gui.NewImageView()
	var displayDigitView5 = gui.NewImageView()
	var displayDigitView6 = gui.NewImageView()
	var displayDigitView7 = gui.NewImageView()
	var displayDigitView8 = gui.NewImageView()
	var displayDigitView9 = gui.NewImageView()
	displayDigitView0.SetImage(displayEmptyDigitImage)
	displayDigitView1.SetImage(displayEmptyDigitImage)
	displayDigitView2.SetImage(displayEmptyDigitImage)
	displayDigitView3.SetImage(displayEmptyDigitImage)
	displayDigitView4.SetImage(displayEmptyDigitImage)
	displayDigitView5.SetImage(displayEmptyDigitImage)
	displayDigitView6.SetImage(displayEmptyDigitImage)
	displayDigitView7.SetImage(displayEmptyDigitImage)
	displayDigitView8.SetImage(displayEmptyDigitImage)
	displayDigitView9.SetImage(displayEmptyDigitImage)

	displayEmptyDotImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/empty_display_dot.png")
	if err != nil {
		log.Fatal(err)
	}
	var displayDot0 = gui.NewImageView()
	var displayDot1 = gui.NewImageView()
	var displayDot2 = gui.NewImageView()
	var displayDot3 = gui.NewImageView()
	var displayDot4 = gui.NewImageView()
	var displayDot5 = gui.NewImageView()
	var displayDot6 = gui.NewImageView()
	var displayDot7 = gui.NewImageView()
	var displayDot8 = gui.NewImageView()
	displayDot0.SetImage(displayEmptyDotImage)
	displayDot1.SetImage(displayEmptyDotImage)
	displayDot2.SetImage(displayEmptyDotImage)
	displayDot3.SetImage(displayEmptyDotImage)
	displayDot4.SetImage(displayEmptyDotImage)
	displayDot5.SetImage(displayEmptyDotImage)
	displayDot6.SetImage(displayEmptyDotImage)
	displayDot7.SetImage(displayEmptyDotImage)
	displayDot8.SetImage(displayEmptyDotImage)

	displayLayout.AddView(displayDigitView0)
	displayLayout.AddView(displayDot0)
	displayLayout.AddView(displayDigitView1)
	displayLayout.AddView(displayDot1)
	displayLayout.AddView(displayDigitView2)
	displayLayout.AddView(displayDot2)
	displayLayout.AddView(displayDigitView3)
	displayLayout.AddView(displayDot3)
	displayLayout.AddView(displayDigitView4)
	displayLayout.AddView(displayDot4)
	displayLayout.AddView(displayDigitView5)
	displayLayout.AddView(displayDot5)
	displayLayout.AddView(displayDigitView6)
	displayLayout.AddView(displayDot6)
	displayLayout.AddView(displayDigitView7)
	displayLayout.AddView(displayDot7)
	displayLayout.AddView(displayDigitView8)
	displayLayout.AddView(displayDot8)
	displayLayout.AddView(displayDigitView9)

	iv2Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/200x200.png")
	if err != nil {
		log.Fatal(err)
	}
	var iv2 = gui.NewImageView()
	iv2.SetImage(iv2Image)
	iv2.SetOnClickListener(
		func() {
			os.Exit(0)
		},
	)
	otherLayout.AddView(iv2)

	avImg, _, err := ebitenutil.NewImageFromFile("test/test_assets/a100x200.png")
	if err != nil {
		log.Fatal(err)
	}
	var avSpriteSheetId = gui.AddAnimationSpriteSheet(avImg, 0, 0, 100, 200, 30)
	var avAnimation = gui.NewAnimation(avSpriteSheetId, 30).WithTimes(1000 / 60)
	var av = gui.NewAnimationView()
	av.SetAnimation(avAnimation)
	av.SetOnClickListener(
		func() {
			if avAnimation.IsRunning() {
				avAnimation.Pause()
			} else {
				avAnimation.Resume()
			}
		},
	)
	otherLayout.AddView(av)

	var tv = gui.NewTextView("github.com/WRABZY/gui")
	otherLayout.AddView(tv)

	f.OpenWithExecute(avAnimation.Start)
}
