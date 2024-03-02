package paddle

import (
	"gong/components/UI"
	"gong/components/window"
	"image/color"
	"gong/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const SPEED = 6

type Paddle struct {
	X         int
	Y         int
	playfield ui.Playfield
	Width     int
	Height    int
	velocityY int
	Img       *ebiten.Image
	ImgOpts   ebiten.DrawImageOptions
}

func (p *Paddle) Init(playfield ui.Playfield, color color.Color) {
	p.playfield = playfield

	p.Img = ebiten.NewImage(p.Width, p.Height)
	p.Img.Fill(color)

	// p.Y -= float64(p.Height / 2) // Account for height to center on the Y axis
	p.ImgOpts.GeoM.Translate(float64(p.X), float64(p.Y))
}

func (p *Paddle) move() {
	oldYPos := p.Y
	minY, maxY := p.playfield.Height-1, window.Win.Height - p.playfield.Height

	newYPos := utils.Clamp(p.Y + p.velocityY, maxY - p.Height, minY)
	p.Y = newYPos

	// Update paddle on screen
	p.ImgOpts.GeoM.Translate(0, float64(p.Y - oldYPos))
}

func (p *Paddle) MoveUp() {
	p.velocityY = -SPEED
	p.move()
}

func (p *Paddle) MoveDown() {
	p.velocityY = +SPEED
	p.move()
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.Img, &p.ImgOpts)
}

func (p Paddle) Clear(screen *ebiten.Image) {
	p.Img.Clear()
}
