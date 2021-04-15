package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten"
)

type Test struct {
	core.Sprite
}

func NewTest() *Test {
	return &Test{}
}

func (t *Test) OnLoad() {
	t.SetMaterial(core.Byte2Image(images.My_1))
	t.SetXY(100, 100)
	t.shape2 = resolv.NewRectangle(t.X, t.Y, t.W, t.H)
}

func (t *Test) Update(screen *ebiten.Image) (err error) {

	t.Sprite.Update(screen)
	return nil
}
