package gui

type layout struct {
	elements []view
	layouts  []*layout
	vertical bool
}

func NewVerticalLayout() *layout {
	return &layout{vertical: true}
}

func NewHorizontalLayout() *layout {
	return &layout{vertical: true}
}

func (l *layout) AddLayout(nl *layout) {
	l.layouts = append(l.layouts, nl)
}

func (l *layout) AddView(v view) {
	l.elements = append(l.elements, v)
}

func (l *layout) SizeOfContent() (width, height int) {
	width = 0
	height = 0

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

	if l.vertical {
		for _, element := range l.elements {
			if element.image().Bounds().Dx() > width {
				width = element.image().Bounds().Dx()
			}
			height += element.image().Bounds().Dy()
		}
	} else {
		for _, element := range l.elements {
			width += element.image().Bounds().Dx()
			if element.image().Bounds().Dy() > height {
				height = element.image().Bounds().Dy()
			}
		}
	}

	return
}
