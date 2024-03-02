package main

import (
	ui "gong/components/UI"
	. "gong/components/ball"
	"gong/components/paddle"
	. "gong/components/player"
	. "gong/components/screen"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/colornames"
)

type Game struct {
	rightPlayer Player
	leftPlayer  Player
	ball        Ball
	UI          ui.UI
}

func NewGame() *Game {
	// Get a font for scorecounters
	font, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatalf("Error parsing font: %s", err)
	}
	fontFace := truetype.NewFace(font, &truetype.Options{Size: SCORECOUNTER_FONT_SIZE})

	game := &Game{
		leftPlayer: Player{
			ScoreCounter: ui.ScoreCounter{Score: 0, X: Screen.HalfWidth() - SCORECOUNTER_GAP_FROM_CENTER, Y: SCORECOUNTER_GAP_FROM_TOP, FontFace: fontFace},
			Paddle: paddle.Paddle{
				X:           PADDLE_WALL_GAP,
				Y:           Screen.CenterY(),
				StrokeWidth: PADDLE_STROKE_WIDTH,
				Width:       PADDLE_WIDTH,
				Height:      PADDLE_HEIGHT,
			},
		},
		rightPlayer: Player{
			ScoreCounter: ui.ScoreCounter{Score: 0, X: Screen.HalfWidth() + SCORECOUNTER_GAP_FROM_CENTER, Y: SCORECOUNTER_GAP_FROM_TOP, FontFace: fontFace},
			Paddle: paddle.Paddle{
				X:           Screen.Width - PADDLE_WIDTH - PADDLE_WALL_GAP,
				Y:           Screen.CenterY(),
				StrokeWidth: PADDLE_STROKE_WIDTH,
				Width:       PADDLE_WIDTH,
				Height:      PADDLE_HEIGHT,
			},
		},
		ball: Ball{
			Pos:          BallPosition{X: Screen.CenterX(), Y: Screen.CenterY()},
			Radius:       BALL_RADIUS,
			NormalSpeed:  BALL_SPEED,
			VelocityX:    BALL_INITIAL_VELOCITY_X,
			VelocityY:    BALL_INITIAL_VELOCITY_Y,
			InitialSpeed: BALL_SPEED_INITIAL,
			HasHitPlayer: false,
		},
		UI: ui.UI{
			Separator: ui.Separator{Width: 5},
			Playfield: ui.Playfield{Height: 10},
		},
	}

	return game
}

func (g *Game) Update() error {
	// Quit Game
	if ebiten.IsKeyPressed(ebiten.KeyQ) || ebiten.IsKeyPressed(ebiten.KeyEscape) {
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
	g.leftPlayer.ScoreCounter.Draw(screen)
	g.rightPlayer.ScoreCounter.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Screen.Width, Screen.Height
}
