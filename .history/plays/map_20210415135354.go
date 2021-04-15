package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten"
)

type Map struct {
	core.Sprite
	count  int
	shape2 *resolv.Rectangle
}

func NewMap() (m *Map) {
	m = &Map{}
	m.Sprite.Create("map")
	m.shape2 = resolv.NewRectangle(m.X, m.Y, m.W, m.H)
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
		a := NewStar()
		a.Move = func(a *core.Animator) {
			a.Y += 1
			//a.X = int32(0.5 * float64(a.Y^2))
		}
		core.Nodes.AddPrefab(a)
	}

	d.Sprite.Update(screen)
	return nil
}
