package ball

import (
	ui "gong/components/UI"
	"gong/components/paddle"
	"gong/components/player"
	"gong/components/window"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type BallPosition struct {
	X, Y int
}

type Ball struct {
	Radius                    int
	NormalSpeed, InitialSpeed int
	Pos                       BallPosition
	VelocityY                 int
	VelocityX                 int
	HasHitPlayer              bool
	Img                       *ebiten.Image
	ImgOpts                   ebiten.DrawImageOptions
}

func (b *Ball) Init(color color.Color) {
	b.Img = ebiten.NewImage(b.Radius, b.Radius)

	b.Img.Fill(color)

	// Centers, accounting for ball's size
	// DISABLED for now, makes calculating positions relative to paddles, etc. hard
	// b.Pos.X -= float64(b.Radius)
	// b.Pos.Y -= float64(b.Radius)

	b.VelocityX = 2
	b.VelocityY = 3

	// Initializes ball to center of window area
	b.ImgOpts.GeoM.Translate(float64(window.Win.CenterX()), float64(window.Win.CenterY()))
}

func (b *Ball) Reset() {
	b.Pos.X, b.Pos.Y = window.Win.CenterX(), window.Win.CenterY()
	b.HasHitPlayer = false
	b.VelocityX = -(b.VelocityX)
	b.VelocityY = -(b.VelocityY)
}

func (b *Ball) collidedWithPaddle(p paddle.Paddle) bool {
	if (b.Pos.X == p.X-b.Radius || b.Pos.X == p.X) && p.Y <= b.Pos.Y+b.Radius && b.Pos.Y <= p.Y+p.Height {
		return true
	}
	return false
}

func (b *Ball) collidedWithPlayfield(pf ui.Playfield) bool {
	if b.Pos.Y >= window.Win.Height-(pf.Height+b.Radius) || b.Pos.Y <= pf.Height {
		return true
	}
	return false
}

func (b *Ball) Update(playfield *ui.Playfield, rightPaddle, leftPaddle *paddle.Paddle, rightPlayer, leftPlayer *player.Player) {
	var SPEED = b.InitialSpeed
	oldPos := BallPosition{X: b.Pos.X, Y: b.Pos.Y}

	if b.Pos.X > window.Win.Width {
		leftPlayer.Score += 1
		b.Reset()
	} else if b.Pos.X <= 0.0 {
		rightPlayer.Score += 1
		b.Reset()
	}

	// Check playfield collision
	if b.collidedWithPlayfield(*playfield) {
		b.VelocityY = -(b.VelocityY)
	}

	if b.collidedWithPaddle(*leftPaddle) || b.collidedWithPaddle(*rightPaddle) {
		if !b.HasHitPlayer {
			SPEED = b.NormalSpeed
		}
		b.VelocityX = -b.VelocityX
	}

	b.Pos.X += SPEED * b.VelocityX
	b.Pos.Y += SPEED * b.VelocityY

	b.ImgOpts.GeoM.Translate(float64(b.Pos.X-oldPos.X), float64(b.Pos.Y-oldPos.Y))
}

func (b *Ball) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.Img, &b.ImgOpts)
}
