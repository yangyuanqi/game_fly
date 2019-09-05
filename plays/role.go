package plays

import (
	"bytes"
	"fmt"
	"game_fly/core"
	"game_fly/core/component"
	"game_fly/plays/assets/images"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
	"math"
)

type Role struct {
	core.Sprite
	RootNode  *Game
	Left      int
	Right     int
	RoleImg   *ebiten.Image
	BulletImg *ebiten.Image
}

var RoleImg *ebiten.Image

func NewRole() (role *Role) {
	role = &Role{}

	//role.BulletImg = bulletImg
	//role.RootNode = rootNode
	role.SetScale(0.3, 0.3)
	return
}

func (r *Role) OnLoad() {
	//角色贴图
	img, _, err := image.Decode(bytes.NewReader(images.My_1))
	if err != nil {
		log.Fatal(err)
	}
	RoleImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	r.RoleImg = RoleImg

	w, h := RoleImg.Size()
	r.SetWH(int(math.Round(float64(w)*r.ScaleW)), int(math.Round(float64(h)*r.ScaleH)))
	fmt.Println(1, core.GetScene("game"))
	game := core.GetScene("game").(*Game)
	r.SetXY(math.Ceil(float64((game.W-r.W)/2)), float64(game.H-r.H))

	component.AddComponent(NewInput(), "input")

	//1号自动子弹
	//core.SetTicker(time.Millisecond*200, r.AutoBullet01, 0)
}

func (r *Role) Update(screen *ebiten.Image) (err error) {
	//for _, v := range r.GroupBullet {
	//	v.Update()
	//}


	r.Drow()

	//r.CheckCollied()

	return nil
}

func (r *Role) Drow() {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Reset()
	opts.GeoM.Scale(r.ScaleW, r.ScaleH)

	//灰机倾斜处理
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
	core.GetScene("game").(*Game).Screen.DrawImage(r.RoleImg, opts)
	//r.RootNode.Screen.DrawImage(r.RoleImg, opts)
}

//func (r *Role) CheckCollied() {
//	//敌人和英雄
//
//	//英雄子弹和敌人
//	for _, v := range r.RootNode.Enemys {
//		for _, v2 := range r.GroupBullet {
//			if v.Visible == true && v2.Visible == true && core.CheckCollision(v, v2) {
//				v.Visible = false
//				v2.Visible = false
//
//			}
//		}
//	}
//}

//var bullet *Bullet
//
//func (r *Role) AutoBullet01() {
//	for i := 1; i < 7; i++ {
//		if i > 3 {
//			bullet = NewBullet(r.RootNode, r, r.BulletImg)
//			bullet.MX = float64(i-3) / 10
//			bullet.MY = -10
//		} else {
//			bullet = NewBullet(r.RootNode, r, r.BulletImg)
//			bullet.MX = float64(-i) / 10
//			bullet.MY = -10
//		}
//		r.GroupBullet = append(r.GroupBullet, bullet)
//	}
//}
