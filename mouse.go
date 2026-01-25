package gui

import (
	"sync"
)

type Clickable interface {
	SetOnClickListener(func())
	Click()
}

type ClickableBound struct {
	X0, Y0, X1, Y1 int
}

type MouseState struct {
	sync.RWMutex
	pressed bool
}

func (ms *MouseState) Pressed() bool {
	ms.RLock()
	defer ms.RUnlock()

	return ms.pressed
}

func (ms *MouseState) Press() {
	ms.Lock()
	defer ms.Unlock()

	ms.pressed = true
}

func (ms *MouseState) Release() {
	ms.Lock()
	defer ms.Unlock()

	ms.pressed = false
}
