package main

import (
	"gong/components/window"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450

	PADDLE_WIDTH  = 10
	PADDLE_HEIGHT = 100
	PADDLE_SPEED  = 3.0

	BALL_RADIUS        = 22
	BALL_SPEED         = 4
	BALL_SPEED_INITIAL = 2
)

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Gong - The Go Pong!")

	window.Win = window.Window{Width: WINDOW_WIDTH, Height: WINDOW_HEIGHT}

	var game ebiten.Game = NewGame()

	// Start a new game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
