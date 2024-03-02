package screen

type ScreenHelper struct {
	Width, Height int
}

func (w *ScreenHelper) HalfWidth() int {
	return w.Width / 2
}

func (w *ScreenHelper) HalfHeight() int {
	return w.Height / 2
}

func (w ScreenHelper) CenterX() int {
	return w.Width / 2
}

func (w ScreenHelper) CenterY() int {
	return w.Height / 2
}

var (
	Screen = ScreenHelper{}
)
