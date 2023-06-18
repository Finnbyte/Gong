package window

type Window struct {
    Width, Height int
}

func (w *Window) HalfWidth() int {
   return w.Width / 2
}

func (w *Window) HalfHeight() int {
   return w.Height / 2
}

func (w Window) CenterX() float64 {
	return float64(w.Width / 2)
}

func (w Window) CenterY() float64 {
	return float64(w.Height / 2)
}

var (
    Win = Window{}
)
