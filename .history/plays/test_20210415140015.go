package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"

	"github.com/hajimehoshi/ebiten"
)

type Test struct {
	core.Sprite
}

func (t *Test) OnLoad() {
	t.SetMaterial(core.Byte2Image(images.My_1))
}

func (t *Test) Update(screen *ebiten.Image) (err error) {

	t.Sprite.Update(screen)
	return nil
}
