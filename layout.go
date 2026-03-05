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
	return &layout{vertical: true, throughElementIndex: -1, throughLayoutIndex: -1}
}

func NewHorizontalLayout() *layout {
	return &layout{vertical: true, throughElementIndex: -1, throughLayoutIndex: -1}
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
		if l.throughElementIndex < len(l.elements)-1 {
			l.throughElementIndex++
		} else {
			l.throughElementIndex = 0
		}
		l.throughX += l.throughXDelta
		l.throughY += l.throughYDelta
		l.elements[l.throughElementIndex].setPosition(l.throughX, l.throughY)
		if l.vertical {
			l.throughYDelta = l.elements[l.throughElementIndex].image().Bounds().Dy()
		} else {
			l.throughXDelta = l.elements[l.throughElementIndex].image().Bounds().Dx()
		}
		return l.elements[l.throughElementIndex], l.throughElementIndex == len(l.elements)-1
	}

	// TODO MOVEMENT THROUGH LAYOUT
	/*
		if len(l.layouts) > 0 {
			if l.throughLayoutIndex < len(l.layouts)-1 {
				l.throughLayoutIndex++
				element, _ = l.layouts[l.throughLayoutIndex].ThroughLayout()
				l.throughX += l.throughXDelta
				l.throughY += l.throughYDelta
				element.setPosition(l.throughX, l.throughY)
				if l.vertical {
					l.throughYDelta = l.elements[l.throughElementIndex].image().Bounds().Dy()
				} else {
					l.throughXDelta = l.elements[l.throughElementIndex].image().Bounds().Dx()
				}
				last = !(l.throughLayoutIndex < len(l.layouts)-1)
				return
			}

			l.throughLayoutIndex = -1
		}*/

	return nil, true
}
