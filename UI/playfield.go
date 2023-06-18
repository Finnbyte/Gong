package ui

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Playfield struct {
	ScreenWidth, ScreenHeight int
    Img *ebiten.Image 
    ImgOptsTop ebiten.DrawImageOptions
    ImgOptsBottom ebiten.DrawImageOptions
}

func (pf *Playfield) Init() {
    pf.Img = ebiten.NewImage(pf.ScreenWidth, 3)
	pf.Img.Fill(color.White)

	// Draw the line across the screen
	pf.ImgOptsTop = ebiten.DrawImageOptions{}
	pf.ImgOptsTop.GeoM.Translate(0, 20)

	pf.ImgOptsBottom = ebiten.DrawImageOptions{}
	pf.ImgOptsBottom.GeoM.Translate(0, float64(pf.ScreenHeight - 20))
}
