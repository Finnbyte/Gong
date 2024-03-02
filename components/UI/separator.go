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
	const LINE_HEIGHT = 12
	const LINE_GAP = 30
	for y := 8; y < screenHeight; y += LINE_GAP {
		vector.StrokeLine(screen, float32(screenWidth)/2, float32(y), float32(screenWidth)/2, float32(y)+LINE_HEIGHT, float32(s.Width), colornames.White, false)
	}
}
