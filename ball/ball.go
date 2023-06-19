package ball

import (
	"fmt"
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
	xDirection BallDirection
	yDirection BallDirection
	HasHitPlayer bool
	Img *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

const (
	UNDEFINED BallDirection = 0
	LEFT BallDirection = -1
	RIGHT BallDirection = 1 
	UP BallDirection = -1
	DOWN BallDirection = 1
)

func (b *Ball) Init(color color.Color) {
	b.Img = ebiten.NewImage(b.Radius, b.Radius)

	b.Img.Fill(color)

	// Centers, accounting for ball's size
	b.Pos.X -= float64(b.Radius / 2)
	b.Pos.Y -= float64(b.Radius / 2)

	// Sets initial direction of ball
	b.xDirection = UNDEFINED
	b.yDirection = UP

	// Initializes ball to center of window area
	b.ImgOpts.GeoM.Translate(window.Win.CenterX(), window.Win.CenterY())
}

func (b *Ball) SwapDirection(direction *BallDirection) {
	switch *direction {
	case LEFT:
	    *direction = RIGHT
	case RIGHT:
	    *direction = LEFT
	}
}

func (b *Ball) Update(rightPaddle, leftPaddle paddle.Paddle) {
	var currentSpeed int
	if b.HasHitPlayer {
		currentSpeed = b.Speed
	} else {
		currentSpeed = b.InitialSpeed
	}

	oldPos := BallPosition{ X: b.Pos.X, Y: b.Pos.Y }

	// Check playfield collision
	if b.Pos.Y + float64(b.Radius * 2) >= float64(window.Win.Height) || b.Pos.Y - float64(currentSpeed) <= 0 {
		fmt.Println("touched border")
		b.SwapDirection(&b.yDirection)
	}

	// Check paddles collision (X only for now)
	if int(b.Pos.X + float64(currentSpeed)) >= int(rightPaddle.X - float64(rightPaddle.Width*3)) || int(b.Pos.X + float64(currentSpeed)) <= int(leftPaddle.X) {
		b.HasHitPlayer = true
		currentSpeed = b.Speed
		b.SwapDirection(&b.xDirection)
	}

	b.Pos.X += float64(currentSpeed * b.xDirection.Int())
	b.Pos.Y += float64(currentSpeed * b.yDirection.Int())

	b.ImgOpts.GeoM.Translate(b.Pos.X - oldPos.X, b.Pos.Y - oldPos.Y)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.Img, &b.ImgOpts)
}
