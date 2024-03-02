package paddle

import (
	"fmt"
	"gong/UI"
	"gong/window"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PaddleBody struct {
    body []float64 // Each element is an y location, and first element represents head
	Mid float64
}

func (body *PaddleBody) Head() float64 {
    return body.body[0]
}

func (body *PaddleBody) Tail() float64 {
    return body.body[len(body.body) - 1]
}

func (body *PaddleBody) Center() float64 {
    return body.body[2]
}

type Paddle struct {
    X float64
	Y float64
    Body PaddleBody
	playfield ui.Playfield
    Width int
	Height int
	Speed float64
	Img *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

func (p *Paddle) Init(playfield ui.Playfield, color color.Color) {
	p.playfield = playfield

	// Create and color image representing segment of Paddle
	p.Img = ebiten.NewImage(p.Width, p.Height)
	p.Img.Fill(color)

	// p.Y -= float64(p.Height / 2) // Account for height to center on the Y axis
	p.ImgOpts.GeoM.Translate(p.X, p.Y)

	// Initialize body variable
	p.Body.Body = make([]float64, p.Height)
	for i := 0; i < p.Height; i++ {
		p.Body.Body[i] = p.Y + float64(i)
	}
}

func (p *Paddle) move(directionModifier float64) {
	oldYPos := p.Y
	maxY, minY := p.playfield.Height, window.Win.Height - p.playfield.Height

	// Bottom
	if int(p.Body.Tail()) + int(directionModifier) >= minY {
		const boundaryCollisionPush = 100
		p.Y = float64(minY) - boundaryCollisionPush

	// Top
	} else if int(p.Body.Head()) + int(directionModifier) <= maxY {
		p.Y = float64(maxY)

	// Movement allowed
	} else {
		p.Y += directionModifier
	}

	// Update body
	for i := 0; i < p.Height; i++ {
		p.Body.Body[i] = p.Y + float64(i)
	}

	fmt.Println(p.Body.Tail())

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
