package ball

import (
	"gong/paddle"
	"gong/window"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type BallVelocity struct {
	X, Y float64
}

type BallPosition struct {
	X, Y float64
}

type BallDirection int

func (bd BallDirection) Int() int {
	return int(bd)
}

type Ball struct {
	Radius int
	Speed, InitialSpeed int
	Pos BallPosition
	direction BallDirection
	HasHitPlayer bool
	Img *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

const (
	UNDEFINED BallDirection = 0
	LEFT BallDirection = 1
	RIGHT BallDirection = -1 
)

func (b *Ball) Init(color color.Color) {
	b.Img = ebiten.NewImage(b.Radius, b.Radius)

	b.Img.Fill(color)

	// Centers, accounting for ball's size
	b.Pos.X -= float64(b.Radius / 2)
	b.Pos.Y -= float64(b.Radius / 2)

	// Initializes ball to center of window area
	b.ImgOpts.GeoM.Translate(window.Win.CenterX(), window.Win.CenterY())
}

func (b *Ball) SwapDirection() {
	switch b.direction {
	case LEFT:
	    b.direction = RIGHT
	case RIGHT:
	    b.direction = LEFT
	}
}

func (b *Ball) Update(rightPaddle, leftPaddle paddle.Paddle) {
	var currentSpeed int

	if b.HasHitPlayer {
		currentSpeed = b.Speed
	} else {
		currentSpeed = b.InitialSpeed
	}

	if b.direction == UNDEFINED {
		b.direction = LEFT
	}

	oldPos := BallPosition{ X: b.Pos.X, Y: b.Pos.Y }

	if int(b.Pos.X + float64(currentSpeed)) >= int(rightPaddle.X - float64(rightPaddle.Width*3)) || int(b.Pos.X + float64(currentSpeed)) <= int(leftPaddle.X) {
		if !b.HasHitPlayer {
			b.HasHitPlayer = true
			currentSpeed = b.Speed
		}
		b.SwapDirection()
	}

	b.Pos.X += float64(currentSpeed * b.direction.Int())

	b.ImgOpts.GeoM.Translate(b.Pos.X - oldPos.X, b.Pos.Y - oldPos.Y)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.Img, &b.ImgOpts)
}
