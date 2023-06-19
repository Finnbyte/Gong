package main

import (
	"gong/UI"
	"gong/window"
	pongBall "gong/ball"
	"gong/paddle"
	"gong/player"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WINDOW_WIDTH = 600
	WINDOW_HEIGHT = 600

	PADDLE_WIDTH = 10
	PADDLE_HEIGHT = 100
	PADDLE_SPEED = 3.0

	BALL_RADIUS = 22
	BALL_SPEED = 4 
	BALL_SPEED_INITIAL = 2
)

// Game implements ebiten.Game interface.
type Game struct{
	rightPlayer player.Player
	leftPlayer player.Player
	ball pongBall.Ball

	UI ui.UI

	background color.Color
}

func NewGame() *Game {
	game := &Game{ 
		leftPlayer: player.Player{ 
			Score: 0, 
			Paddle: paddle.Paddle{
				X: 20.0, 
				Y: window.Win.CenterY(),
				Width: PADDLE_WIDTH,
				Height: PADDLE_HEIGHT,
				Speed: PADDLE_SPEED,
			},
		},
		rightPlayer: player.Player{ 
			Score: 0, 
			Paddle: paddle.Paddle{
				X: float64(window.Win.Width) - 30, 
				Y: window.Win.CenterY(),
				Width: PADDLE_WIDTH,
				Height: PADDLE_HEIGHT,
				Speed: PADDLE_SPEED,
			},
		},
		ball: pongBall.Ball{
			Pos: pongBall.BallPosition{X: window.Win.CenterX(), Y: window.Win.CenterY() - 100},
			Radius: BALL_RADIUS,
			Speed: BALL_SPEED,
			InitialSpeed: BALL_SPEED_INITIAL,
			HasHitPlayer: false,
		},
		UI: ui.UI{
			Separator: ui.Separator{Width: 3},
			Playfield: ui.Playfield{Window: window.Win},
		},
	}

	// Define colors
	separatorColor := color.RGBA{R: 173, G: 127, B: 168, A: 255}
	topBottomBorderColor := color.RGBA{R: 181, G: 137, B: 0, A: 255}
	paddleColor := color.RGBA{R: 253, G: 235, B: 208, A: 255}
	ballColor := color.RGBA{R: 235, G: 219, B: 178, A: 255}

	// Initializing components
	game.UI.Separator.Init(window.Win.Width, window.Win.Height, separatorColor)
	game.UI.Playfield.Init(topBottomBorderColor)
	game.leftPlayer.Paddle.Init(game.UI.Playfield, paddleColor)
	game.rightPlayer.Paddle.Init(game.UI.Playfield, paddleColor)
	game.ball.Init(ballColor)

	// Set background color
	game.background = color.RGBA{R: 40, G: 40, B: 40, A: 1}

	// Return struct instance for runGame()
	return game
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
    // Quit Game
	if (ebiten.IsKeyPressed(ebiten.KeyQ)) {
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

	g.ball.Update(&g.rightPlayer.Paddle, &g.leftPlayer.Paddle, &g.rightPlayer, &g.leftPlayer)

    return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
    // Set bg 
	screen.Fill(g.background)

	// Draw separator
	screen.DrawImage(g.UI.Separator.Img, &g.UI.Separator.ImgOpts)

	// Draw Playfield
	screen.DrawImage(g.UI.Playfield.Img, &g.UI.Playfield.ImgOptsTop)
	screen.DrawImage(g.UI.Playfield.Img, &g.UI.Playfield.ImgOptsBottom)

	// Draw paddles
	g.leftPlayer.Paddle.Draw(screen)
	g.rightPlayer.Paddle.Draw(screen)

	// Draw ball
	g.ball.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return window.Win.Width, window.Win.Height
}

func main() {
    ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
    ebiten.SetWindowTitle("Gong - The Go Pong!")

    window.Win = window.Window{ Width: WINDOW_WIDTH, Height: WINDOW_HEIGHT }

	// Start a new game
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
