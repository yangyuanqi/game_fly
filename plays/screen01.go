package plays

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
	"test/play/plays/assets/images"
)

type Screen01 struct {
	CC
}

var screen01Image *ebiten.Image

func NewScreen01(game *Game) (screen02 *Screen01) {
	screen02 = &Screen01{}
	screen02.RootNode = game
	return
}

func (s *Screen01) Onload() {
	img, _, err := image.Decode(bytes.NewReader(images.Load))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func (s *Screen01) Update() {

	s.Draw()
}

func (s *Screen01) Draw() {
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(0, 0)
	s.RootNode.screen.DrawImage(ebitenImage, op)
}
