package main

import (
	"os"

	"github.com/wrabzy/gui"
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
	var viewButtons1ControlL1 = gui.NewImageView().SetImage(assetImgControlL1)
	var viewButtons1OperatorTop = gui.NewImageView().SetImage(assetImgOperatorTop)
	var viewButtons1ControlR1 = gui.NewImageView().SetImage(assetImgControlR1)
	var viewButtons1Empty = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons1ControlL2 = gui.NewImageView().SetImage(assetImgControlL2)
	var viewButtons1ActionTop = gui.NewImageView().SetImage(assetImgActionTop)
	var viewButtons1ControlR2 = gui.NewImageView().SetImage(assetImgControlR2)
	var viewButtons1BorderRightV = gui.NewImageView().SetImage(assetImgBorderV)
	layoutButtons1.AddViews(
		viewButtons1BorderLeftV,
		viewButtons1ControlL1,
		viewButtons1OperatorTop,
		viewButtons1ControlR1,
		viewButtons1Empty,
		viewButtons1ControlL2,
		viewButtons1ActionTop,
		viewButtons1ControlR2,
		viewButtons1BorderRightV,
	)

	layoutButtons2 := gui.NewHorizontalLayout()
	var viewButtons2BorderLeftV = gui.NewImageView().SetImage(assetImgBorderV)
	var viewButtons2OperatorLeft = gui.NewImageView().SetImage(assetImgOperatorLeft)
	var viewButtons2OperatorCenter = gui.NewImageView().SetImage(assetImgOperatorCenter)
	var viewButtons2OperatorRight = gui.NewImageView().SetImage(assetImgOperatorRight)
	var viewButtons2Empty1 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons2ActionLeft = gui.NewImageView().SetImage(assetImgActionLeft)
	var viewButtons2Empty2 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons2ActionRight = gui.NewImageView().SetImage(assetImgActionRight)
	var viewButtons2BorderRightV = gui.NewImageView().SetImage(assetImgBorderV)
	layoutButtons2.AddViews(
		viewButtons2BorderLeftV,
		viewButtons2OperatorLeft,
		viewButtons2OperatorCenter,
		viewButtons2OperatorRight,
		viewButtons2Empty1,
		viewButtons2ActionLeft,
		viewButtons2Empty2,
		viewButtons2ActionRight,
		viewButtons2BorderRightV,
	)

	layoutButtons3 := gui.NewHorizontalLayout()
	var viewButtons3BorderLeftV = gui.NewImageView().SetImage(assetImgBorderV)
	var viewButtons3Empty1 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons3OperatorBot = gui.NewImageView().SetImage(assetImgOperatorBot)
	var viewButtons3Numpad7 = gui.NewImageView().SetImage(assetImgNumpad7)
	var viewButtons3Numpad8 = gui.NewImageView().SetImage(assetImgNumpad8)
	var viewButtons3Numpad9 = gui.NewImageView().SetImage(assetImgNumpad9)
	var viewButtons3ActionBot = gui.NewImageView().SetImage(assetImgActionBot)
	var viewButtons3Empty2 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons3BorderRightV = gui.NewImageView().SetImage(assetImgBorderV)
	layoutButtons3.AddViews(
		viewButtons3BorderLeftV,
		viewButtons3Empty1,
		viewButtons3OperatorBot,
		viewButtons3Numpad7,
		viewButtons3Numpad8,
		viewButtons3Numpad9,
		viewButtons3ActionBot,
		viewButtons3Empty2,
		viewButtons3BorderRightV,
	)

	layoutButtons4 := gui.NewHorizontalLayout()
	var viewButtons4BorderLeftV = gui.NewImageView().SetImage(assetImgBorderV)
	var viewButtons4Empty1 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons4Empty2 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons4Numpad4 = gui.NewImageView().SetImage(assetImgNumpad4)
	var viewButtons4Numpad5 = gui.NewImageView().SetImage(assetImgNumpad5)
	var viewButtons4Numpad6 = gui.NewImageView().SetImage(assetImgNumpad6)
	var viewButtons4Empty3 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons4Empty4 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons4BorderRightV = gui.NewImageView().SetImage(assetImgBorderV)
	layoutButtons4.AddViews(
		viewButtons4BorderLeftV,
		viewButtons4Empty1,
		viewButtons4Empty2,
		viewButtons4Numpad4,
		viewButtons4Numpad5,
		viewButtons4Numpad6,
		viewButtons4Empty3,
		viewButtons4Empty4,
		viewButtons4BorderRightV,
	)

	layoutButtons5 := gui.NewHorizontalLayout()
	var viewButtons5BorderLeftV = gui.NewImageView().SetImage(assetImgBorderV)
	var viewButtons5Empty1 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons5Empty2 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons5Numpad1 = gui.NewImageView().SetImage(assetImgNumpad1)
	var viewButtons5Numpad2 = gui.NewImageView().SetImage(assetImgNumpad2)
	var viewButtons5Numpad3 = gui.NewImageView().SetImage(assetImgNumpad3)
	var viewButtons5Empty3 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons5Empty4 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons5BorderRightV = gui.NewImageView().SetImage(assetImgBorderV)
	layoutButtons5.AddViews(
		viewButtons5BorderLeftV,
		viewButtons5Empty1,
		viewButtons5Empty2,
		viewButtons5Numpad1,
		viewButtons5Numpad2,
		viewButtons5Numpad3,
		viewButtons5Empty3,
		viewButtons5Empty4,
		viewButtons5BorderRightV,
	)

	layoutButtons6 := gui.NewHorizontalLayout()
	var viewButtons6BorderLeftV = gui.NewImageView().SetImage(assetImgBorderV)
	var viewButtons6Empty1 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons6Empty2 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons6NumpadMinus = gui.NewImageView().SetImage(assetImgNumpadMinus)
	var viewButtons6Numpad0 = gui.NewImageView().SetImage(assetImgNumpad0)
	var viewButtons6NumpadDot = gui.NewImageView().SetImage(assetImgNumpadDot)
	var viewButtons6Empty3 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons6Empty4 = gui.NewImageView().SetImage(assetImgEmpty)
	var viewButtons6BorderRightV = gui.NewImageView().SetImage(assetImgBorderV)
	layoutButtons6.AddViews(
		viewButtons6BorderLeftV,
		viewButtons6Empty1,
		viewButtons6Empty2,
		viewButtons6NumpadMinus,
		viewButtons6Numpad0,
		viewButtons6NumpadDot,
		viewButtons6Empty3,
		viewButtons6Empty4,
		viewButtons6BorderRightV,
	)

	layoutBorderBotH := gui.NewHorizontalLayout()
	var viewBorderBotH = gui.NewImageView().SetImage(assetImgBorderH)
	layoutBorderBotH.AddView(viewBorderBotH)

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

	var form = gui.NewForm("Calculator")
	form.SetLayout(rootLayout)
	form.SetIcons(assetImgIcon16, assetImgIcon32, assetImgIcon48)

	soundSelect := gui.NewInterruptableSound(assetSndSelect).Volume(0.2)

	allButtons := [...]*gui.ImageView{
		viewButtons1ControlL1,
		viewButtons1OperatorTop,
		viewButtons1ControlR1,
		viewButtons1ControlL2,
		viewButtons1ActionTop,
		viewButtons1ControlR2,
		viewButtons2OperatorLeft,
		viewButtons2OperatorRight,
		viewButtons2ActionLeft,
		viewButtons2ActionRight,
		viewButtons3OperatorBot,
		viewButtons3Numpad7,
		viewButtons3Numpad8,
		viewButtons3Numpad9,
		viewButtons3ActionBot,
		viewButtons4Numpad4,
		viewButtons4Numpad5,
		viewButtons4Numpad6,
		viewButtons5Numpad1,
		viewButtons5Numpad2,
		viewButtons5Numpad3,
		viewButtons6NumpadMinus,
		viewButtons6Numpad0,
		viewButtons6NumpadDot,
	}
	for _, view := range allButtons {
		view.SetOnHoverListener(soundSelect.Play)
	}

	viewButtons1ControlL2.SetOnClickListener(
		func() {
			form.SwitchFullscreen()
		},
	)
	viewButtons1ControlR2.SetOnClickListener(
		func() {
			os.Exit(0)
		},
	)
	viewButtons2ActionLeft.SetOnClickListener(calculator.ClearOnePosition)
	viewButtons2ActionRight.SetOnClickListener(calculator.Clear)
	viewButtons3Numpad7.SetOnClickListener(func() {
		calculator.AppendSeven()
	})
	viewButtons3Numpad8.SetOnClickListener(func() {
		calculator.AppendEight()
	})
	viewButtons3Numpad9.SetOnClickListener(func() {
		calculator.AppendNine()
	})
	viewButtons4Numpad4.SetOnClickListener(func() {
		calculator.AppendFour()
	})
	viewButtons4Numpad5.SetOnClickListener(func() {
		calculator.AppendFive()
	})
	viewButtons4Numpad6.SetOnClickListener(func() {
		calculator.AppendSix()
	})
	viewButtons5Numpad1.SetOnClickListener(func() {
		calculator.AppendOne()
	})
	viewButtons5Numpad2.SetOnClickListener(func() {
		calculator.AppendTwo()
	})
	viewButtons5Numpad3.SetOnClickListener(func() {
		calculator.AppendThree()
	})
	viewButtons6NumpadMinus.SetOnClickListener(calculator.SwitchSign)
	viewButtons6Numpad0.SetOnClickListener(calculator.AppendZero)
	viewButtons6NumpadDot.SetOnClickListener(calculator.AppendDot)

	calculator.SetDisplayDigits(
		viewDisplayDigit0,
		viewDisplayDigit1,
		viewDisplayDigit2,
		viewDisplayDigit3,
		viewDisplayDigit4,
		viewDisplayDigit5,
		viewDisplayDigit6,
		viewDisplayDigit7,
		viewDisplayDigit8,
		viewDisplayDigit9,
	)

	calculator.SetDisplayDots(
		viewDisplayDot0,
		viewDisplayDot1,
		viewDisplayDot2,
		viewDisplayDot3,
		viewDisplayDot4,
		viewDisplayDot5,
		viewDisplayDot6,
		viewDisplayDot7,
		viewDisplayDot8,
		viewDisplayDot9,
	)

	calculator.SetDisplayImages(
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
