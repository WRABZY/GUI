package gui

import (
	"fmt"
	"image"
	"math"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type animationView struct {
	identifier int
	anim       *Animation
	x, y       float64
	onClick    func()
}

func NewAnimationView() *animationView {
	return &animationView{identifier: assignID()}
}

func (av *animationView) SetPosition(x, y int) {
	av.x = float64(x)
	av.y = float64(y)
}

func (av *animationView) SetAnimation(animation *Animation) {
	av.anim = animation
}

func (av *animationView) SetOnClickListener(onClick func()) {
	av.onClick = onClick
}

func (av *animationView) Click() {
	if av.onClick != nil {
		av.onClick()
	}
}

func (av *animationView) id() int {
	return av.identifier
}

func (av *animationView) image() *ebiten.Image {
	return av.anim.Frame()
}

func (av *animationView) coordinates() (float64, float64) {
	return av.x, av.y
}

// ---------------------------------------------------

var animationSpriteSheets = make(map[int][]*ebiten.Image)
var nextAnimationSpriteSheetID = math.MinInt

func AddAnimationSpriteSheet(sheet *ebiten.Image, x, y, w, h, frames int) (id int) {
	id = nextAnimationSpriteSheetID
	animationSpriteSheets[id] = make([]*ebiten.Image, frames)
	var x0, x1 int
	y1 := y + h
	for i := range frames {
		x0 = x + i*w
		x1 = x0 + w
		animationSpriteSheets[id][i] = sheet.SubImage(image.Rect(x0, y, x1, y1)).(*ebiten.Image)
	}
	nextAnimationSpriteSheetID++
	return
}

type Animation struct {
	frames                    int
	times                     []time.Duration
	spriteSheetID             int
	nextFrameDirectionForward bool
	loop                      bool
	actualFrameIndex          int
	startFrameIndex           int
	actualFrameIndexLock      sync.Mutex
	isRunning                 bool

	stopChan chan struct{}
	stopLock sync.Mutex
}

func NewAnimation(spriteSheetID int, frames int) *Animation {
	return &Animation{
		frames:                    frames,
		times:                     make([]time.Duration, frames),
		spriteSheetID:             spriteSheetID,
		nextFrameDirectionForward: true,
		stopChan:                  make(chan struct{}),
	}
}

func (a *Animation) WithTimes(milliseconds ...int) *Animation {
	if len(milliseconds) == 1 {
		for i := range a.frames {
			a.times[i] = time.Millisecond * time.Duration(milliseconds[0])
		}
		return a
	}
	if a.frames != len(milliseconds) {
		panic(fmt.Sprintf("the animation has %d frames but the duration is set for %d frames", a.frames, len(milliseconds)))
	}
	for i, ms := range milliseconds {
		a.times[i] = time.Millisecond * time.Duration(ms)
	}
	return a
}

func (a *Animation) WithDirectionForward(forward bool) *Animation {
	a.nextFrameDirectionForward = forward
	return a
}

func (a *Animation) WithLoop(loop bool) *Animation {
	a.loop = loop
	return a
}

func (a *Animation) WithStartFrame(index int) *Animation {
	a.startFrameIndex = index
	return a
}

func (a *Animation) Start() {
	go a.run(true)
}

func (a *Animation) Pause() {
	if a.isRunning {
		a.stopLock.Lock()
		defer a.stopLock.Unlock()

		close(a.stopChan)
		a.stopChan = make(chan struct{})
	}
}

func (a *Animation) Resume() {
	if !a.isRunning {
		go a.run(false)
	}
}

func (a *Animation) IsRunning() bool {
	return a.isRunning
}

func (a *Animation) run(fromBegin bool) {
	a.stopLock.Lock()
	defer a.stopLock.Unlock()

	close(a.stopChan)
	a.stopChan = make(chan struct{})

	a.actualFrameIndexLock.Lock()
	if fromBegin {
		a.actualFrameIndex = a.startFrameIndex
	}

	go func(sc <-chan struct{}) {
		a.isRunning = true
		timer := time.NewTicker(a.times[a.actualFrameIndex])
		for {
			select {
			case <-sc:
				timer.Stop()
				a.isRunning = false
				a.actualFrameIndexLock.Unlock()
				return
			case <-timer.C:
				if a.nextFrameDirectionForward {
					a.actualFrameIndex++
					if a.actualFrameIndex == a.frames {
						if a.loop {
							a.actualFrameIndex = 0
						} else {
							a.nextFrameDirectionForward = false
							a.actualFrameIndex -= 2
						}
					}
				} else {
					a.actualFrameIndex--
					if a.actualFrameIndex == -1 {
						if a.loop {
							a.actualFrameIndex = a.frames - 1
						} else {
							a.nextFrameDirectionForward = true
							a.actualFrameIndex = 1
						}
					}
				}
				timer.Reset(a.times[a.actualFrameIndex])
			}
		}
	}(a.stopChan)
}

func (a *Animation) Frame() *ebiten.Image {
	return animationSpriteSheets[a.spriteSheetID][a.actualFrameIndex]
}
