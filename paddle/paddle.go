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
    X int
	Y int
    body PaddleBody
	playfield ui.Playfield
    Width int
	Height int
	Speed int
	Img *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

func (p *Paddle) Init(playfield ui.Playfield) {
	p.playfield = playfield

	p.Img = ebiten.NewImage(p.Width, p.Height)
	p.Img.Fill(color.White)

	p.ImgOpts.GeoM.Translate(float64(p.X), float64(p.Y))
}

func (p *Paddle) checkCanMove() bool {
    // Checks all things NOT supposed to happen after moving
    // e.g. going out of screen
	// topPixel := p.Img.SubImage(image.Rect(p.Width, 1, p.Width, 1))
	// bottomPixel := p.Img.SubImage(image.Rect(p.Width, p.Height - 1, p.Width, p.Height - 1))

	if float64(p.Y + p.Speed) > float64(window.Win.Height - 20) {
		fmt.Println("shit hit the fan")
		return false
	}

	return true
}

func (p *Paddle) move(directionModifier int) {
	if p.checkCanMove() {
		p.Y += directionModifier
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
