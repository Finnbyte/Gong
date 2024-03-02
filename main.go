package main

import (
	. "gong/components/screen"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450

	PADDLE_WIDTH        = 15
	PADDLE_STROKE_WIDTH = 3
	PADDLE_HEIGHT       = 90
	PADDLE_SPEED        = 3.0
	PADDLE_WALL_GAP     = 20

	BALL_INITIAL_VELOCITY_X = 2
	BALL_INITIAL_VELOCITY_Y = 3
	BALL_RADIUS             = 22
	BALL_SPEED              = 4
	BALL_SPEED_INITIAL      = 2
)

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Gong - The Go Pong!")

	Screen = ScreenHelper{Width: WINDOW_WIDTH, Height: WINDOW_HEIGHT}

	var game ebiten.Game = NewGame()

	// Start a new game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
