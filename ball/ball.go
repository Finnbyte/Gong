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
	verticality float64
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

	b.verticality = -1.5

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

	b.verticality = -1.5
}

func (b *Ball) determineDirectionOnPaddleCollision(bY float64, paddle *paddle.Paddle) {
	centerAreaSeparator := 9
	centerTopLimit := paddle.Y + float64(paddle.Height/2) + float64(centerAreaSeparator)
	centerBottomLimit := paddle.Y + float64(paddle.Height/2) - float64(centerAreaSeparator)
	fmt.Println(centerBottomLimit, centerTopLimit, bY)
	// center segment
	if bY <= centerTopLimit && bY >= centerBottomLimit {
		fmt.Println("hit center")
		b.SwapDirection(&b.xDirection)
		b.verticality = -1
		//b.SwapDirection(&b.yDirection)
		b.yDirection = DOWN

	// head segment
	} else if bY >= paddle.Y && bY <= paddle.Y + float64(paddle.Height/2) {
		fmt.Println("hit head")
		b.SwapDirection(&b.xDirection)
		b.verticality = 0
		b.yDirection = UP

	// tail segment
	} else {
		b.SwapDirection(&b.xDirection)
		fmt.Println("hit tail")
		b.verticality = 0
		b.yDirection = DOWN
	}
}

func (b *Ball) Update(playfield *ui.Playfield, rightPaddle, leftPaddle *paddle.Paddle, rightPlayer, leftPlayer *player.Player) {
	var currentSpeed int

	if b.HasHitPlayer {
		currentSpeed = b.Speed
	} else {
		currentSpeed = b.InitialSpeed
	}

	oldPos := BallPosition{ X: b.Pos.X, Y: b.Pos.Y }

	// Check playfield collision
	if b.Pos.Y >= float64(window.Win.Height - playfield.Height) || b.Pos.Y <= float64(playfield.Height) {
		b.SwapDirection(&b.yDirection)
	}

	// Check paddles collision
	if b.Pos.X >= rightPaddle.X - float64(rightPaddle.Width) &&
	   b.Pos.Y > rightPaddle.Y &&
	   b.Pos.Y < rightPaddle.Y + float64(rightPaddle.Height) || 

	   b.Pos.X <= leftPaddle.X + float64(leftPaddle.Width) && 
	   b.Pos.Y > leftPaddle.Y &&
	   b.Pos.Y < leftPaddle.Y + float64(leftPaddle.Height) {
		if !b.HasHitPlayer {
			b.HasHitPlayer = true
			currentSpeed = b.Speed
		}

		// To use DRY principles, recheck which paddle the ball hit
		// Handling this in one function seems more intuitive for me
		if b.Pos.X >= rightPaddle.X - float64(rightPaddle.Width) {
			b.determineDirectionOnPaddleCollision(b.Pos.Y, rightPaddle)
		} else {
			b.determineDirectionOnPaddleCollision(b.Pos.Y, leftPaddle)
		}
	}

	// Ball went in for right player
	if b.Pos.X > float64(window.Win.Width) {
		// Set score
		leftPlayer.Score += 1
		fmt.Println("Left player scored! Score now:", leftPlayer.Score)
		// Reset ball
		b.Reset()

	// Ball went in for left player
	} else if b.Pos.X <= 0.0 {
		// Set score
		rightPlayer.Score += 1
		fmt.Println("Right player scored! Score now:", rightPlayer.Score)
		// Reset ball
		b.Reset()
	}

	b.Pos.X += float64(currentSpeed * b.xDirection.Int())
	b.Pos.Y += float64(currentSpeed * b.yDirection.Int()) + float64(b.verticality)

	b.ImgOpts.GeoM.Translate(b.Pos.X - oldPos.X, b.Pos.Y - oldPos.Y)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.Img, &b.ImgOpts)
}
