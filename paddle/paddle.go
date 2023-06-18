package paddle

import (
	"fmt"
	"gong/window"
	"gong/UI"
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
    Body PaddleBody
	playfield ui.Playfield
    Width int
	Height int
	Speed float64
	Img *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

func (p *Paddle) Init(playfield ui.Playfield) {
	// Assign playfield
	p.playfield = playfield

	// Create and color image representing segment of Paddle
	p.Img = ebiten.NewImage(p.Width, p.Height)
	red := color.RGBA{R: 255, G: 0, B: 0, A: 1}
	p.Img.Fill(red)

	// Account for height to center on the Y axis
	p.Body.Mid -= float64(p.Height / 2) 

	// Initialize body parts
	p.Body.body = make([]float64, p.Height)
	p.Body.body[0] = p.Body.Mid - 1.0
	p.Body.body[1] = p.Body.Mid - 2.0
	p.Body.body[2] = p.Body.Mid
	p.Body.body[3] = p.Body.Mid + 1.0
	p.Body.body[4] = p.Body.Mid + 2.0
}

func (p *Paddle) checkCanMove() bool {
    // Checks all things NOT supposed to happen after moving
    // e.g. going out of screen
	const TOP_LIMIT = 20

	fmt.Printf("Tail position: %f\n", p.Body.Tail())
	fmt.Printf("Head position: %f\n", p.Body.Head())

	if p.Body.Tail() + p.Speed > float64(window.Win.Height - 20) {
		return false
	}

	if p.Body.Head() - p.Speed < TOP_LIMIT {
		return false
	}
	return true
}

func (p *Paddle) move(directionModifier float64) {
	if p.checkCanMove() {
		// Update body
		for i := 0; i < len(p.Body.body); i++ {
			p.Body.body[i] += directionModifier
		}
	}
}

func (p *Paddle) MoveUp() {
	p.move(-p.Speed)
}

func (p *Paddle) MoveDown() {
	p.move(+p.Speed)
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	bodyPartOpts := ebiten.DrawImageOptions{}
	for i := 0; i < len(p.Body.body); i++ {
		bodyPartOpts.GeoM.Translate(0, float64(p.Body.body[i]))
	}

	for i := 0; i < len(p.Body.body); i++ {
		bodyPartOpts.GeoM.Translate(0, float64(p.Body.body[i]))
	    screen.DrawImage(p.Img, &bodyPartOpts)
	}
}

func (p Paddle) Clear(screen *ebiten.Image) {
	p.Img.Clear()
}
