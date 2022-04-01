package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Sprite struct {
	Active         bool
	CustomMaterial *ebiten.Image
	Color          color.Color
	SpriteAtlas    []*ebiten.Image
	SpriteFrame    *ebiten.Image
	Grayscale      bool
	SizeMode       uint
	Type           uint
	Trim           bool
}

func (s *Sprite) Update(dt float64) error {
	return nil
}

func (s *Sprite) Draw(screen *ebiten.Image) {

}
