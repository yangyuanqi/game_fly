package plays

import (
	"bytes"
	"fmt"
	"game_fly/plays/assets/images"
	"github.com/SolarLune/resolv/resolv"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
	"math"
	"time"
)

type Role struct {
	Sprite
	bullet             []*Bullet
	CollideSpaceBullet *resolv.Space
	Left               int
	Right              int
}

var RoleImg *ebiten.Image
var count int
var bullet *Bullet
var ticker *time.Ticker

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.My_1))
	if err != nil {
		log.Fatal(err)
	}
	RoleImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	ticker = time.NewTicker(time.Second * 1)
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
	for k, v := range r.bullet {
		//到达边界删除
		if v.Y < 0 {
			if k < len(r.bullet) {
				r.bullet = append(r.bullet[:k], r.bullet[k+1:]...)
			}
		}
		v.Update("hero")
	}

	//1号自动子弹
	r.AutoBullet01()

	r.Drow(count)

	r.CheckCollied()

	return nil
}

func (r *Role) Drow(count int) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Reset()
	opts.GeoM.Scale(r.ScaleW, r.ScaleH)

	if r.Right > 0 || r.Left > 0 {
		if r.Right <= 20 {
			opts.GeoM.Skew(0, float64(r.Right)/100)
		} else if r.Right > 20 {
			opts.GeoM.Skew(0, 0.2)
		}
		if r.Left <= 20 {
			opts.GeoM.Skew(0, -float64(r.Left)/100)
		} else if r.Left > 20 {
			opts.GeoM.Skew(0, -0.2)
		}
	}

	opts.GeoM.Translate(r.X, r.Y)
	r.RootNode.screen.DrawImage(RoleImg, opts)
}

func (r *Role) CheckCollied() {
	//敌人和英雄
	/*	colliding := r.RootNode.EnemySpace.FilterByTags("enemy")
		if colliding.IsColliding(r.Collide) {
			fmt.Println(strconv.Itoa(Count) + "WHOA! shape1 and shape2 are colliding.")
			r.RootNode.scenesIng = 100
		}*/

	//英雄子弹和敌人
	for k, v := range r.RootNode.Enemys {
		for k2, v2 := range r.bullet {
			if CheckCollision(v, v2) {
				if k < len(r.RootNode.Enemys) {
					r.RootNode.Enemys = append(r.RootNode.Enemys[:k], r.RootNode.Enemys[k+1:]...)
					r.bullet = append(r.bullet[:k2], r.bullet[k2+1:]...)
				}
			}
		}
	}

	//敌人子弹和英雄

	fmt.Println(r.CollideSpaceBullet.Length())
}

func (r *Role) AutoBullet01() {
	go func() {
		select {
		case <-ticker.C:
			func() {
				for i := 1; i < 7; i++ {
					if i > 3 {
						bullet = NewBullet(r.RootNode, r, float64(i-3))
					} else {
						bullet = NewBullet(r.RootNode, r, float64(-i))
					}
					r.bullet = append(r.bullet, bullet)
				}
			}()
		}
	}()
}
