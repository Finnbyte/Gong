package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Separator struct {
	Width   int
	Img     *ebiten.Image
	ImgOpts ebiten.DrawImageOptions
}

func (s *Separator) Init(screenWidth, screenHeight int, color color.Color) {
	s.Img = ebiten.NewImage(s.Width, screenHeight) // - 70 to leave room for top and bottom parts
	s.Img.Fill(color)

	imageRect := s.Img.Bounds().Size()

	halfWidth, halfHeight := screenWidth/2, screenHeight/2
	s.ImgOpts.GeoM.Translate(float64(halfWidth-int(imageRect.X)/2), float64(halfHeight-int(imageRect.Y)/2))
}

func (s *Separator) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.Img, &s.ImgOpts)
}
