package paddle

import (
	ui "gong/components/UI"
	. "gong/components/screen"
	"gong/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

const SPEED = 6

type Paddle struct {
	X           int
	Y           int
	playfield   ui.Playfield
	StrokeWidth int
	Width       int
	Height      int
	velocityY   int
}

func (p *Paddle) update() {
	minY, maxY := p.playfield.Height-1, Screen.Height-p.playfield.Height
	newYVelocity := p.Y + p.velocityY
	newYPos := utils.Clamp(newYVelocity, maxY-p.Height, minY+p.StrokeWidth)
	p.Y = newYPos
}

func (p *Paddle) MoveUp() {
	p.velocityY = -SPEED
	p.update()
}

func (p *Paddle) MoveDown() {
	p.velocityY = +SPEED
	p.update()
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	vector.StrokeRect(screen, float32(p.X), float32(p.Y), float32(p.Width), float32(p.Height), 3, colornames.White, false)
}
