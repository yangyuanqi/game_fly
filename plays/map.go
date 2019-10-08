package plays

import (
	"game_fly/core"
	"game_fly/plays/conf"
	"github.com/hajimehoshi/ebiten"
)

type Map struct {
	core.Sprite
}

func NewMap() (m *Map) {
	m = &Map{}
	m.Sprite.Create("map")
	return
}

func (m *Map) OnLoad() {
	//背景
	m.SetMaterial(MapImg)

	sW, sH := conf.GetConfInt("scenes_width"), conf.GetConfInt("scenes_height")
	m.SetScale(core.ScaliEq(float64(sW), float64(m.W)), core.ScaliEq(float64(sH), float64(m.H)))
}

func (d *Map) Start() {

}

func (d *Map) Update(screen *ebiten.Image) (err error) {
	d.Sprite.Update(screen)
	return nil
}
