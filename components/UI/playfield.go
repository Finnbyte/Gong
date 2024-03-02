package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Playfield struct {
	Height int
}

func (pf *Playfield) Draw(screen *ebiten.Image) {
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	vector.StrokeLine(screen, 0, 0, float32(screenWidth), 0, float32(pf.Height), colornames.White, false)
	vector.StrokeLine(screen, 0, float32(screenHeight), float32(screenWidth), float32(screenHeight), 2, colornames.White, false)
}
