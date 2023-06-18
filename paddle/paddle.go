package paddle

import (
	"fmt"
	"gong/window"
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
    Body PaddleBody
    Width int
	Speed int
	Img *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

func (p *Paddle) Init() {
	p.Img = ebiten.NewImage(p.Width, len(p.Body.Body) * 20)
	p.Img.Fill(color.White)

	p.ImgOpts.GeoM.Translate(float64(p.X), float64(p.Body.Head()))
}

func (p *Paddle) checkCanMove() bool {
    // Checks all things NOT supposed to happen after moving
    // e.g. going out of screen
	if (p.Body.Head() + p.Speed) > window.Win.Height {
		fmt.Println("Pää keissi")
		return false
	}

	if (p.Body.Tail() + p.Speed) > window.Win.Height * 2 - 2 {
		fmt.Println("Tail keissi")
		return false
	}
	return true
}

func (p *Paddle) move(directionModifier int) {
	if p.checkCanMove() {
		for i := 0; i < len(p.Body.Body); i++ {
			p.Body.Body[i] += directionModifier
		}
	}
}

func (p *Paddle) MoveUp() {
	p.move(-p.Speed)
}

func (p *Paddle) MoveDown() {
	p.move(+p.Speed)
}

func (p Paddle) Draw(screen *ebiten.Image) {
	for i := 0; i < len(p.Body.Body); i++ {
		bodyPartOpts := ebiten.DrawImageOptions{}
		bodyPartOpts.GeoM.Translate(float64(p.X), float64(p.Body.Body[i]))
		screen.DrawImage(p.Img, &bodyPartOpts)
	}
}

func (p Paddle) Clear(screen *ebiten.Image) {
	p.Img.Clear()
}
