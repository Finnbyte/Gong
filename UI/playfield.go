package ui

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"gong/window"
)

type Playfield struct {
	Window window.Window
	Img *ebiten.Image 
    ImgOptsTop ebiten.DrawImageOptions
    ImgOptsBottom ebiten.DrawImageOptions
}

func (pf *Playfield) Init() {
    pf.Img = ebiten.NewImage(window.Win.Width, 10)
	pf.Img.Fill(color.White)

	// Draw the line across the screen
	pf.ImgOptsTop = ebiten.DrawImageOptions{}
	pf.ImgOptsTop.GeoM.Translate(0, 0)

	pf.ImgOptsBottom = ebiten.DrawImageOptions{}
	pf.ImgOptsBottom.GeoM.Translate(0, float64(window.Win.Height - 10))
}
