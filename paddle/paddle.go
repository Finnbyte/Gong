package paddle

import (
	"fmt"
	"gong/window"
	"gong/UI"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type PaddleBody struct {
    Body [5]int // Each element is an y location, and first element represents head
}

func (body *PaddleBody) Head() int {
    return body.Body[0]
}

func (body *PaddleBody) Tail() int {
    return body.Body[len(body.Body) - 1]
}

func (body *PaddleBody) Center() int {
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

func (p *Paddle) Init(playfield ui.Playfield) {
	p.playfield = playfield

	p.Img = ebiten.NewImage(p.Width, p.Height)
	p.Img.Fill(color.White)

	p.Y -= float64(p.Height / 2) // Account for height to center on the Y axis
	p.ImgOpts.GeoM.Translate(p.X, p.Y)
}

func (p *Paddle) checkCanMove() bool {
    // Checks all things NOT supposed to happen after moving
    // e.g. going out of screen
	// topPixel := p.Img.SubImage(image.Rect(p.Width, 1, p.Width, 1))
	// bottomPixel := p.Img.SubImage(image.Rect(p.Width, p.Height - 1, p.Width, p.Height - 1))
	
	const TOP_LIMIT = 20

	fmt.Printf("Position above head: %f\n", p.Y - p.Speed)

	if p.Y + (p.Speed * 2) > float64(window.Win.Height - 20) {
		return false
	}

	if p.Y - (p.Speed * 2) < TOP_LIMIT {
		return false
	}

	return true
}

func (p *Paddle) move(directionModifier float64) {
	if p.checkCanMove() {
		p.Y += directionModifier
		p.ImgOpts.GeoM.Translate(0, directionModifier)
	}
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
