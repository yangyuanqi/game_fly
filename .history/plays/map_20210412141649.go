package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"

	"github.com/hajimehoshi/ebiten"
)

type Map struct {
	core.Sprite
	count int
}

func NewMap() (m *Map) {
	m = &Map{}
	m.Sprite.Create("map")
	return
}

func (m *Map) OnLoad() {
	//背景
	m.SetMaterial(core.Byte2Image(images.M3_jpg))
}

func (d *Map) Start() {

}

func (d *Map) Update(screen *ebiten.Image) (err error) {
	d.count++
	if d.count%60 == 0 {

	}

	d.Sprite.Update(screen)
	return nil
}
