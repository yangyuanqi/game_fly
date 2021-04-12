package plays

import (
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
	if b.Y < 0 {
		b.Destroy()
		var a core.Animator = &core.Animator{}
	}
	return
}
