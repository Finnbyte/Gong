package main

import (
	"gong/UI"
	"gong/window"
	ballImport "gong/ball"
	"gong/paddle"
	"gong/player"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WINDOW_WIDTH = 800
	WINDOW_HEIGHT = 600

	PADDLE_WIDTH = 10
	PADDLE_HEIGHT = 100
	PADDLE_SPEED = 3.0
)

// Game implements ebiten.Game interface.
type Game struct{
	rightPlayer player.Player
	leftPlayer player.Player
	ball ballImport.Ball

	UI ui.UI
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
				X: float64(window.Win.Width - 20), 
				Y: window.Win.CenterY(),
				Width: PADDLE_WIDTH,
				Height: PADDLE_HEIGHT,
				Speed: PADDLE_SPEED,
			},
		},
		ball: ballImport.Ball{
			Pos: [2]int{0, 0},
			Direction: 1,        
			HasHitPlayer: false,
		},
		UI: ui.UI{
			Separator: ui.Separator{Width: 3, Color: color.White},
			Playfield: ui.Playfield{Window: window.Win},
		},
	}

	// Initializing components
	game.UI.Separator.Init(window.Win.Width, window.Win.Height)
	game.UI.Playfield.Init()
	game.leftPlayer.Paddle.Init(game.UI.Playfield)
	game.rightPlayer.Paddle.Init(game.UI.Playfield)

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

    return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw separator
	screen.DrawImage(g.UI.Separator.Img, &g.UI.Separator.ImgOpts)

	// Draw Playfield
	screen.DrawImage(g.UI.Playfield.Img, &g.UI.Playfield.ImgOptsTop)
	screen.DrawImage(g.UI.Playfield.Img, &g.UI.Playfield.ImgOptsBottom)

	// Draw paddles
	g.leftPlayer.Paddle.Draw(screen)
	g.rightPlayer.Paddle.Draw(screen)
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
