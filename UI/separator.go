package ui

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Separator struct {
    Width int
    Color color.Color
    Img *ebiten.Image
    ImgOpts ebiten.DrawImageOptions
}


func (s *Separator) Init(screenWidth, screenHeight int) {
	s.Img = ebiten.NewImage(s.Width, screenHeight) // - 70 to leave room for top and bottom parts
	s.Img.Fill(s.Color)

	imageRect := s.Img.Bounds().Size()

    halfWidth, halfHeight := screenWidth/2, screenHeight/2
	s.ImgOpts.GeoM.Translate(float64(halfWidth - int(imageRect.X)/2), float64(halfHeight - int(imageRect.Y)/2))
}

