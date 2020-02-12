package scripts

import (
	"game_fly/assets/prefabs"
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/plays/assets/images"
	"github.com/325Gerbils/go-vector"
	"github.com/hajimehoshi/ebiten"
	"image"
	"math"
)

type Enemy struct {
	component.GameObject
	prefab prefabs.Bullet
}

func EnemyGameObject() (enemy *Enemy) {
	e := &Enemy{}
	e.Sprite.Img = core.Byte2Image(images.Ep_1)

	//
	w, h := e.Sprite.Img.Size()
	e.Collider = *component.NewBoxCollider(e.Transform.Position, math.Floor(float64(w)/2.9), math.Floor(float64(h)/4))
	return e
}

func (e *Enemy) Update(screen *ebiten.Image) (err error) {
	e.Collider.Position.Add(vector.New(0, 1))

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(0.5, 0.5)
	opts.GeoM.Translate(e.Collider.Position.X, e.Collider.Position.Y)
	screen.DrawImage(e.Sprite.Img.SubImage(image.Rect(30, 50, 180, 200)).(*ebiten.Image), opts)
	return
}

func (e *Enemy) OnCollisionEnter(order component.Render) {
	//order.GetCollider().Position.Add(vector.New(1, 0))
	e.GetGameObject().Destroy()
}
