package scripts

import (
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/plays/assets/images"
	"github.com/325Gerbils/go-vector"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type Enemy struct {
	component.Prefab
}

func EnemyGameObject() (enemy *Enemy) {
	e := &Enemy{
		Prefab: component.Prefab{
			GameObject: component.GameObject{
				Name:      "enemy",
				IsPrefab:  true,
				Del:       false,
				Speed:     1,
				Direction: vector.New(0, 1),
				Transform: component.Transform{
					ScaleX: 0.5,
					ScaleY: 0.5,
				},
				Sprite: component.Sprite{
					Img: core.Byte2Image(images.Ep_1).SubImage(image.Rect(30, 50, 170, 160)).(*ebiten.Image),
				},
			},
		},
	}

	//添加碰撞
	w, h := e.Sprite.Img.Size()
	e.Collider = *component.NewBoxCollider(e.Transform.Position, float64(w)*e.Transform.ScaleX, float64(h)*e.Transform.ScaleY)
	return e
}

func (e *Enemy) Update(screen *ebiten.Image) (err error) {
	e.Draw(core.Game.GameObjects, screen, "enemy")
	return
}

//检测碰撞
func (e *Enemy) OnCollisionEnter(order component.Render) {
	e.GetGameObject().Destroy()
}

func (e *Enemy) OnTriggerEnter(order component.Render) {
	e.GetGameObject().Destroy()
}
