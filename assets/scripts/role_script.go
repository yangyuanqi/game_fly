package scripts

import (
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/plays/assets/images"
	"github.com/325Gerbils/go-vector"
	"github.com/hajimehoshi/ebiten"
	"image"
	"math"
)

type Role struct {
	component.GameObject
}

func RoleGameObject(xy vector.Vector) (role *Role) {
	r := &Role{}
	r.Sprite.Img = core.Byte2Image(images.My_1)
	r.Transform.Position = xy
	//添加碰撞
	w, h := r.Sprite.Img.Size()
	r.Collider = *component.NewBoxCollider(r.Transform.Position, math.Ceil(float64(w)/2.2), math.Ceil(float64(h)/3))
	return r
}

func (r *Role) Update(screen *ebiten.Image) (err error) {
	opts := &ebiten.DrawImageOptions{}
	//	//opts.GeoM.Reset()
	opts.GeoM.Scale(0.5, 0.5)
	//opts.GeoM.Rotate(-1.1)//旋转
	//opts.GeoM.Skew(1,1)//倾斜
	opts.GeoM.Translate(r.Collider.Position.X, r.Collider.Position.Y)

	screen.DrawImage(r.Sprite.Img.SubImage(image.Rect(10, 30, 200, 200)).(*ebiten.Image), opts)

	//移动
	moveX := core.GetAxisRaw("X")
	moveY := core.GetAxisRaw("Y")
	r.Collider.Move(vector.New(float64(moveX*5), float64(moveY*5)))
	return
}
