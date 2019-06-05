package plays

import (
	"bytes"
	"fmt"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
	"math"
	"strconv"
	"test/play/plays/assets/images"
)

type Role struct {
	CC
	bullet             []*Bullet
	CollideSpaceBullet *resolv.Space
}

var RoleImg *ebiten.Image
var count int

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.My_1))
	if err != nil {
		log.Fatal(err)
	}
	RoleImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

}

func NewRole(game *Game) (role *Role) {
	role = &Role{}
	role.RootNode = game
	role.ScaleW = 0.3
	role.ScaleH = 0.3
	w, h := RoleImg.Size()

	role.W = int(math.Round(float64(w) * role.ScaleW))
	role.H = int(math.Round(float64(h) * role.ScaleH))
	role.X = math.Ceil(float64((game.W - role.W) / 2))
	role.Y = float64(game.H - role.H)

	role.Collide = resolv.NewRectangle(int32(role.X), int32(role.Y), int32(role.W), int32(role.H))
	role.CollideSpaceBullet = resolv.NewSpace()
	return
}

func (r *Role) Update() (err error) {
	count++

	for k, v := range r.bullet {
		v.Update("hero")
		//到达边界删除
		if v.Y < 0 {
			if len(r.bullet) > 1 {
				r.bullet = append(r.bullet[:k], r.bullet[k+1:]...)
			}
		}
	}

	if count == 100 {
		count = 0
	}
	if count%50 == 0 {
		var bullet *Bullet
		for i := 1; i < 7; i++ {
			if i > 3 {
				bullet = NewBullet(r.RootNode, r, float64(i-3))
			} else {
				bullet = NewBullet(r.RootNode, r, float64(-i))
			}

			r.bullet = append(r.bullet, bullet)

			bullet.Collide.AddTags("heroBullet")
			r.CollideSpaceBullet.Add(bullet.Collide)
		}
	}

	r.CheckCollied()

	r.Drow()

	return nil
}

func (r *Role) Drow() {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(r.ScaleW, r.ScaleH)
	opts.GeoM.Translate(r.X, r.Y)
	r.RootNode.screen.DrawImage(RoleImg, opts)
}

func (r *Role) GetPosition() (x, y float64, w, h int) {
	return r.X, r.Y, r.W, r.H
}

func (r *Role) CheckCollied() {
	//敌人和英雄
	colliding := r.RootNode.EnemySpace.FilterByTags("enemy")
	if colliding.IsColliding(r.Collide) {
		fmt.Println(strconv.Itoa(Count) + "WHOA! shape1 and shape2 are colliding.")
		r.RootNode.scenesIng = 100
	}

	//英雄子弹和敌人
	for k, v := range r.RootNode.Enemys {
		colliding2 := v.Collide
		if colliding2.IsColliding(r.CollideSpaceBullet) {
			r.RootNode.Enemys = append(r.RootNode.Enemys[:k], r.RootNode.Enemys[k+1:]...)
		}
	}

	//敌人子弹和英雄
}
