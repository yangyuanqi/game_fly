package plays

import (
	"game_fly/core"

	"github.com/SolarLune/resolv"
	"github.com/hajimehoshi/ebiten"
)

type Bullet struct {
	core.Sprite
}

func NewBullet() (bullet *Bullet) {
	b := &Bullet{}
	b.Sprite.Create("bullet")
	//b.SetMaterial(core.Byte2Image(images.Myb_1))
	b.SetMaterial(core.Nodes.Img["bullet"])
	b.SetScale(.3, .3)
	b.Collision = resolv.NewRectangle(b.X, b.Y, b.W, b.H)
	return b
}

func (b *Bullet) Update(screen *ebiten.Image) (err error) {
	b.shape1.SetXY(b.X, b.Y)

	b.Sprite.Update(screen)
	if b.Y < 0 {
		var a = &core.Animator{
			Sprite: core.Sprite{
				Material: core.Nodes.Img["t1png"],
				Base: core.Base{
					X: b.X,
					Y: b.Y,
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
