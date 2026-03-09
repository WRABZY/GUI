package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/wrabzy/gui"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {

	layoutBorderTopH := gui.NewHorizontalLayout()
	var viewBorderTopH = gui.NewImageView().SetImage(assetImgBorderH)
	layoutBorderTopH.AddView(viewBorderTopH)

	layoutDisplay := gui.NewHorizontalLayout()
	var viewBorderDisplayLeftV = gui.NewImageView().SetImage(assetImgBorderDisplayVertical)
	var viewDisplayDigitFiller = gui.NewImageView().SetImage(assetImgDisplayDigitFiller)
	var viewDisplayDigit0 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot0 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit1 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot1 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit2 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot2 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit3 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot3 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit4 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot4 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit5 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot5 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit6 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot6 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit7 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot7 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit8 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot8 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewDisplayDigit9 = gui.NewImageView().SetImage(assetImgDisplayDigitEmpty)
	var viewDisplayDot9 = gui.NewImageView().SetImage(assetImgDisplayDotEmpty)
	var viewBorderDisplayRightV = gui.NewImageView().SetImage(assetImgBorderDisplayVertical)
	layoutDisplay.AddViews(
		viewBorderDisplayLeftV,
		viewDisplayDigitFiller,
		viewDisplayDigit0,
		viewDisplayDot0,
		viewDisplayDigit1,
		viewDisplayDot1,
		viewDisplayDigit2,
		viewDisplayDot2,
		viewDisplayDigit3,
		viewDisplayDot3,
		viewDisplayDigit4,
		viewDisplayDot4,
		viewDisplayDigit5,
		viewDisplayDot5,
		viewDisplayDigit6,
		viewDisplayDot6,
		viewDisplayDigit7,
		viewDisplayDot7,
		viewDisplayDigit8,
		viewDisplayDot8,
		viewDisplayDigit9,
		viewDisplayDot9,
		viewBorderDisplayRightV,
	)

	layoutBorderDisplayBotH := gui.NewHorizontalLayout()
	var viewBorderDisplayBotH = gui.NewImageView().SetImage(assetImgBorderH)
	layoutBorderDisplayBotH.AddView(viewBorderDisplayBotH)

	layoutButtons1 := gui.NewHorizontalLayout()
	var viewButtons1BorderLeftV = gui.NewImageView().SetImage(assetImgBorderV)
	var buttonsRow1L1 = gui.NewImageView().SetImage(assetImgControlL1)
	var buttonsRow1Up = gui.NewImageView().SetImage(assetImgOperatorTop)
	var buttonsRow1R1 = gui.NewImageView().SetImage(assetImgControlR1)
	var buttonsRow1Empty = gui.NewImageView().SetImage(assetImgEmpty)
	var buttonsRow1L2 = gui.NewImageView().SetImage(assetImgControlL2)
	var buttonsRow1Y = gui.NewImageView().SetImage(assetImgActionTop)
	var buttonsRow1R2 = gui.NewImageView().SetImage(assetImgControlR2)
	var buttonsRow1VerticalBorderRight = gui.NewImageView().SetImage(assetImgBorderV)
	layoutButtons1.AddViews(
		viewButtons1BorderLeftV,
		buttonsRow1L1,
		buttonsRow1Up,
		buttonsRow1R1,
		buttonsRow1Empty,
		buttonsRow1L2,
		buttonsRow1Y,
		buttonsRow1R2,
		buttonsRow1VerticalBorderRight,
	)

	layoutButtons2 := gui.NewHorizontalLayout()
	layoutButtons3 := gui.NewHorizontalLayout()
	layoutButtons4 := gui.NewHorizontalLayout()
	layoutButtons5 := gui.NewHorizontalLayout()
	layoutButtons6 := gui.NewHorizontalLayout()
	layoutBorderBotH := gui.NewHorizontalLayout()

	rootLayout := gui.NewVerticalLayout()
	rootLayout.AddLayout(layoutBorderTopH)
	rootLayout.AddLayout(layoutDisplay)
	rootLayout.AddLayout(layoutBorderDisplayBotH)
	rootLayout.AddLayout(layoutButtons1)
	rootLayout.AddLayout(layoutButtons2)
	rootLayout.AddLayout(layoutButtons3)
	rootLayout.AddLayout(layoutButtons4)
	rootLayout.AddLayout(layoutButtons5)
	rootLayout.AddLayout(layoutButtons6)
	rootLayout.AddLayout(layoutBorderBotH)

	form.SetLayout(rootLayout)
	form.SetIcons(assetImgIcon16, assetImgIcon32, assetImgIcon48)

	buttonsRow1L2.SetOnClickListener(
		func() {
			form.SwitchFullscreen()
		},
	)
	buttonsRow1R2.SetOnClickListener(
		func() {
			os.Exit(0)
		},
	)

	calculator.SetDisplayDigits(
		viewDisplayDigit0,
		displayDigitView1,
		displayDigitView2,
		displayDigitView3,
		displayDigitView4,
		displayDigitView5,
		displayDigitView6,
		displayDigitView7,
		displayDigitView8,
		displayDigitView9,
	)

	calculator.SetDisplayDots(
		viewDisplayDot0,
		displayDot1,
		displayDot2,
		displayDot3,
		displayDot4,
		displayDot5,
		displayDot6,
		displayDot7,
		displayDot8,
		displayDot9,
	)

	calculator.SetDisplaassetImgActionTops(
		assetImgDisplayDigit0,
		assetImgDisplayDigit1,
		assetImgDisplayDigit2,
		assetImgDisplayDigit3,
		assetImgDisplayDigit4,
		assetImgDisplayDigit5,
		assetImgDisplayDigit6,
		assetImgDisplayDigit7,
		assetImgDisplayDigit8,
		assetImgDisplayDigit9,
		assetImgDisplayDot,
		assetImgDisplayDigitE,
		assetImgDisplayDigitR,
		assetImgDisplayDigitO,
		assetImgDisplayDigitMinus,
		assetImgDisplayDigitEmpty,
		assetImgDisplayDotEmpty,
	)

	var buttonsRow2VerticalBorderLeft = gui.NewImageView()
	var buttonsRow2Left = gui.NewImageView()
	var buttonsRow2Center = gui.NewImageView()
	var buttonsRow2Right = gui.NewImageView()
	var buttonsRow2Empty1 = gui.NewImageView()
	var buttonsRow2X = gui.NewImageView()
	var buttonsRow2Empty2 = gui.NewImageView()
	var buttonsRow2B = gui.NewImageView()
	var buttonsRow2VerticalBorderRight = gui.NewImageView()
	buttonsRow2VerticalBorderLeft.SetImage(assetImgBorderV)
	buttonsRow2Left.SetImage(leftImage)
	buttonsRow2Center.SetImage(centerImage)
	buttonsRow2Right.SetImage(rightImage)
	buttonsRow2Empty1.SetImage(assetImgEmpty)
	buttonsRow2X.SetImage(xImage)
	buttonsRow2Empty2.SetImage(assetImgEmpty)
	buttonsRow2B.SetImage(bImage)
	buttonsRow2VerticalBorderRight.SetImage(assetImgBorderV)
	layoutButtons2.AddView(buttonsRow2VerticalBorderLeft)
	layoutButtons2.AddView(buttonsRow2Left)
	layoutButtons2.AddView(buttonsRow2Center)
	layoutButtons2.AddView(buttonsRow2Right)
	layoutButtons2.AddView(buttonsRow2Empty1)
	layoutButtons2.AddView(buttonsRow2X)
	layoutButtons2.AddView(buttonsRow2Empty2)
	layoutButtons2.AddView(buttonsRow2B)
	layoutButtons2.AddView(buttonsRow2VerticalBorderRight)
	buttonsRow2X.SetOnClickListener(calculator.ClearOnePosition)
	buttonsRow2B.SetOnClickListener(calculator.Clear)

	var buttonsRow3VerticalBorderLeft = gui.NewImageView()
	var buttonsRow3Empty1 = gui.NewImageView()
	var buttonsRow3Bot = gui.NewImageView()
	var buttonsRow3B7 = gui.NewImageView()
	var buttonsRow3B8 = gui.NewImageView()
	var buttonsRow3B9 = gui.NewImageView()
	var buttonsRow3A = gui.NewImageView()
	var buttonsRow3Empty2 = gui.NewImageView()
	var buttonsRow3VerticalBorderRight = gui.NewImageView()
	buttonsRow3VerticalBorderLeft.SetImage(assetImgBorderV)
	buttonsRow3Empty1.SetImage(assetImgEmpty)
	buttonsRow3Bot.SetImage(botImage)
	buttonsRow3B7.SetImage(b7Image)
	buttonsRow3B8.SetImage(b8Image)
	buttonsRow3B9.SetImage(b9Image)
	buttonsRow3A.SetImage(aImage)
	buttonsRow3Empty2.SetImage(assetImgEmpty)
	buttonsRow3VerticalBorderRight.SetImage(assetImgBorderV)
	layoutButtons3.AddView(buttonsRow3VerticalBorderLeft)
	layoutButtons3.AddView(buttonsRow3Empty1)
	layoutButtons3.AddView(buttonsRow3Bot)
	layoutButtons3.AddView(buttonsRow3B7)
	layoutButtons3.AddView(buttonsRow3B8)
	layoutButtons3.AddView(buttonsRow3B9)
	layoutButtons3.AddView(buttonsRow3A)
	layoutButtons3.AddView(buttonsRow3Empty2)
	layoutButtons3.AddView(buttonsRow3VerticalBorderRight)
	buttonsRow3B7.SetOnClickListener(func() {
		calculator.AppendSeven()
	})
	buttonsRow3B8.SetOnClickListener(func() {
		calculator.AppendEight()
	})
	buttonsRow3B9.SetOnClickListener(func() {
		calculator.AppendNine()
	})

	file, err := os.Open("example/calculator/assets/select.wav")
	if err != nil {
		panic(err)
	}
	sound, err := wav.DecodeF32(file)
	form.AddSound(sound)

	buttonsRow3B7.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})
	buttonsRow3B8.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})
	buttonsRow3B9.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})

	var buttonsRow4VerticalBorderLeft = gui.NewImageView()
	var buttonsRow4Empty1 = gui.NewImageView()
	var buttonsRow4Empty2 = gui.NewImageView()
	var buttonsRow4B4 = gui.NewImageView()
	var buttonsRow4B5 = gui.NewImageView()
	var buttonsRow4B6 = gui.NewImageView()
	var buttonsRow4Empty3 = gui.NewImageView()
	var buttonsRow4Empty4 = gui.NewImageView()
	var buttonsRow4VerticalBorderRight = gui.NewImageView()
	buttonsRow4VerticalBorderLeft.SetImage(assetImgBorderV)
	buttonsRow4Empty1.SetImage(assetImgEmpty)
	buttonsRow4Empty2.SetImage(assetImgEmpty)
	buttonsRow4B4.SetImage(b4Image)
	buttonsRow4B5.SetImage(b5Image)
	buttonsRow4B6.SetImage(b6Image)
	buttonsRow4Empty3.SetImage(assetImgEmpty)
	buttonsRow4Empty4.SetImage(assetImgEmpty)
	buttonsRow4VerticalBorderRight.SetImage(assetImgBorderV)
	layoutButtons4.AddView(buttonsRow4VerticalBorderLeft)
	layoutButtons4.AddView(buttonsRow4Empty1)
	layoutButtons4.AddView(buttonsRow4Empty2)
	layoutButtons4.AddView(buttonsRow4B4)
	layoutButtons4.AddView(buttonsRow4B5)
	layoutButtons4.AddView(buttonsRow4B6)
	layoutButtons4.AddView(buttonsRow4Empty3)
	layoutButtons4.AddView(buttonsRow4Empty4)
	layoutButtons4.AddView(buttonsRow4VerticalBorderRight)
	buttonsRow4B4.SetOnClickListener(func() {
		calculator.AppendFour()
	})
	buttonsRow4B5.SetOnClickListener(func() {
		calculator.AppendFive()
	})
	buttonsRow4B6.SetOnClickListener(func() {
		calculator.AppendSix()
	})

	buttonsRow4B4.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})
	buttonsRow4B5.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})
	buttonsRow4B6.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})

	var buttonsRow5VerticalBorderLeft = gui.NewImageView()
	var buttonsRow5Empty1 = gui.NewImageView()
	var buttonsRow5Empty2 = gui.NewImageView()
	var buttonsRow5B1 = gui.NewImageView()
	var buttonsRow5B2 = gui.NewImageView()
	var buttonsRow5B3 = gui.NewImageView()
	var buttonsRow5Empty3 = gui.NewImageView()
	var buttonsRow5Empty4 = gui.NewImageView()
	var buttonsRow5VerticalBorderRight = gui.NewImageView()
	buttonsRow5VerticalBorderLeft.SetImage(assetImgBorderV)
	buttonsRow5Empty1.SetImage(assetImgEmpty)
	buttonsRow5Empty2.SetImage(assetImgEmpty)
	buttonsRow5B1.SetImage(b1Image)
	buttonsRow5B2.SetImage(b2Image)
	buttonsRow5B3.SetImage(b3Image)
	buttonsRow5Empty3.SetImage(assetImgEmpty)
	buttonsRow5Empty4.SetImage(assetImgEmpty)
	buttonsRow5VerticalBorderRight.SetImage(assetImgBorderV)
	layoutButtons5.AddView(buttonsRow5VerticalBorderLeft)
	layoutButtons5.AddView(buttonsRow5Empty1)
	layoutButtons5.AddView(buttonsRow5Empty2)
	layoutButtons5.AddView(buttonsRow5B1)
	layoutButtons5.AddView(buttonsRow5B2)
	layoutButtons5.AddView(buttonsRow5B3)
	layoutButtons5.AddView(buttonsRow5Empty3)
	layoutButtons5.AddView(buttonsRow5Empty4)
	layoutButtons5.AddView(buttonsRow5VerticalBorderRight)
	buttonsRow5B1.SetOnClickListener(func() {
		calculator.AppendOne()
	})
	buttonsRow5B2.SetOnClickListener(func() {
		calculator.AppendTwo()
	})
	buttonsRow5B3.SetOnClickListener(func() {
		calculator.AppendThree()
	})

	buttonsRow5B1.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})
	buttonsRow5B2.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})
	buttonsRow5B3.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})

	var buttonsRow6VerticalBorderLeft = gui.NewImageView()
	var buttonsRow6Empty1 = gui.NewImageView()
	var buttonsRow6Empty2 = gui.NewImageView()
	var buttonsRow6UnaryMinus = gui.NewImageView()
	var buttonsRow6B0 = gui.NewImageView()
	var buttonsRow6Bdot = gui.NewImageView()
	var buttonsRow6Empty3 = gui.NewImageView()
	var buttonsRow6Empty4 = gui.NewImageView()
	var buttonsRow6VerticalBorderRight = gui.NewImageView()
	buttonsRow6VerticalBorderLeft.SetImage(assetImgBorderV)
	buttonsRow6Empty1.SetImage(assetImgEmpty)
	buttonsRow6Empty2.SetImage(assetImgEmpty)
	buttonsRow6UnaryMinus.SetImage(unaryMinusImage)
	buttonsRow6B0.SetImage(b0Image)
	buttonsRow6Bdot.SetImage(dotImage)
	buttonsRow6Empty3.SetImage(assetImgEmpty)
	buttonsRow6Empty4.SetImage(assetImgEmpty)
	buttonsRow6VerticalBorderRight.SetImage(assetImgBorderV)
	layoutButtons6.AddView(buttonsRow6VerticalBorderLeft)
	layoutButtons6.AddView(buttonsRow6Empty1)
	layoutButtons6.AddView(buttonsRow6Empty2)
	layoutButtons6.AddView(buttonsRow6UnaryMinus)
	layoutButtons6.AddView(buttonsRow6B0)
	layoutButtons6.AddView(buttonsRow6Bdot)
	layoutButtons6.AddView(buttonsRow6Empty3)
	layoutButtons6.AddView(buttonsRow6Empty4)
	layoutButtons6.AddView(buttonsRow6VerticalBorderRight)
	buttonsRow6UnaryMinus.SetOnClickListener(calculator.SwitchSign)
	buttonsRow6B0.SetOnClickListener(calculator.AppendZero)
	buttonsRow6Bdot.SetOnClickListener(calculator.AppendDot)

	buttonsRow6B0.SetOnHoverListener(func() {
		if form.IsPlaying() {
			return
		}
		form.PlaySound()
	})

	var botHBorder = gui.NewImageView()
	botHBorder.SetImage(assetImgBorderH)
	layoutBorderBotH.AddView(botHBorder)

	var form = gui.NewForm("Calculator")

	/*
		ivImage, _, err := ebitenutil.NewImageFromFile("example/calculator/assets/48x48.png")
		if err != nil {
			log.Fatal(err)
		}
		var iv = gui.NewImageView()
		iv.SetImage(ivImage)

		iv2Image, _, err := ebitenutil.NewImageFromFile("example/calculator/assets/200x200.png")
		if err != nil {
			log.Fatal(err)
		}
		var iv2 = gui.NewImageView()
		iv2.SetImage(iv2Image)

		avImg, _, err := ebitenutil.NewImageFromFile("example/calculator/assets/a100x200.png")
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

	// form.OpenWithExecute(avAnimation.Start)
	form.Open()
}
