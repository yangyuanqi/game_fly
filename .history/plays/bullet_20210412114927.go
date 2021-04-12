package plays

import (
	"fmt"
	"game_fly/core"

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
	return b
}

func (b *Bullet) Update(screen *ebiten.Image) (err error) {
	b.Y -= 10
	b.Sprite.Update(screen)
	if b.Y < 100 {
		fmt.Println(123)
		if b.Destroy() {

			var a = &core.Animator{
				Sprite: core.Sprite{

					Base: core.Base{
						X:        b.X,
						Y:        b.Y,
						Material: core.Nodes.Img["t1png"],
						SkewX:    1,
						SkewY:    1,
					},
				},
				Interval: 5,
				Order:    "line",
				Imgs:     []*ebiten.Image{core.Nodes.Img["t1png"], core.Nodes.Img["t2png"]},
			}
			core.Nodes.AddPrefab(a)
		}
	}
	return
}
