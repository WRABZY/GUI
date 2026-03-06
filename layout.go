package gui

type layout struct {
	elements      []view
	layouts       []*layout
	vertical      bool
	width, height int

	throughElementIndex                              int
	throughLayoutIndex                               int
	throughX, throughY, throughXDelta, throughYDelta int
}

func NewVerticalLayout() *layout {
	return &layout{vertical: true}
}

func NewHorizontalLayout() *layout {
	return &layout{vertical: false}
}

func (l *layout) AddLayout(nl *layout) {
	if len(l.elements) > 0 {
		panic("Layout can contain only elements or only other layouts")
	}
	l.layouts = append(l.layouts, nl)
}

func (l *layout) AddView(v view) {
	if len(l.layouts) > 0 {
		panic("Layout can contain only elements or only other layouts")
	}
	l.elements = append(l.elements, v)
	if l.vertical {
		if v.image().Bounds().Dx() > l.width {
			l.width = v.image().Bounds().Dx()
		}
		l.height += v.image().Bounds().Dy()
	} else {
		l.width += v.image().Bounds().Dx()
		if v.image().Bounds().Dy() > l.height {
			l.height = v.image().Bounds().Dy()
		}
	}
}

func (l *layout) SizeOfContent() (width, height int) {
	if len(l.elements) == 0 {
		if len(l.layouts) == 0 {
			return
		}
		if l.vertical {
			for _, layout := range l.layouts {
				lw, lh := layout.SizeOfContent()
				if lw > width {
					width = lw
				}
				height += lh
			}
		} else {
			for _, layout := range l.layouts {
				lw, lh := layout.SizeOfContent()
				width += lw
				if lh > height {
					height = lh
				}
			}
		}
		return
	}

	return l.width, l.height
}

func (l *layout) ThroughLayout() (element view, last bool) {
	if len(l.elements) > 0 {
		element = l.elements[l.throughElementIndex]
		l.throughX += l.throughXDelta
		l.throughY += l.throughYDelta
		element.setPosition(l.throughX, l.throughY)

		l.throughElementIndex++
		if l.throughElementIndex == len(l.elements) {
			l.throughX = 0
			l.throughY = 0
			l.throughXDelta = 0
			l.throughYDelta = 0
			l.throughElementIndex = 0
			last = true
		} else if l.vertical {
			l.throughYDelta = element.image().Bounds().Dy()
		} else {
			l.throughXDelta = element.image().Bounds().Dx()
		}
		return
	}

	if len(l.layouts) > 0 {
		layout := l.layouts[l.throughLayoutIndex]
		element, last = layout.ThroughLayout()
		if element == nil {
			return nil, true
		}
		elemX, elemY := element.coordinates()
		element.setPosition(int(elemX)+l.throughX, int(elemY)+l.throughY)

		if last {
			l.throughLayoutIndex++

			if l.throughLayoutIndex == len(l.layouts) {
				l.throughXDelta = 0
				l.throughYDelta = 0
				l.throughLayoutIndex = 0
			} else {
				if l.vertical {
					l.throughY += layout.height
				} else {
					l.throughX += layout.width
				}
				last = false
			}
		}
		return
	}

	return nil, true
}
