package plays

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
	"test/play/plays/assets/images"
)

type Screen02 struct {
	CC
}

var ebitenImage *ebiten.Image

func NewScreen02(game *Game) (screen02 *Screen02) {
	screen02 = &Screen02{}
	screen02.RootNode = game
	return
}

func (s *Screen02) Onload() {
	img, _, err := image.Decode(bytes.NewReader(images.Load))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func (s *Screen02) Update() {

	s.Draw()
}

func (s *Screen02) Draw() {
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(0, 0)
	s.RootNode.screen.DrawImage(ebitenImage, op)
}
