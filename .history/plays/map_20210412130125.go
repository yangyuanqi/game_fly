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
	if d.count%5 == 0 {
		var a = &core.Animator{
			Sprite: core.Sprite{
				Base: core.Base{
					X:        100,
					Y:        100,
					Material: core.Nodes.Img["t1png"],
				},
			},
			Interval: 1,
			Order:    "line",
			Imgs:     []*ebiten.Image{core.Nodes.Img["t1png"], core.Nodes.Img["t1png"], core.Nodes.Img["t1png"], core.Nodes.Img["t2png"], core.Nodes.Img["t2png"], core.Nodes.Img["t2png"]},
		}
		a.Y++
		core.Nodes.AddPrefab(a)
	}

	d.Sprite.Update(screen)
	return nil
}
