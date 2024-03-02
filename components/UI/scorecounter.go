package ui

import (
	"fmt"
	"image/color"

	"gong/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type ScoreCounter struct {
	X        int
	Y        int
	FontFace font.Face
	Score    int
}

func (sc *ScoreCounter) Draw(screen *ebiten.Image) {
	utils.DrawCenteredText(screen, fmt.Sprint(sc.Score), sc.FontFace, sc.X, sc.Y, color.White)
}
