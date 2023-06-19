package paddle

import (
	"gong/window"
	"gong/UI"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type PaddleBody struct {
    Body []float64 // Each element is an y location, and first element represents head
}

func (body *PaddleBody) Head() float64 {
    return body.Body[0]
}

func (body *PaddleBody) Tail() float64 {
    return body.Body[len(body.Body) - 1]
}

func (body *PaddleBody) Center() float64 {
    return body.Body[2]
}

type Paddle struct {
    X float64
	Y float64
    body PaddleBody
	playfield ui.Playfield
    Width int
	Height int
	Speed float64
	Img *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

func (p *Paddle) Init(playfield ui.Playfield, color color.Color) {
	p.playfield = playfield

	p.Img = ebiten.NewImage(p.Width, p.Height)
	p.Img.Fill(color)

	p.Y -= float64(p.Height / 2) // Account for height to center on the Y axis
	p.ImgOpts.GeoM.Translate(p.X, p.Y)

	// Initialize body variable
	p.body.Body = make([]float64, p.Height)
	for i := 0; i < p.Height; i++ {
		p.body.Body[i] = p.Y + 1.0
	}
}

func (p *Paddle) move(directionModifier float64) {
	oldYPos := p.Y
	maxY, minY := 20, window.Win.Height - p.Height - 20
	boundaryCollisionPush := p.Speed * 2 // Pushes paddle by value when colliding with playfield

	// Bottom
	if int(p.body.Tail() + p.Speed) >= minY {
		p.Y = float64(minY) + -boundaryCollisionPush 

	// Top
	} else if int(p.body.Head() - p.Speed) <= maxY {
		p.Y = float64(maxY) + boundaryCollisionPush 

	// Movement allowed
	} else {
		p.Y += directionModifier
	}

	// Update body
	for i := 0; i < p.Height; i++ {
		p.body.Body[i] = p.Y - 1.0
	}

	// Update paddle on screen
	p.ImgOpts.GeoM.Translate(0, p.Y - oldYPos)
}

func (p *Paddle) MoveUp() {
	p.move(-p.Speed)
}

func (p *Paddle) MoveDown() {
	p.move(+p.Speed)
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.Img, &p.ImgOpts)
}

func (p Paddle) Clear(screen *ebiten.Image) {
	p.Img.Clear()
}
