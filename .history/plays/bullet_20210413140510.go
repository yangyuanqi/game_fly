package plays

import (
	"game_fly/core"

	"github.com/hajimehoshi/ebiten"
)

type Bullet struct {
	core.Sprite
	Move func(b *Bullet)
}

func NewBullet() (bullet *Bullet) {
	b := &Bullet{}
	b.Sprite.Create("bullet")
	//b.SetMaterial(core.Byte2Image(images.Myb_1))
	b.SetMaterial(core.Nodes.Img["bullet"])
	b.SetScale(.3, .3)
	return b
}

func (b *Bullet) Update(screen *ebiten.Image) (err error) {
	b.Move(b)
	b.Sprite.Update(screen)
	if b.Y < 0 {
		var a = &core.Animator{
			Sprite: core.Sprite{
				Base: core.Base{
					X:        b.X,
					Y:        b.Y,
					Material: core.Nodes.Img["t1png"],
				},
			},
			Interval: 1,
			Order:    "line",
			Imgs:     []*ebiten.Image{core.Nodes.Img["t1png"], core.Nodes.Img["t1png"], core.Nodes.Img["t1png"], core.Nodes.Img["t2png"], core.Nodes.Img["t2png"], core.Nodes.Img["t2png"]},
		}
		a.Move = func(a *core.Animator) {}
		core.Nodes.AddPrefab(a)
		b.Destroy()
	}
	return
}
