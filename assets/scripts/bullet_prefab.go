package scripts

import (
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/plays/assets/images"
	"github.com/325Gerbils/go-vector"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type Bullet struct {
	component.Prefab
}

func BulletGameObject(vector2, direction vector.Vector, speed float64) (bullet *Bullet) {
	e := &Bullet{
		Prefab: component.Prefab{
			GameObject: component.GameObject{
				Name:      "bullet",
				IsPrefab:  true,
				Del:       false,
				Speed:     speed,
				Direction: direction,
				Transform: component.Transform{
					ScaleX:   0.4,
					ScaleY:   0.4,
					Position: vector2,
				},
				Sprite: component.Sprite{
					Img: core.Byte2Image(images.Myb_1).SubImage(image.Rect(30, 10, 65, 80)).(*ebiten.Image),
				},
			},
		},
	}

	//添加碰撞
	w, h := e.Sprite.Img.Size()
	e.Collider = *component.NewBoxCollider(e.Transform.Position, float64(w)*e.Transform.ScaleX, float64(h)*e.Transform.ScaleY)
	return e
}

func (e *Bullet) Update(screen *ebiten.Image) (err error) {
	e.Draw(core.Game.GameObjects, screen, "bullet")
	return
}

//检测碰撞
func (b *Bullet) OnCollisionEnter(order component.Render) {

	//fmt.Printf("%T", order.GetGameObject().GetGameObject())
	//b.GetGameObject().Destroy()
}

func (b *Bullet) OnTriggerEnter(order component.Render) {
	b.GetGameObject().Destroy()
}
