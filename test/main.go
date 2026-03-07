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

	topBorderLayout := gui.NewHorizontalLayout()
	displayLayout := gui.NewHorizontalLayout()
	displayBorderLayout := gui.NewHorizontalLayout()
	buttonsRow1Layout := gui.NewHorizontalLayout()
	buttonsRow2Layout := gui.NewHorizontalLayout()
	buttonsRow3Layout := gui.NewHorizontalLayout()
	buttonsRow4Layout := gui.NewHorizontalLayout()
	buttonsRow5Layout := gui.NewHorizontalLayout()
	buttonsRow6Layout := gui.NewHorizontalLayout()
	botBorderLayout := gui.NewHorizontalLayout()

	rootLayout.AddLayout(topBorderLayout)
	rootLayout.AddLayout(displayLayout)
	rootLayout.AddLayout(displayBorderLayout)
	rootLayout.AddLayout(buttonsRow1Layout)
	rootLayout.AddLayout(buttonsRow2Layout)
	rootLayout.AddLayout(buttonsRow3Layout)
	rootLayout.AddLayout(buttonsRow4Layout)
	rootLayout.AddLayout(buttonsRow5Layout)
	rootLayout.AddLayout(buttonsRow6Layout)
	rootLayout.AddLayout(botBorderLayout)

	horizontalBorderImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/border_horizontal.png")
	if err != nil {
		log.Fatal(err)
	}
	var topHBorder = gui.NewImageView()
	topHBorder.SetImage(horizontalBorderImage)
	topBorderLayout.AddView(topHBorder)

	displayVerticalBorderImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/border_vertical_display.png")
	if err != nil {
		log.Fatal(err)
	}
	displayDigitFiller, _, err := ebitenutil.NewImageFromFile("test/test_assets/display_digit_start_end_filler.png")
	if err != nil {
		log.Fatal(err)
	}
	displayEmptyDigitImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/display_digit_empty.png")
	if err != nil {
		log.Fatal(err)
	}
	displayEmptyDotImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/display_dot_empty.png")
	if err != nil {
		log.Fatal(err)
	}
	var displayVerticalBorderLeft = gui.NewImageView()
	var displayDigitFillerStart = gui.NewImageView()
	var displayDigitView0 = gui.NewImageView()
	var displayDot0 = gui.NewImageView()
	var displayDigitView1 = gui.NewImageView()
	var displayDot1 = gui.NewImageView()
	var displayDigitView2 = gui.NewImageView()
	var displayDot2 = gui.NewImageView()
	var displayDigitView3 = gui.NewImageView()
	var displayDot3 = gui.NewImageView()
	var displayDigitView4 = gui.NewImageView()
	var displayDot4 = gui.NewImageView()
	var displayDigitView5 = gui.NewImageView()
	var displayDot5 = gui.NewImageView()
	var displayDigitView6 = gui.NewImageView()
	var displayDot6 = gui.NewImageView()
	var displayDigitView7 = gui.NewImageView()
	var displayDot7 = gui.NewImageView()
	var displayDigitView8 = gui.NewImageView()
	var displayDot8 = gui.NewImageView()
	var displayDigitView9 = gui.NewImageView()
	var displayDigitFillerEnd = gui.NewImageView()
	var displayVerticalBorderRight = gui.NewImageView()
	displayVerticalBorderLeft.SetImage(displayVerticalBorderImage)
	displayDigitFillerStart.SetImage(displayDigitFiller)
	displayDigitView0.SetImage(displayEmptyDigitImage)
	displayDot0.SetImage(displayEmptyDotImage)
	displayDigitView1.SetImage(displayEmptyDigitImage)
	displayDot1.SetImage(displayEmptyDotImage)
	displayDigitView2.SetImage(displayEmptyDigitImage)
	displayDot2.SetImage(displayEmptyDotImage)
	displayDigitView3.SetImage(displayEmptyDigitImage)
	displayDot3.SetImage(displayEmptyDotImage)
	displayDigitView4.SetImage(displayEmptyDigitImage)
	displayDot4.SetImage(displayEmptyDotImage)
	displayDigitView5.SetImage(displayEmptyDigitImage)
	displayDot5.SetImage(displayEmptyDotImage)
	displayDigitView6.SetImage(displayEmptyDigitImage)
	displayDot6.SetImage(displayEmptyDotImage)
	displayDigitView7.SetImage(displayEmptyDigitImage)
	displayDot7.SetImage(displayEmptyDotImage)
	displayDigitView8.SetImage(displayEmptyDigitImage)
	displayDot8.SetImage(displayEmptyDotImage)
	displayDigitView9.SetImage(displayEmptyDigitImage)
	displayDigitFillerEnd.SetImage(displayDigitFiller)
	displayVerticalBorderRight.SetImage(displayVerticalBorderImage)
	displayLayout.AddView(displayVerticalBorderLeft)
	displayLayout.AddView(displayDigitFillerStart)
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
	displayLayout.AddView(displayDigitFillerEnd)
	displayLayout.AddView(displayVerticalBorderRight)

	var displayHBorder = gui.NewImageView()
	displayHBorder.SetImage(horizontalBorderImage)
	displayBorderLayout.AddView(displayHBorder)

	verticalBorderImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/border_vertical.png")
	if err != nil {
		log.Fatal(err)
	}
	l1Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/op_l1.png")
	if err != nil {
		log.Fatal(err)
	}
	upImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/up.png")
	if err != nil {
		log.Fatal(err)
	}
	r1Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/op_r1.png")
	if err != nil {
		log.Fatal(err)
	}
	emptyImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/empty.png")
	if err != nil {
		log.Fatal(err)
	}
	l2Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/l2.png")
	if err != nil {
		log.Fatal(err)
	}
	yImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/y.png")
	if err != nil {
		log.Fatal(err)
	}
	r2Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/r2.png")
	if err != nil {
		log.Fatal(err)
	}
	var buttonsRow1VerticalBorderLeft = gui.NewImageView()
	var buttonsRow1L1 = gui.NewImageView()
	var buttonsRow1Up = gui.NewImageView()
	var buttonsRow1R1 = gui.NewImageView()
	var buttonsRow1Empty = gui.NewImageView()
	var buttonsRow1L2 = gui.NewImageView()
	var buttonsRow1Y = gui.NewImageView()
	var buttonsRow1R2 = gui.NewImageView()
	var buttonsRow1VerticalBorderRight = gui.NewImageView()
	buttonsRow1VerticalBorderLeft.SetImage(verticalBorderImage)
	buttonsRow1L1.SetImage(l1Image)
	buttonsRow1Up.SetImage(upImage)
	buttonsRow1R1.SetImage(r1Image)
	buttonsRow1Empty.SetImage(emptyImage)
	buttonsRow1L2.SetImage(l2Image)
	buttonsRow1Y.SetImage(yImage)
	buttonsRow1R2.SetImage(r2Image)
	buttonsRow1VerticalBorderRight.SetImage(verticalBorderImage)
	buttonsRow1Layout.AddView(buttonsRow1VerticalBorderLeft)
	buttonsRow1Layout.AddView(buttonsRow1L1)
	buttonsRow1Layout.AddView(buttonsRow1Up)
	buttonsRow1Layout.AddView(buttonsRow1R1)
	buttonsRow1Layout.AddView(buttonsRow1Empty)
	buttonsRow1Layout.AddView(buttonsRow1L2)
	buttonsRow1Layout.AddView(buttonsRow1Y)
	buttonsRow1Layout.AddView(buttonsRow1R2)
	buttonsRow1Layout.AddView(buttonsRow1VerticalBorderRight)
	buttonsRow1L2.SetOnClickListener(
		func() {
			f.SwitchFullscreen()
		},
	)
	buttonsRow1R2.SetOnClickListener(
		func() {
			os.Exit(0)
		},
	)

	leftImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/left.png")
	if err != nil {
		log.Fatal(err)
	}
	centerImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/center.png")
	if err != nil {
		log.Fatal(err)
	}
	rightImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/right.png")
	if err != nil {
		log.Fatal(err)
	}
	xImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/x.png")
	if err != nil {
		log.Fatal(err)
	}
	bImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/b.png")
	if err != nil {
		log.Fatal(err)
	}
	var buttonsRow2VerticalBorderLeft = gui.NewImageView()
	var buttonsRow2Left = gui.NewImageView()
	var buttonsRow2Center = gui.NewImageView()
	var buttonsRow2Right = gui.NewImageView()
	var buttonsRow2Empty1 = gui.NewImageView()
	var buttonsRow2X = gui.NewImageView()
	var buttonsRow2Empty2 = gui.NewImageView()
	var buttonsRow2B = gui.NewImageView()
	var buttonsRow2VerticalBorderRight = gui.NewImageView()
	buttonsRow2VerticalBorderLeft.SetImage(verticalBorderImage)
	buttonsRow2Left.SetImage(leftImage)
	buttonsRow2Center.SetImage(centerImage)
	buttonsRow2Right.SetImage(rightImage)
	buttonsRow2Empty1.SetImage(emptyImage)
	buttonsRow2X.SetImage(xImage)
	buttonsRow2Empty2.SetImage(emptyImage)
	buttonsRow2B.SetImage(bImage)
	buttonsRow2VerticalBorderRight.SetImage(verticalBorderImage)
	buttonsRow2Layout.AddView(buttonsRow2VerticalBorderLeft)
	buttonsRow2Layout.AddView(buttonsRow2Left)
	buttonsRow2Layout.AddView(buttonsRow2Center)
	buttonsRow2Layout.AddView(buttonsRow2Right)
	buttonsRow2Layout.AddView(buttonsRow2Empty1)
	buttonsRow2Layout.AddView(buttonsRow2X)
	buttonsRow2Layout.AddView(buttonsRow2Empty2)
	buttonsRow2Layout.AddView(buttonsRow2B)
	buttonsRow2Layout.AddView(buttonsRow2VerticalBorderRight)

	botImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/bot.png")
	if err != nil {
		log.Fatal(err)
	}
	b7Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/7.png")
	if err != nil {
		log.Fatal(err)
	}
	b8Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/8.png")
	if err != nil {
		log.Fatal(err)
	}
	b9Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/9.png")
	if err != nil {
		log.Fatal(err)
	}
	aImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/a.png")
	if err != nil {
		log.Fatal(err)
	}
	var buttonsRow3VerticalBorderLeft = gui.NewImageView()
	var buttonsRow3Empty1 = gui.NewImageView()
	var buttonsRow3Bot = gui.NewImageView()
	var buttonsRow3B7 = gui.NewImageView()
	var buttonsRow3B8 = gui.NewImageView()
	var buttonsRow3B9 = gui.NewImageView()
	var buttonsRow3A = gui.NewImageView()
	var buttonsRow3Empty2 = gui.NewImageView()
	var buttonsRow3VerticalBorderRight = gui.NewImageView()
	buttonsRow3VerticalBorderLeft.SetImage(verticalBorderImage)
	buttonsRow3Empty1.SetImage(emptyImage)
	buttonsRow3Bot.SetImage(botImage)
	buttonsRow3B7.SetImage(b7Image)
	buttonsRow3B8.SetImage(b8Image)
	buttonsRow3B9.SetImage(b9Image)
	buttonsRow3A.SetImage(aImage)
	buttonsRow3Empty2.SetImage(emptyImage)
	buttonsRow3VerticalBorderRight.SetImage(verticalBorderImage)
	buttonsRow3Layout.AddView(buttonsRow3VerticalBorderLeft)
	buttonsRow3Layout.AddView(buttonsRow3Empty1)
	buttonsRow3Layout.AddView(buttonsRow3Bot)
	buttonsRow3Layout.AddView(buttonsRow3B7)
	buttonsRow3Layout.AddView(buttonsRow3B8)
	buttonsRow3Layout.AddView(buttonsRow3B9)
	buttonsRow3Layout.AddView(buttonsRow3A)
	buttonsRow3Layout.AddView(buttonsRow3Empty2)
	buttonsRow3Layout.AddView(buttonsRow3VerticalBorderRight)

	b4Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/4.png")
	if err != nil {
		log.Fatal(err)
	}
	b5Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/5.png")
	if err != nil {
		log.Fatal(err)
	}
	b6Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/6.png")
	if err != nil {
		log.Fatal(err)
	}
	var buttonsRow4VerticalBorderLeft = gui.NewImageView()
	var buttonsRow4Empty1 = gui.NewImageView()
	var buttonsRow4Empty2 = gui.NewImageView()
	var buttonsRow4B4 = gui.NewImageView()
	var buttonsRow4B5 = gui.NewImageView()
	var buttonsRow4B6 = gui.NewImageView()
	var buttonsRow4Empty3 = gui.NewImageView()
	var buttonsRow4Empty4 = gui.NewImageView()
	var buttonsRow4VerticalBorderRight = gui.NewImageView()
	buttonsRow4VerticalBorderLeft.SetImage(verticalBorderImage)
	buttonsRow4Empty1.SetImage(emptyImage)
	buttonsRow4Empty2.SetImage(emptyImage)
	buttonsRow4B4.SetImage(b4Image)
	buttonsRow4B5.SetImage(b5Image)
	buttonsRow4B6.SetImage(b6Image)
	buttonsRow4Empty3.SetImage(emptyImage)
	buttonsRow4Empty4.SetImage(emptyImage)
	buttonsRow4VerticalBorderRight.SetImage(verticalBorderImage)
	buttonsRow4Layout.AddView(buttonsRow4VerticalBorderLeft)
	buttonsRow4Layout.AddView(buttonsRow4Empty1)
	buttonsRow4Layout.AddView(buttonsRow4Empty2)
	buttonsRow4Layout.AddView(buttonsRow4B4)
	buttonsRow4Layout.AddView(buttonsRow4B5)
	buttonsRow4Layout.AddView(buttonsRow4B6)
	buttonsRow4Layout.AddView(buttonsRow4Empty3)
	buttonsRow4Layout.AddView(buttonsRow4Empty4)
	buttonsRow4Layout.AddView(buttonsRow4VerticalBorderRight)

	b1Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/1.png")
	if err != nil {
		log.Fatal(err)
	}
	b2Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/2.png")
	if err != nil {
		log.Fatal(err)
	}
	b3Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/3.png")
	if err != nil {
		log.Fatal(err)
	}
	var buttonsRow5VerticalBorderLeft = gui.NewImageView()
	var buttonsRow5Empty1 = gui.NewImageView()
	var buttonsRow5Empty2 = gui.NewImageView()
	var buttonsRow5B1 = gui.NewImageView()
	var buttonsRow5B2 = gui.NewImageView()
	var buttonsRow5B3 = gui.NewImageView()
	var buttonsRow5Empty3 = gui.NewImageView()
	var buttonsRow5Empty4 = gui.NewImageView()
	var buttonsRow5VerticalBorderRight = gui.NewImageView()
	buttonsRow5VerticalBorderLeft.SetImage(verticalBorderImage)
	buttonsRow5Empty1.SetImage(emptyImage)
	buttonsRow5Empty2.SetImage(emptyImage)
	buttonsRow5B1.SetImage(b1Image)
	buttonsRow5B2.SetImage(b2Image)
	buttonsRow5B3.SetImage(b3Image)
	buttonsRow5Empty3.SetImage(emptyImage)
	buttonsRow5Empty4.SetImage(emptyImage)
	buttonsRow5VerticalBorderRight.SetImage(verticalBorderImage)
	buttonsRow5Layout.AddView(buttonsRow5VerticalBorderLeft)
	buttonsRow5Layout.AddView(buttonsRow5Empty1)
	buttonsRow5Layout.AddView(buttonsRow5Empty2)
	buttonsRow5Layout.AddView(buttonsRow5B1)
	buttonsRow5Layout.AddView(buttonsRow5B2)
	buttonsRow5Layout.AddView(buttonsRow5B3)
	buttonsRow5Layout.AddView(buttonsRow5Empty3)
	buttonsRow5Layout.AddView(buttonsRow5Empty4)
	buttonsRow5Layout.AddView(buttonsRow5VerticalBorderRight)

	unaryMinusImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/unary_minus.png")
	if err != nil {
		log.Fatal(err)
	}
	b0Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/0.png")
	if err != nil {
		log.Fatal(err)
	}
	dotImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/dot.png")
	if err != nil {
		log.Fatal(err)
	}
	var buttonsRow6VerticalBorderLeft = gui.NewImageView()
	var buttonsRow6Empty1 = gui.NewImageView()
	var buttonsRow6Empty2 = gui.NewImageView()
	var buttonsRow6UnaryMinus = gui.NewImageView()
	var buttonsRow6B0 = gui.NewImageView()
	var buttonsRow6Bdot = gui.NewImageView()
	var buttonsRow6Empty3 = gui.NewImageView()
	var buttonsRow6Empty4 = gui.NewImageView()
	var buttonsRow6VerticalBorderRight = gui.NewImageView()
	buttonsRow6VerticalBorderLeft.SetImage(verticalBorderImage)
	buttonsRow6Empty1.SetImage(emptyImage)
	buttonsRow6Empty2.SetImage(emptyImage)
	buttonsRow6UnaryMinus.SetImage(unaryMinusImage)
	buttonsRow6B0.SetImage(b0Image)
	buttonsRow6Bdot.SetImage(dotImage)
	buttonsRow6Empty3.SetImage(emptyImage)
	buttonsRow6Empty4.SetImage(emptyImage)
	buttonsRow6VerticalBorderRight.SetImage(verticalBorderImage)
	buttonsRow6Layout.AddView(buttonsRow6VerticalBorderLeft)
	buttonsRow6Layout.AddView(buttonsRow6Empty1)
	buttonsRow6Layout.AddView(buttonsRow6Empty2)
	buttonsRow6Layout.AddView(buttonsRow6UnaryMinus)
	buttonsRow6Layout.AddView(buttonsRow6B0)
	buttonsRow6Layout.AddView(buttonsRow6Bdot)
	buttonsRow6Layout.AddView(buttonsRow6Empty3)
	buttonsRow6Layout.AddView(buttonsRow6Empty4)
	buttonsRow6Layout.AddView(buttonsRow6VerticalBorderRight)

	var botHBorder = gui.NewImageView()
	botHBorder.SetImage(horizontalBorderImage)
	botBorderLayout.AddView(botHBorder)

	/*
		ivImage, _, err := ebitenutil.NewImageFromFile("test/test_assets/48x48.png")
		if err != nil {
			log.Fatal(err)
		}
		var iv = gui.NewImageView()
		iv.SetImage(ivImage)

		iv2Image, _, err := ebitenutil.NewImageFromFile("test/test_assets/200x200.png")
		if err != nil {
			log.Fatal(err)
		}
		var iv2 = gui.NewImageView()
		iv2.SetImage(iv2Image)

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
	*/

	// var tv = gui.NewTextView("github.com/WRABZY/gui")

	// f.OpenWithExecute(avAnimation.Start)
	f.Open()
}
