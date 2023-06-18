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
	SCREEN_WIDTH = 600
	SCREEN_HEIGHT = 300

	PADDLE_WIDTH = 5
)

// Game implements ebiten.Game interface.
type Game struct{
	rightPlayer player.Player
	leftPlayer player.Player
	ball ballImport.Ball

	UI ui.UI
}

func NewGame() *Game {
	center := window.Win.Height / 2 / 2

	game := &Game{ 
		leftPlayer: player.Player{ 
			Score: 0, 
			Paddle: paddle.Paddle{
				X: 20, 
				Body: paddle.PaddleBody{
					Body: [5]int{center-2, center-1, center, center+1, center+2},
				},
				Width: PADDLE_WIDTH,
				Speed: 4,
			},
		},
		rightPlayer: player.Player{ 
			Score: 0, 
			Paddle: paddle.Paddle{
				X: window.Win.HalfWidth() - 20, 
				Body: paddle.PaddleBody{
					Body: [5]int{center-2, center-1, center, center+1, center+2},
				},
				Width: PADDLE_WIDTH,
				Speed: 4,
			},
		},
		ball: ballImport.Ball{
			Pos: [2]int{center, center},
			Direction: 1,        
			HasHitPlayer: false,
		},
		UI: ui.UI{
			Separator: ui.Separator{Width: 3, Color: color.White},
			Playfield: ui.Playfield{ScreenWidth: window.Win.Width, ScreenHeight: window.Win.Height},
		},
	}

	// Initializing components
	game.UI.Separator.Init(window.Win.Width, window.Win.Height)
	game.UI.Playfield.Init()
	game.leftPlayer.Paddle.Init()
	game.rightPlayer.Paddle.Init()

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
	g.rightPlayer.Paddle.Draw(screen)
	g.leftPlayer.Paddle.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return window.Win.Width, window.Win.Height
}

func main() {
    ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
    ebiten.SetWindowTitle("Gong - The Go Pong!")

    window.Win = window.Window{ Width: SCREEN_WIDTH * 2, Height: SCREEN_HEIGHT * 2}

	// Start a new game
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
