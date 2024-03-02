package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gong/components/window"
	"image/color"
)

type Playfield struct {
	Height        int
	Window        window.Window
	Img           *ebiten.Image
	ImgOptsTop    ebiten.DrawImageOptions
	ImgOptsBottom ebiten.DrawImageOptions
}

func (pf *Playfield) Init(color color.Color) {
	pf.Img = ebiten.NewImage(window.Win.Width, pf.Height)
	pf.Img.Fill(color)

	// Draw the line across the screen
	pf.ImgOptsTop = ebiten.DrawImageOptions{}
	pf.ImgOptsTop.GeoM.Translate(0, 0)

	pf.ImgOptsBottom = ebiten.DrawImageOptions{}
	pf.ImgOptsBottom.GeoM.Translate(0, float64(window.Win.Height-10))
}

func (pf *Playfield) Draw(screen *ebiten.Image) {
	screen.DrawImage(pf.Img, &pf.ImgOptsBottom)
	screen.DrawImage(pf.Img, &pf.ImgOptsTop)
}
