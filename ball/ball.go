package ball

import (
	"image/color"
	"gong/window"
	"github.com/hajimehoshi/ebiten/v2"
)

type BallVelocity struct {
	x, y float64
}

type Ball struct {
	Width, Height int
	Pos          [2]float64
	Velocity BallVelocity
	HasHitPlayer bool
	Img *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

func (b *Ball) Init() {
	b.Img = ebiten.NewImage(b.Width, b.Height)

	pinkColor := color.RGBA{R: 255, B: 203, G: 192, A: 1}
	b.Img.Fill(pinkColor)

	// Centers, accounting for ball's size
	b.Pos[0] -= float64(b.Width / 2)
	b.Pos[1] -= float64(b.Height / 2)

	// Initializes ball to center of window area
	b.ImgOpts.GeoM.Translate(window.Win.CenterX(), window.Win.CenterY())
}

func (b *Ball) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.Img, &b.ImgOpts)
}
