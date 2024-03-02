package player

import (
	ui "gong/components/UI"
	"gong/components/paddle"
)

type Player struct {
	ScoreCounter ui.ScoreCounter
	Paddle       paddle.Paddle
}
