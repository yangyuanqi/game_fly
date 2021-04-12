package plays

import (
	"game_fly/core"
	"game_fly/plays/assets/images"
	"image"

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
		var a = &core.Animator{
			Sprite: core.Sprite{
				Base: core.Base{
					X:        100,
					Y:        100,
					Material: core.Nodes.Img["t1png"],
				},
			},
			Interval: 1,
			Order:    "loop",
			Imgs: []*ebiten.Image{
				core.Nodes.Img["itempng"].SubImage(image.Rect(0, 0, 100, 100)).(*ebiten.Image),
				core.Nodes.Img["itempng"].SubImage(image.Rect(0, 0, 100, 100)).(*ebiten.Image),
				core.Nodes.Img["itempng"].SubImage(image.Rect(0, 0, 100, 100)).(*ebiten.Image),
				core.Nodes.Img["itempng"].SubImage(image.Rect(0, 0, 100, 100)).(*ebiten.Image),
				core.Nodes.Img["itempng"].SubImage(image.Rect(0, 0, 100, 100)).(*ebiten.Image),
				core.Nodes.Img["itempng"].SubImage(image.Rect(0, 0, 100, 100)).(*ebiten.Image),
			},
		}
		a.Move = func(a *core.Animator) {
			a.Y += 1
		}
		core.Nodes.AddPrefab(a)
	}

	d.Sprite.Update(screen)
	return nil
}
