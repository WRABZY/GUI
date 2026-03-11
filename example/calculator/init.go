package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	assetImgActionTop             *ebiten.Image
	assetImgActionLeft            *ebiten.Image
	assetImgActionBot             *ebiten.Image
	assetImgActionRight           *ebiten.Image
	assetImgControlL1             *ebiten.Image
	assetImgControlR1             *ebiten.Image
	assetImgControlL2             *ebiten.Image
	assetImgControlR2             *ebiten.Image
	assetImgDisplayDigitFiller    *ebiten.Image
	assetImgDisplayDigitEmpty     *ebiten.Image
	assetImgDisplayDotEmpty       *ebiten.Image
	assetImgDisplayDigit0         *ebiten.Image
	assetImgDisplayDigit1         *ebiten.Image
	assetImgDisplayDigit2         *ebiten.Image
	assetImgDisplayDigit3         *ebiten.Image
	assetImgDisplayDigit4         *ebiten.Image
	assetImgDisplayDigit5         *ebiten.Image
	assetImgDisplayDigit6         *ebiten.Image
	assetImgDisplayDigit7         *ebiten.Image
	assetImgDisplayDigit8         *ebiten.Image
	assetImgDisplayDigit9         *ebiten.Image
	assetImgDisplayDot            *ebiten.Image
	assetImgDisplayDigitE         *ebiten.Image
	assetImgDisplayDigitR         *ebiten.Image
	assetImgDisplayDigitO         *ebiten.Image
	assetImgDisplayDigitMinus     *ebiten.Image
	assetImgEmpty                 *ebiten.Image
	assetImgBorderH               *ebiten.Image
	assetImgBorderV               *ebiten.Image
	assetImgBorderDisplayVertical *ebiten.Image
	assetImgIcon16                *ebiten.Image
	assetImgIcon32                *ebiten.Image
	assetImgIcon48                *ebiten.Image
	assetImgNumpad0               *ebiten.Image
	assetImgNumpad1               *ebiten.Image
	assetImgNumpad2               *ebiten.Image
	assetImgNumpad3               *ebiten.Image
	assetImgNumpad4               *ebiten.Image
	assetImgNumpad5               *ebiten.Image
	assetImgNumpad6               *ebiten.Image
	assetImgNumpad7               *ebiten.Image
	assetImgNumpad8               *ebiten.Image
	assetImgNumpad9               *ebiten.Image
	assetImgNumpadDot             *ebiten.Image
	assetImgNumpadMinus           *ebiten.Image
	assetImgOperatorTop           *ebiten.Image
	assetImgOperatorBot           *ebiten.Image
	assetImgOperatorLeft          *ebiten.Image
	assetImgOperatorRight         *ebiten.Image
	assetImgOperatorCenter        *ebiten.Image

	assetSndSelect *wav.Stream
)

func init() {
	var err error

	assetImgActionTop, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/actions/top.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgActionLeft, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/actions/left.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgActionBot, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/actions/bot.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgActionRight, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/actions/right.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgControlL1, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/control/l1.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgControlR1, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/control/r1.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgControlL2, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/control/l2.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgControlR2, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/control/r2.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigitFiller, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_filler.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigitEmpty, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_empty.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDotEmpty, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_dot_empty.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit0, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_0.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit1, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_1.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit2, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_2.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit3, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_3.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit4, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_4.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit5, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_5.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit6, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_6.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit7, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_7.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit8, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_8.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigit9, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_9.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDot, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_dot.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigitE, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_e.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigitR, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_r.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigitO, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_digit_o.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgDisplayDigitMinus, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/display/display_minus.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgEmpty, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/extra/empty.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgBorderH, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/extra/border_horizontal.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgBorderV, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/extra/border_vertical.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgBorderDisplayVertical, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/extra/border_vertical_display.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgIcon16, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/icons/icon16.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgIcon32, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/icons/icon32.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgIcon48, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/icons/icon48.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad0, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/0.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad1, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/1.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad2, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/2.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad3, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/3.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad4, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/4.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad5, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/5.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad6, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/6.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad7, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/7.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad8, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/8.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpad9, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/9.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpadDot, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/dot.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgNumpadMinus, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/num_pad/minus.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgOperatorTop, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/op_pad/top.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgOperatorBot, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/op_pad/bot.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgOperatorLeft, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/op_pad/left.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgOperatorRight, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/op_pad/right.png")
	if err != nil {
		log.Fatal(err)
	}
	assetImgOperatorCenter, _, err = ebitenutil.NewImageFromFile("example/calculator/assets/op_pad/center.png")
	if err != nil {
		log.Fatal(err)
	}

	var fileSndSelect *os.File
	fileSndSelect, err = os.Open("example/calculator/assets/sound/select.wav")
	if err != nil {
		log.Fatal(err)
	}
	assetSndSelect, err = wav.DecodeF32(fileSndSelect)
	if err != nil {
		log.Fatal(err)
	}
}
