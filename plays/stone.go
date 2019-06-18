package plays

import (
	"game_fly/core"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Stone struct {
	core.BaseComponent
}

var (
	img2 *ebiten.Image
)

func (s *Stone) OnLoad() {
	s.W = 20
	s.H = 200
	s.X = 200
	s.Y = 540

	img2, _ = ebiten.NewImage(int(s.W), int(s.H), ebiten.FilterDefault)
	img2.Fill(color.White)
}

func (s *Stone) Update(screen *ebiten.Image) {
	opts2 := &ebiten.DrawImageOptions{}
	//opts2.GeoM.Scale(b.ScaleW, b.ScaleH)
	opts2.GeoM.Rotate(0.9)
	opts2.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(img2, opts2)
}
