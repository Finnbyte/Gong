package ball

import (
	"fmt"
	ui "gong/UI"
	"gong/paddle"
	"gong/player"
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

type BallDirection string

func (bd BallDirection) Int() int {
	switch bd {
	case "LEFT":
		return -1
	case "RIGHT":
		return 1
	case "UP":
		return -1
	case "DOWN":
		return 1
	default: 
		return 0
	}
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
	UNDEFINED BallDirection = "UNDEFINED"
	LEFT BallDirection = "LEFT"
	RIGHT BallDirection = "RIGHT"
	UP BallDirection = "UP" 
	DOWN BallDirection = "DOWN" 
)

func (b *Ball) Init(color color.Color) {
	b.Img = ebiten.NewImage(b.Radius, b.Radius)

	b.Img.Fill(color)

	// Centers, accounting for ball's size
	// DISABLED for now, makes calculating positions relative to paddles, etc. hard
	// b.Pos.X -= float64(b.Radius)
	// b.Pos.Y -= float64(b.Radius)

	// Sets initial direction of ball
	b.xDirection = RIGHT
	b.yDirection = DOWN

	// Initializes ball to center of window area
	b.ImgOpts.GeoM.Translate(window.Win.CenterX(), window.Win.CenterY())
}

func (b *Ball) SwapDirection(direction *BallDirection) {
	switch *direction {
	case LEFT:
	    *direction = RIGHT
	case RIGHT:
	    *direction = LEFT
	case UP:
	    *direction = DOWN
	case DOWN:
	    *direction = UP
	}
}

func (b *Ball) Reset() {
	// Reset positions
	b.Pos.X, b.Pos.Y = window.Win.CenterX(), window.Win.CenterY()
	//b.Pos.X -= float64(b.Radius / 2)
	//b.Pos.Y -= float64(b.Radius / 2)

	// Reset variable
	b.HasHitPlayer = false

	// Sets direction
	b.SwapDirection(&b.xDirection)
	b.yDirection = DOWN
}

func (b *Ball) Update(playfield *ui.Playfield, rightPaddle, leftPaddle *paddle.Paddle, rightPlayer, leftPlayer *player.Player) {
	var currentSpeed int
	var yAxisNormalizer float64

	if b.HasHitPlayer {
		currentSpeed = b.Speed
		yAxisNormalizer = 0
	} else {
		currentSpeed = b.InitialSpeed
		yAxisNormalizer = -1.5 // Causes ball to go down slowly
	}

	oldPos := BallPosition{ X: b.Pos.X, Y: b.Pos.Y }

	// Check playfield collision
	if b.Pos.Y >= float64(window.Win.Height - playfield.Height) || b.Pos.Y <= float64(playfield.Height) {
		b.SwapDirection(&b.yDirection)
	}

	// Check paddles collision
	if b.Pos.X >= rightPaddle.X - float64(rightPaddle.Width) &&
	   b.Pos.Y < rightPaddle.Y + float64(rightPaddle.Height/2) &&
	   b.Pos.Y > rightPaddle.Y - float64(rightPaddle.Height/2) || 

	   b.Pos.X <= leftPaddle.X + float64(leftPaddle.Width) && 
	   b.Pos.Y < leftPaddle.Y + float64(leftPaddle.Height/2) &&
	   b.Pos.Y > leftPaddle.Y - float64(leftPaddle.Height/2) {
		if !b.HasHitPlayer {
			b.HasHitPlayer = true
			currentSpeed = b.Speed
		}
		b.SwapDirection(&b.xDirection)
	}

	// Ball went in for right player
	if b.Pos.X > float64(window.Win.Width) {
		// Set score
		rightPlayer.Score += 1
		fmt.Println("Right player scored! Score now:", rightPlayer.Score)
		// Reset ball
		b.Reset()

	// Ball went in for left player
	} else if b.Pos.X <= float64(0 + leftPaddle.Width) {
		// Set score
		leftPlayer.Score += 1
		fmt.Println("Left player scored! Score now:", leftPlayer.Score)
		// Reset ball
		b.Reset()
	}

	b.Pos.X += float64(currentSpeed * b.xDirection.Int())
	b.Pos.Y += float64(currentSpeed * b.yDirection.Int()) + float64(yAxisNormalizer)

	b.ImgOpts.GeoM.Translate(b.Pos.X - oldPos.X, b.Pos.Y - oldPos.Y)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.Img, &b.ImgOpts)
}
