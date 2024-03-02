package main

import (
	ui "gong/components/UI"
	pongBall "gong/components/ball"
	"gong/components/paddle"
	"gong/components/player"
	. "gong/components/screen"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Game struct {
	rightPlayer player.Player
	leftPlayer  player.Player
	ball        pongBall.Ball
	UI          ui.UI
}

func NewGame() *Game {
	game := &Game{
		leftPlayer: player.Player{
			Score: 0,
			Paddle: paddle.Paddle{
				X:           PADDLE_WALL_GAP,
				Y:           Screen.CenterY(),
			},
		},
		rightPlayer: player.Player{
			Score: 0,
			Paddle: paddle.Paddle{
				X:           Screen.Width - PADDLE_WIDTH - PADDLE_WALL_GAP,
				Y:           Screen.CenterY(),
			},
		},
		ball: pongBall.Ball{
			Pos:          pongBall.BallPosition{X: Screen.CenterX(), Y: Screen.CenterY()},
			Radius:       BALL_RADIUS,
			NormalSpeed:  BALL_SPEED,
			InitialSpeed: BALL_SPEED_INITIAL,
			HasHitPlayer: false,
		},
		UI: ui.UI{
			Separator: ui.Separator{Width: 3},
			Playfield: ui.Playfield{Height: 10},
		},
	}

	return game
}

func (g *Game) Update() error {
	// Quit Game
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	// Left player controls
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.leftPlayer.Paddle.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.leftPlayer.Paddle.MoveDown()
	}

	// Right player controls
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.rightPlayer.Paddle.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.rightPlayer.Paddle.MoveDown()
	}

	g.ball.Update(&g.UI.Playfield, &g.rightPlayer.Paddle, &g.leftPlayer.Paddle, &g.rightPlayer, &g.leftPlayer)

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)

	g.UI.Separator.Draw(screen)
	g.UI.Playfield.Draw(screen)
	g.leftPlayer.Paddle.Draw(screen)
	g.rightPlayer.Paddle.Draw(screen)
	g.ball.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Screen.Width, Screen.Height
}
