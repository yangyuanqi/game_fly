package scripts

import (
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/plays/assets/images"
	"github.com/325Gerbils/go-vector"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image"
)

type Role struct {
	//component.GameObject
	component.Script
	//prefabs []component.Render
	speed float64
}

func RoleGameObject(xy vector.Vector) (role *Role) {
	r := &Role{
		Script: component.Script{
			GameObject: component.GameObject{
				Del: false,
				Transform: component.Transform{
					Position: xy,
					ScaleX:   0.5,
					ScaleY:   0.5,
				},
				Sprite: component.Sprite{
					Img: core.Byte2Image(images.My_1).SubImage(image.Rect(10, 30, 190, 160)).(*ebiten.Image),
				},
			},
		},
	}
	//添加碰撞
	w, h := r.Sprite.Img.Size()
	r.Collider = *component.NewBoxCollider(r.Transform.Position, float64(w)*r.Transform.ScaleX, float64(h)*r.Transform.ScaleY)
	return r
}

func (r *Role) Update(screen *ebiten.Image) (err error) {
	r.Draw(screen)
	//移动
	moveX := core.GetAxisRaw("X")
	moveY := core.GetAxisRaw("Y")
	r.Collider.Move(vector.New(float64(moveX*5), float64(moveY*5)))

	//发射子弹

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if v := inpututil.KeyPressDuration(ebiten.KeySpace); 0 < v {
			if v == 1 {
				bullet := BulletGameObject(vector.New(r.Transform.Position.X+37, r.Transform.Position.Y), vector.New(0, -1), 8)
				//core.Game.AddPrefab(bullet, "bullet")
				core.Game.AddGameObject(bullet)
			}
		}
	}

	return
}

func (r *Role) OnCollisionEnter(order component.Render) {
	//fmt.Println(r.Transform.Position,order.GetCollider().Transform.Position)
}

func (r *Role) OnTriggerEnter(order component.Render) {
	w, h := order.GetCollider().GetSize()
	w2, h2 := r.GetCollider().GetSize()
	if r.Collider.Transform.Position.X <= 0 {
		r.Collider.Transform.Position.X = 0
	}
	if r.Collider.Transform.Position.X >= w-w2 {
		r.Collider.Transform.Position.X = w - w2
	}
	if r.Collider.Transform.Position.Y <= 0 {
		r.Collider.Transform.Position.Y = 0
	}
	if r.Collider.Transform.Position.Y >= h-h2 {
		r.Collider.Transform.Position.Y = h - h2
	}
}
