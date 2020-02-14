package scripts

import (
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
)

type Background struct {
	component.Script
}

func BackgroundObject() (background *Background) {
	b := &Background{
		Script:component.Script{
			GameObject: component.GameObject{
				Del: false,
				Transform: component.Transform{
					ScaleX: 0.8,
					ScaleY: 0.8,
				},
				Sprite: component.Sprite{
					Img: core.Byte2Image(images.M3_jpg),
				},
			},
		},
	}

	w, h := b.Sprite.Img.Size()
	b.Collider = *component.NewBoxCollider(b.Transform.Position, float64(w)*b.Transform.ScaleX, float64(h)*b.Transform.ScaleY)
	b.Collider.IsTrigger = true
	return b
}

func (b *Background) Update(screen *ebiten.Image) (err error) {
	b.Draw(screen)
	return
}
