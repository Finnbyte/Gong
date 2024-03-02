package ball

import (
	ui "gong/components/UI"
	"gong/components/paddle"
	"gong/components/player"
	. "gong/components/screen"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
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
}

func (b *Ball) Reset() {
	b.Pos.X, b.Pos.Y = Screen.CenterX(), Screen.CenterY()
	b.HasHitPlayer = false
	b.VelocityX = -(b.VelocityX)
	b.VelocityY = -(b.VelocityY)
}

func (b *Ball) collidedWithPaddle(p paddle.Paddle) bool {
	if p.X <= b.Pos.X && b.Pos.X <= p.X+p.Width+p.StrokeWidth || p.X >= b.Pos.X && p.X-p.Width-p.StrokeWidth <= b.Pos.X {
		if p.Y <= b.Pos.Y+b.Radius && b.Pos.Y <= p.Y+p.Height {
			return true
		}
	}
	return false
}

func (b *Ball) collidedWithPlayfield(pf ui.Playfield) bool {
	if b.Pos.Y >= Screen.Height-(pf.Height+b.Radius) || b.Pos.Y <= pf.Height {
		return true
	}
	return false
}

func (b *Ball) Update(playfield *ui.Playfield, rightPaddle, leftPaddle *paddle.Paddle, rightPlayer, leftPlayer *player.Player) {
	var SPEED = b.InitialSpeed

	if b.Pos.X > Screen.Width {
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
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.Pos.X), float32(b.Pos.Y), float32(b.Radius), float32(b.Radius), colornames.White, false)
}
