package main

import (
	"log"
	"math/big"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wrabzy/gui"
)

type Сalculator struct {
	operand       *big.Float
	onDisplay     string
	displayDigits []*gui.ImageView
	displayDots   []*gui.ImageView
	displayImages []*ebiten.Image
}

var calculator = &Сalculator{
	operand: big.NewFloat(0),
}

func (c *Сalculator) SetDisplayDigits(digits ...*gui.ImageView) {
	c.displayDigits = digits
}

func (c *Сalculator) SetDisplayDots(dots ...*gui.ImageView) {
	c.displayDots = dots
}

// SetDisplayImages accepts images in order 0-9 - digits, 10 - dot, 11 - E, 12 - r, 13 - o, 14 - minus, 15 - empty digit, 16 - empty dot
func (c *Сalculator) SetDisplayImages(images ...*ebiten.Image) {
	c.displayImages = images
}

func (c *Сalculator) AppendNine() {
	c.append("9")
}

func (c *Сalculator) AppendEight() {
	c.append("8")
}

func (c *Сalculator) AppendSeven() {
	c.append("7")
}
func (c *Сalculator) AppendSix() {
	c.append("6")
}
func (c *Сalculator) AppendFive() {
	c.append("5")
}

func (c *Сalculator) AppendFour() {
	c.append("4")
}

func (c *Сalculator) AppendThree() {
	c.append("3")
}

func (c *Сalculator) AppendTwo() {
	c.append("2")
}

func (c *Сalculator) AppendOne() {
	c.append("1")
}

func (c *Сalculator) AppendZero() {
	c.append("0")
}

func (c *Сalculator) AppendDot() {
	if len(c.onDisplay) == 0 || strings.Contains(c.onDisplay, ".") {
		return
	}
	c.append(".")
}

func (c *Сalculator) SwitchSign() {
	defer func() {
		// TODO Delete it
		log.Println("On display:", c.onDisplay)
	}()
	if len(c.onDisplay) > 0 {
		if c.onDisplay[0] == '-' {
			c.onDisplay = c.onDisplay[1:]
		} else {
			borderSize := 10
			if strings.Contains(c.onDisplay, "+") {
				borderSize++
			}
			if strings.Contains(c.onDisplay, ".") {
				borderSize++
			}
			if len(c.onDisplay) < borderSize {
				c.onDisplay = "-" + c.onDisplay
			}
		}
		c.RefreshDisplay()
	}
}

func (c *Сalculator) append(unit string) {
	defer func() {
		// TODO Delete it
		log.Println("On display:", c.onDisplay)
	}()

	borderSize := 10
	if strings.Contains(c.onDisplay, "+") {
		borderSize++
	}
	if strings.Contains(c.onDisplay, ".") {
		borderSize++
	}
	if len(c.onDisplay) < borderSize {
		newValue := c.onDisplay + unit
		var ok bool
		c.operand, ok = c.operand.SetString(newValue)
		if !ok {
			c.Clear()
			c.ShowErr()
			return
		}

		c.onDisplay = newValue
		c.RefreshDisplay()
	}
}

func (c *Сalculator) ShowErr() {
	c.clearDisplay()
	for i, digitView := range c.displayDigits {
		switch i {
		case 5:
			digitView.SetImage(c.displayImages[11])
		case 6, 7, 9:
			digitView.SetImage(c.displayImages[12])
		case 8:
			digitView.SetImage(c.displayImages[13])
		}
	}
}

func (c *Сalculator) RefreshDisplay() {
	c.clearDisplay()
	if len(c.onDisplay) > 0 {
		borderSize := 11
		if strings.Contains(c.onDisplay, "+") {
			borderSize++
		}
		if strings.Contains(c.onDisplay, ".") {
			borderSize++
		}
		if len(c.onDisplay) < borderSize {
			iString := len(c.onDisplay) - 1
			iDisplay := 9
			for iDisplay > -1 && iString > -1 {
				switch c.onDisplay[iString] {
				case '0':
					c.displayDigits[iDisplay].SetImage(c.displayImages[0])
					iDisplay--
				case '1':
					c.displayDigits[iDisplay].SetImage(c.displayImages[1])
					iDisplay--
				case '2':
					c.displayDigits[iDisplay].SetImage(c.displayImages[2])
					iDisplay--
				case '3':
					c.displayDigits[iDisplay].SetImage(c.displayImages[3])
					iDisplay--
				case '4':
					c.displayDigits[iDisplay].SetImage(c.displayImages[4])
					iDisplay--
				case '5':
					c.displayDigits[iDisplay].SetImage(c.displayImages[5])
					iDisplay--
				case '6':
					c.displayDigits[iDisplay].SetImage(c.displayImages[6])
					iDisplay--
				case '7':
					c.displayDigits[iDisplay].SetImage(c.displayImages[7])
					iDisplay--
				case '8':
					c.displayDigits[iDisplay].SetImage(c.displayImages[8])
					iDisplay--
				case '9':
					c.displayDigits[iDisplay].SetImage(c.displayImages[9])
					iDisplay--
				case '.':
					c.displayDots[iDisplay].SetImage(c.displayImages[10])
				case 'E':
					c.displayDigits[iDisplay].SetImage(c.displayImages[11])
					iDisplay--
				case '-':
					c.displayDigits[iDisplay].SetImage(c.displayImages[14])
					iDisplay--
				case '+':
				}
				iString--
			}
		} else {
			c.ShowErr()
		}
	}
}

func (c *Сalculator) clearDisplay() {
	for i, digitView := range c.displayDigits {
		digitView.SetImage(c.displayImages[15])
		c.displayDots[i].SetImage(c.displayImages[16])
	}
}

func (c *Сalculator) Clear() {
	c.operand = big.NewFloat(0)
	c.onDisplay = ""
	c.clearDisplay()
}

func (c *Сalculator) ClearOnePosition() {
	if len(c.onDisplay) < 2 {
		c.Clear()
	} else {
		newValue := c.onDisplay[:len(c.onDisplay)-1]
		var ok bool
		c.operand, ok = c.operand.SetString(newValue)
		if !ok {
			c.Clear()
			c.ShowErr()
			return
		}

		c.onDisplay = newValue
		c.RefreshDisplay()
	}
}
