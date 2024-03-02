package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Separator struct {
	Width int
}

func (s *Separator) Draw(screen *ebiten.Image) {
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	vector.StrokeLine(screen, float32(screenWidth)/2, 0, float32(screenWidth)/2, float32(screenHeight), float32(s.Width), colornames.White, false)
}
